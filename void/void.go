package main

import (
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/preprocessor"
    "github.com/galaxia-team/void/void/src/utils"
    "bufio"
    "fmt"
    "os"
)

func main() {
    a := os.Args[1:]
    if utils.Contains(a, "help") {
        fmt.Println(utils.GetHelp())
        os.Exit(0)
    }
    if utils.Contains(a, "version") {
        fmt.Println(utils.GetVersion())
        os.Exit(0)
    }
    fp := args[0]
    var fd []string
    if fp == "" {
        exception.Except("file_not_specified", 0)
    } else if fp[len(fp)-5:] != ".void" {
        exception.Except("file_not_void", 0)
    } else {
        _, serr := os.Stat(fp)
        if serr != nil {
            exception.Except("file_not_found", 0)
        }
    }
    file, ferr := os.Open(fp)
    if ferr == nil {
        sc := bufio.NewScanner(file)
        sc.Split(bufio.ScanLines)
        for sc.Scan() {
            fd = append(fd, sc.Text())
        }
    } else {
        exception.Except("file_not_accessible", 0)
    }
    preprocessor.PreProcess(fd)
}
