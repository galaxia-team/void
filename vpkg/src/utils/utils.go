package utils

import (
    "github.com/go-git/go-git/v5"
    "runtime"
    "fmt"
)

var (
    Arguments = []string {"help", "version", "install", "uninstall", "update"}
    Help = []string {
        "usage:",
        "vpkg <argument> <package (optional)>",
        "",
        "arguments:",
        "help - print vpkg or vpkg package help screen",
        "version - print vpkg or vpkg package version",
        "install - install vpkg package",
        "uninstall - uninstall vpkg package",
        "update - update vpkg or vpkg package to latest version",
    }
    Commit = "N/A"
    RootDir = "./"
)

const (
    Version = "0.0.1"
    VersionName = "abyss"
    VersionState = "pre-alpha"
)

func GitClone() {
    return
}

func PrintVersion() {
    fmt.Printf("void version %s '%s' (%s) %s/%s commit %s\n", Version, VersionName, VersionState, runtime.GOOS, runtime.GOARCH, Commit)
}

func PrintHelp() {
    for _, h := range Help {
        fmt.Println(h)
    }
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
