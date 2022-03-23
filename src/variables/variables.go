package variables

type Var struct {
    mutable bool
    name, scope, val string
}

var Vars = map[string]interface{} {}

func InitVar(scmd []string, vscope string) {
    Vars[scmd[1]] = Var{
        mutable: scmd[0] == "var",
        name: scmd[1],
        scope: vscope,
        val: "wip",
    }
}