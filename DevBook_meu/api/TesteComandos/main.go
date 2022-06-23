package main

import (
	"fmt"
	"strconv"
)

func main() {

	a, erro := strconv.ParseUint("7.000000", 10, 64)

	fmt.Println(a)
	fmt.Println(erro)

}

/*
f := "3.14159265"
if s, err := strconv.ParseFloat(f, 32); err == nil {
    fmt.Println(s) // 3.1415927410125732
}
if s, err := strconv.ParseFloat(f, 64); err == nil {
    fmt.Println(s) // 3.14159265
}



*/
