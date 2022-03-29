package test

import (
    "github.com/galaxia-team/void/src/lib"
)

func main() {
    c := lib.ImportLib("std.console")
    c.Println("Hello, world!")
}