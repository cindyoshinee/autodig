package main

import (
	"fmt"

	"github.com/cindyoshinee/autodig/dep"
)

func main() {
	err := dep.NewAutodig([]string{"./demo"}, "./demo", "").GenDigFile()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("=========autodig success!!==========")
}
