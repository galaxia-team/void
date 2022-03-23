package console

import "fmt"

func Print(m ...interface{}) {
    fmt.Print(m...)
}

func Printf(m ...interface{}) {
    fmt.Printf(m...)
}

func Println(m ...interface{}) {
    fmt.Println(m...)
}