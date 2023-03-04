package sqlDialects

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	dialect := Oracle()
	sql, err := dialect.Page(3, 10, "select * from users where age>?")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql)
}
