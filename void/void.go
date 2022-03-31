package main

import (
    "github.com/galaxia-team/void/void/src/preprocessor"
    "github.com/galaxia-team/void/void/src/gitwrapper"
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/utils"
    "bufio"
    "fmt"
    "os"
)

func main() {
    a := os.Args[1:]
    if len(a) == 0 {
        fmt.Println("must specify a command\n")
        utils.PrintHelp()
        os.Exit(0)
    }
    if !utils.Contains(utils.Arguments, a[0]) {
        fmt.Println("invalid usage\n")
        utils.PrintHelp()
        os.Exit(0)
    }
    if a[0] == "run" && len(a) == 1 {
        exception.Except("file_not_specified", 0)
    }
    if a[0] == "help" {
        utils.PrintHelp()
        os.Exit(0)
    }
    if a[0] == "version" {
        utils.PrintVersion()
        os.Exit(0)
    }
    if a[0] == "update" {
        gitwrapper.Update()
        os.Exit(0)
    }

    fp := a[1]
    if fp[len(fp)-5:] != ".void" {
        exception.Except("file_not_void", 0)
    } else {
        _, serr := os.Stat(fp)
        if serr != nil {
            exception.Except("file_not_found", 0)
        }
    }
    var fd []string
    f, ferr := os.Open(fp)
    if ferr == nil {
        sc := bufio.NewScanner(f)
        sc.Split(bufio.ScanLines)
        for sc.Scan() {
            fd = append(fd, sc.Text())
        }
    } else {
        exception.Except("file_not_accessible", 0)
    }
    f.Close()
    preprocessor.PreProcess(fd)
}
