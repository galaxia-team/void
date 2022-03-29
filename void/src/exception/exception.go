package exception

import (
    "fmt"
    "os"
)

var (
    NoLineExceptions = map[string]string {
        "file_not_specified": "a file must be specified if using the 'run' flag.",
        "file_not_found": "the specified file cannot be found.",
        "file_not_accessible": "the specified file cannot be accessed.",
        "file_not_void": "the specified file must have the .void file extension.",
        "incorrect_install": "void cannot be updated, as it has not been installed correctly.",
        "git_unreachable": "void cannot be updated, as its git repo cannot be reached.",
        "unexpected_error": "an unexpected and inexplainable error has occurred. please report this to the developers.", 
    }

    Exceptions = map[string]string {
        "str_add_int": "cannot add int to string.",
        "str_add_float": "cannot add int to float.",
        "not_bool": "boolean argument not a boolean.",
        "invalid_op": "invalid operation.",
        "invalid_type": "invalid type.",
        "const_immutable": "variables defined with const are immutable; they cannot be modified.",
        "invalid_syntax": "invalid syntax.",
        "loc_not_int": "if changing an item in a list, the index must be an integer.",
    }
)

func Except(e string, n int) {
    if _, ok := NoLineExceptions[e]; ok {
        fmt.Println(fmt.Sprintf("EXCEPTION | %s", NoLineExceptions[e]))
    } else {
        fmt.Println(fmt.Sprintf("EXCEPTION (line %d) | %s", n, Exceptions[e]))
    }
    os.Exit(1)
}
