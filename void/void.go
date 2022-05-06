package main

import (
    "github.com/galaxia-team/void/void/src/preprocessor"
    "github.com/galaxia-team/void/void/src/gitwrapper"
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/system"
    "github.com/galaxia-team/void/void/src/utils"
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
        if !system.PathExists(fp) {
            exception.Except("file_not_found", 0)
        }
    }
    preprocessor.PreProcess(system.ReadFile(fp))
}
