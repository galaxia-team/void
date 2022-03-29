package utils

import (
    "runtime"
    "fmt"
)

type Var struct {
    mutable bool
    name, scope, val string
}

type Func struct {
    name string
    data func(map[string]interface{}) interface{}
}

var (
    GlobalData = map[string]string {}
    LocalData = map[string]map[string]string {}
    GlobalVars = map[string]Var {}
    LocalVars = map[string]map[string]Var {}
    GlobalFuncs = map[string]map[string]Func {}
    LocalFuncs = map[string]map[string]map[string]Func {}
    Arguments = []string {
        "help",
        "version",
        "run",
        "update",
    }
    Help = []string {
        "usage:",
        "void <argument> <file (optional)>",
        "",
        "arguments:",
        "help - print this help screen",
        "version - print void version",
        "run - run specified void file",
        "update - update void to the latest version.",
    }
    Commit = "N/A"
    RootDir = "./"
)

const (
    Version = "0.0.1"
    VersionName = "abyss"
    VersionState = "pre-alpha"
)

func PrintVersion() {
    fmt.Printf("void version %s '%s' (%s) %s/%s commit %s\n", Version, VersionName, VersionState, runtime.GOOS, runtime.GOARCH, Commit)
    return
}

func PrintHelp() {
    for _, h := range Help {
        fmt.Println(h)
    }
    return
}

func Extend(s1, s2 []string) []string {
    return append(s1, s2...)
}

func RemoveIndex(s []string, i int) []string {
    return append(s[:i], s[i+1:]...)
}

func Remove(s []string, r interface{}) []string {
    i := GetIndex(s, r)
    return append(s[:i], s[i+1:]...)
}

func GetIndex(s []string, v interface{}) int {
    for n, cv := range s {
        if cv == v {
            return n
        }
    }
    return -1
}

func Contains(s []string, v interface{}) bool {
    for _, cv := range s {
        if cv == v {
            return true
        }
    }
    return false
}
