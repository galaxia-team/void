package gitwrapper

import (
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/utils"
    "github.com/go-git/go-git/v5"
    "fmt"
    "os"
)

func GitClone(u string, d string) (interface{}, error) {
	gd, e := git.PlainClone(d, false, &git.CloneOptions{
        URL:               u,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
    return gd, e
}

func GetCloneData(d interface{}) (interface{}, interface{}) {
    b, be := d.Head()
    c, ce := d.CommitObject(b.Hash())
    return b, c, be, ce
}

func Update() {
    if utils.RootDir == "./" && utils.Commit == "N/A" {
        exception.Except("incorrect_install", 0)
    }
    fmt.Println("updating...\n")
    d, err := GitClone("https://github.com/galaxia-team/void.git", utils.RootDir)
    if err != nil {
        b, c, be, ce = GetCloneData(d)
        if be != nil || ce != nil {
            exception.Except("unexpected_error", 0)
        }
    } else {
        exception.Except("git_unreachable", 0)
    }
}