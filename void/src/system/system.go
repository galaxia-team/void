package system

import (
    "github.com/galaxia-team/void/void/src/exception"
    "strings"
    "os/exec"
    "os/user"
    "fmt"
)

func User(n int) string {
    u, e := user.Current()
    if e != nil {
        exception.Except("cant_fetch_user", n)
    }
	return u.Username
}

func ExecCommand(c []string, n int) []string {
    o, e := exec.Command(c...).Output()
    if e != nil {
        exception.Except("unexpected_error", n)
    }
    return strings.Split(fmt.Sprintf("%v", o), "\n")
}
