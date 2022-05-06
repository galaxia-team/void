package system

import (
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/utils"
    "strings"
    "os/exec"
    "os/user"
    "bufio"
    "fmt"
    "os"
)

func PathExists(p string) bool {
    _, se := os.Stat(p)
    if se != nil {
        return false
    }
    return true
}

func ReadFile(p string) []string {
    var fd []string
    f, fe := os.Open(p)
    if fe == nil {
        sc := bufio.NewScanner(f)
        sc.Split(bufio.ScanLines)
        for sc.Scan() {
            fd = append(fd, sc.Text())
        }
    } else {
        exception.Except("file_not_accessible", 0)
    }
    f.Close()
    return fd
}

func User(n int) string {
    u, e := user.Current()
    if e != nil {
        exception.Except("cant_fetch_user", n)
    }
	return u.Username
}

func ExecCommand(c string, n int) []string {
    sc := strings.Fields(c)
    sc0 := sc[0]
    sc = utils.RemoveIndex(sc, 0)
    o, e := exec.Command(sc0, sc...).Output()
    if e != nil {
        exception.Except("unexpected_error", n)
    }
    return strings.Split(fmt.Sprintf("%v", o), "\n")
}
