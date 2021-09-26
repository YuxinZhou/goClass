Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个`sql.ErrNoRows`的时候，是否应该Wrap这个error，抛给上层。为什么，应该怎么做请写出代码？  
A: 应该。用wrap可以保存堆栈信息。抛给上层，上层可以根据业务逻辑来决定如何处理这个逻辑。
```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "source")
	if err != nil {
		panic(err)
	}
}

func queryUserByID(UserID string) (string, error) {
	var name string
	const userByIDQuery = "select name from user where id=$1"
	err := db.QueryRow(userByIDQuery, UserID).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrapf(err, fmt.Sprintf("user not found: id = %v", UserID))
		}
		return "", errors.Wrapf(err, fmt.Sprintf("unexpected DB err when query user: id = %v", UserID))
	}
	return name, nil
}
```
