# Dao 层未找到记录的处理
## 基本思路

1. 针对语义是 GetXX (获取单条记录) 的操作，要报错
    - 因为已经不是底层错误，所以要封装
    - 调用方要能够明确判断是否为找到记录的错误
    - 要能够记录详细的堆栈信息方便排查问题
2. 针对语义是 Select (拉取多条记录) 的操作，返回空的 slice

## 代码示例

定义业务通用错误类型
```Golang
package unionerr
import "errors"

RecordNotFound = errors.New("record not found")
```

dao 层识别 `sql.ErrNoRows` 并抛出通用业务错误

```Golang
// dao
type User Struct {
    ID uint64
    Name string
}

func NewDb(dbsource) (sqlx.DB, error) {
	return sqlx.Open("mysql", dbsource)
}

func (db *DB)GetUserById(ctx context.Context, id uint64) (*User, error) {
    var user User
    err := db.Get(ctx, &user, "select name from user where id = %d", id)
    if err == sql.ErrNoRows {
        return nil, unionerr.RecordNotFound
    }
    if err != nil {
        return nil, errors.Wrap(err, "getUserById")
    }
    return &user, nil
}

```

service 层处理，识别通用业务错误，返回预定义错误码
```Golang
// service
    user, err := s.dao.GetUserById(ctx, uid)
    if errors.Is(unionerr.RecordNotFound) {
        return code404
    }
    if err != nil {
        return code500
    }
    // ...
```
