package utils

import (
    "runtime"
    "fmt"
)

var (
    Commit = "N/A"
    RootDir = "./"
)

const (
    Version = "0.0.1"
    VersionName = "abyss"
    VersionState = "pre-alpha"
)

func GetVersion() string {
    return fmt.Sprintf("vpkg version %s '%s' (%s) %s/%s commit %s", Version, VersionName, VersionState, runtime.GOOS, runtime.GOARCH, Commit)
}

func GetHelp() string {
    return fmt.Sprintf("usage:\nvpkg <argument> <package>\n\narguments:\nhelp - print this help screen\nversion - print vpkg version\ninstall - install a vpkg package\nuninstall - uninstall a vpkg package\nupdate - update installed vpkg packages and package lists")
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
