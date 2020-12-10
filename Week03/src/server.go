package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ptapp.cn/util/sync/errgroup"
)

func main() {
	serv := http.Server{
		Addr:           ":8081",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    0,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/hello", Hello)
	sigs := make(chan os.Signal, 1)
	//done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	eg := errgroup.Group{}
	eg.Go(func(ctx context.Context) error {
		err := serv.ListenAndServe()
		fmt.Println("err", err)
		fmt.Println("bye bye...")
		return err
	})
	eg.Go(func(ctx context.Context) error {
		sig := <-sigs
		fmt.Println("get signal: ", sig, ". shutting down...")
		if err := serv.Shutdown(ctx); err != nil {
			fmt.Errorf("Shutdown err: %v", err)
			return err
		}
		fmt.Println("shutdown")
		return nil
	})
	if err := eg.Wait(); err != nil {
		fmt.Println("done")
		return
	}
}

func Hello(response http.ResponseWriter, request *http.Request) {
	log.Println("hello msg: %v", request.Body)
	time.Sleep(5 * time.Second)
	if _, err := response.Write([]byte("oh~hello.")); err != nil {
		log.Println("response.Write err", err)
		return
	}
	log.Println("done", request.Body)
}
