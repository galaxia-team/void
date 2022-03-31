package gitwrapper

import (
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/system"
    "github.com/galaxia-team/void/void/src/utils"
    "github.com/go-git/go-git/v5"
    "strings"
    "fmt"
    "os"
)

func Clone(u string, d string) (*git.Repository, error) {
	gd, e := git.PlainClone(d, false, &git.CloneOptions{
        URL: u,
        Progress: os.Stdout,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
    return gd, e
}

func CommitHashString(d *git.Repository) string {
    b, e := d.Head()
    if e != nil {
        exception.Except("unexpected_error", 0)
    }
    return strings.Fields(fmt.Sprintf("%v", b))[0]
}

func Update() {
    var c string
    if system.User(0) != "root" {
        exception.Except("not_root", 0)
    }
    /*
    if utils.RootDir == "./" || utils.Commit == "N/A" {
        exception.Except("incorrect_install", 0)
    }
    */
    fmt.Println("cloning repo...\n")
    d, err := Clone("https://github.com/galaxia-team/void.git", utils.RootDir+"voidupdate")
    if err == nil {
        c = CommitHashString(d)
        bc := []string {
            "go",
            "build",
            "-ldflags",
            `"-X`,
            "'utils.Commit="+c+"'",
            "-X",
            `'utils.RootDir=`+utils.RootDir+`'"`,
            "-v",
        }
        os.Chdir(utils.RootDir+"voidupdate/void")
        fmt.Println("\nbuilding...\n")
        o := system.ExecCommand(bc, 0)
        fmt.Println(o)
    } else {
        exception.Except("git_unreachable", 0)
    }
}
