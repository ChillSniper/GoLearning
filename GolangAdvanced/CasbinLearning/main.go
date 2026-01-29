package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println("fuck")
		panic(err)
	}
	sub, obj, act := "zhangsan", "dataA", "read"

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		panic(err)
	}
	if ok == true {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
