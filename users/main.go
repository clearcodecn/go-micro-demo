package main

import (
	"go-micro-demo/users/cmd"
	"fmt"
)

func main()  {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("start error : %v", err)
	}
}
