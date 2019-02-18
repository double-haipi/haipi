package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cmdParamsTest() {
	for index, args := range os.Args[1:] {
		fmt.Println("arg"+strconv.Itoa(index), args)
	}

	fmt.Println("print2:")
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Println("print3")
	fmt.Println(os.Args[1:])
}

var b = flag.Bool("b", false, "bool 类型参数")
var s = flag.String("s", "", "string 类型参数")

func cmdParasByFlagTest() {
	flag.Parse()
	fmt.Println("-b:", *b)
	fmt.Println("-s:", *s)
	fmt.Println("others:", flag.Args())
}
