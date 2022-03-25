package utils

import (
    "fmt"
    "runtime"
)

func GetVersion() string {
    return fmt.Sprintf("void version %s '%s' (%s) %s/%s", Version, VersionName, VersionState, runtime.GOOS, runtime.GOARCH)
}

func GetHelp() string {
    return fmt.Sprintf("usage:\nvoid <argument or file>\n\narguments:\nhelp - print this help screen\nversion - print void version")
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

var (
    GlobalData = map[string]string {}
    LocalData = map[string]map[string]string {}
)

const (
    Version = "0.0.1"
    VersionName = "abyss"
    VersionState = "pre-alpha"
)