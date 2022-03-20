# imports
from sys import argv
from os import _exit, path
from importlib import import_module
import re
from ast import literal_eval
from time import time
import operator
from signal import signal, SIGINT

# functions
def get_scope(line_num, ccmd):
    if ccmd[-1] == "{":
        return f'{"".join(list(re.findall(".*{", ccmd)[0])[:-1]).strip().replace(" ", "_")}_{line_num}'
    else:
        current_line_num = line_num + 1
        current_line = file[line_num]

        while current_line and current_line[-1] != "{":
            current_line_num -= 1
            if current_line_num >= 0:
                current_line = file[current_line_num]
            else:
                return "root"
            if current_line[-1] == "}":
                while current_line and current_line[-1] != "{":
                    current_line_num -= 1
                    if current_line_num >= 0:
                        current_line = file[current_line_num]
                    else:
                        return "root"
        
        if current_line[-1] == "{":
            return f'{"".join(list(re.findall(".*{", current_line)[0])[:-1]).strip().replace(" ", "_")}_{current_line_num-1}'

def get_data(line_num, start_char, break_char, ccmd):
    if ccmd[-1] == break_char:
        return ["".join(list(re.sub(f".*{start_char}", "", ccmd))[:-1])]
    else:
        current_line_num = line_num
        current_line = ccmd
        data = [current_line]
        while current_line and current_line[-1] != break_char:
            current_line_num += 1
            current_line = file[current_line_num]
            data.append(current_line)
        data = data[1:-1]
    return data

def parse_string(unparsed_string, inter_var, **kw):
    parsed_string = unparsed_string
    
    if re.match('f\".*\"', parsed_string) or re.match("f\'.*\'", parsed_string):
        parsed_string = "".join(list(parsed_string)[1:])
        fstr_var = re.findall("{.*}", parsed_string)
        for fstr in fstr_var:
            val = inter_var[fstr.replace("{", "", 1).replace("}", "", 1)]["val"]
            parsed_string = parsed_string.replace(fstr, val)     
    elif re.match('\".*\"', parsed_string) or re.match("\'.*\'", parsed_string):
        parsed_string = "".join(list(parsed_string)[1:-1])
    
    return parsed_string

def parse_list(unparsed_list, inter_var, scope, ccmd, **kw):
    if unparsed_list == "()" or unparsed_list == "[]":
        return ()

    parsed_list = unparsed_list

    if "," not in parsed_list:
        return parse_expr("".join(list(parsed_list)[1:-1]), inter_var, scope, ccmd, args=kw.get("args"))
        
    if parsed_list[0] == "[" and parsed_list[-1] == "]":
        parsed_list = list(parsed_list)
        parsed_list[0] = "("
        parsed_list[-1] = ")"
        parsed_list = "".join(parsed_list)

    if scope != "root":
        parsing_var = { **var["global"], **var["local"][scope] }
        scopes[scope]["temp"]["arg_strs"] = [ *re.findall('\".*\",', parsed_list), *re.findall("\'.*\',", parsed_list) ]
        for num, string in enumerate(scopes[scope]["temp"]["arg_strs"]):
            parsed_list = parsed_list.replace(string, f'"str{num}",')
    else:
        parsing_var = var["global"]
        root["temp"]["arg_strs"] = [ *re.findall('\".*\,"', parsed_list), *re.findall("\'.*\',", parsed_list) ]
        for num, string in enumerate(root["temp"]["arg_strs"]):
            parsed_list = parsed_list.replace(string, f'"str{num}",')
    
    for parse_var in parsing_var:
        val = parsing_var[parse_var]["val"]
        if f", {parse_var})" in parsed_list or f",{parse_var})" in parsed_list:
            parsed_list = parsed_list.replace(f", {parse_var})", f", {val})").replace(f",{parse_var})", f", {val})")
        if f"({parse_var}," in parsed_list:
            parsed_list = parsed_list.replace(f"({parse_var}, ", f"({val}, ")
        if parsed_list == f"({parse_var})":
            return (val)
        while f", {parse_var}," in parsed_list or f",{parse_var}," in parsed_list:
            parsed_list = parsed_list.replace(f", {parse_var},", f", {val},").replace(f",{parse_var},", f",{val},")

    if scope != "root":
        scopes[scope]["temp"]["arg_fstrs"] = [ *re.findall('f\".*\"', parsed_list), *re.findall("f\'.*\'", parsed_list) ]
        for fstring in scopes[scope]["temp"]["arg_fstrs"]:
            fstring_var = re.findall("{.*}", fstring)
            for fstring_var in fstring_var:
                val = parsing_var[fstring_var.replace("{", "").replace("}", "")]["val"]
                parsed_list = parsed_list.replace(fstring, fstring.replace(fstring_var, val).replace("f", "", 1))
        for num, string in enumerate(scopes[scope]["temp"]["arg_strs"]):
            parsed_list = parsed_list.replace(f'"str{num}",', string)

    else:
        root["temp"]["arg_fstrs"] = [ *re.findall('f\".*\"', parsed_list), *re.findall("f\'.*\'", parsed_list) ]
        for fstring in root["temp"]["arg_fstrs"]:
            fstring_var = re.findall("{.*}", fstring)
            for fstring_var in fstring_var:
                val = parsing_var[fstring_var.replace("{", "").replace("}", "")]["val"]
                parsed_list = parsed_list.replace(fstring, fstring.replace(fstring_var, val).replace("f", "", 1))
        for num, string in enumerate(root["temp"]["arg_strs"]):
            parsed_list = parsed_list.replace(f'"str{num}",', string)
    
    parsed_list = literal_eval(parsed_list)

    return parsed_list

def init_var(line_num, var_mode, ccmd, inter_var, scope):
    var_name = re.sub("=.*", "", ccmd.replace(f"{var_mode} ", "")).strip()
    var_val = parse_expr(re.sub(f"{var_mode} {var_name}.*=", "", ccmd).lstrip(), inter_var, scope, ccmd)

    if scope == "root":
        var["global"][var_name] = {
            "mutable": var_mode == "var",
            "type": type(var_val),
            "scope": "root",
            "val": var_val
        }
    else:
        var["local"][scope][var_name] = {
            "mutable": var_mode == "var",
            "type": type(var_val),
            "scope": scope,
            "val": var_val
        }

def interpret_func(ccmd, inter_var, scope, **kw):
    full_func = re.findall("[^\s]+\(.*\)", ccmd)[0]
    func_name = full_func.replace("(", "")
    args = merge_var_dicts(parse_list(re.findall("\(.*\)", full_func)[0], inter_var, scope, ccmd), kw.get("args"))
    func_name = ccmd.replace(re.findall("\(.*\)", full_func)[0], "")
    if funcs[func_name]["type"] == "included":
        if isinstance(args, tuple):
            return funcs[func_name]["code"](*args)
        else:
            return funcs[func_name]["code"](args)
    else:
        return interpret(funcs[func_name]["code"], args=args)

def modify_var(line_num, scmd, inter_var, scope):
    operator = scmd[1][:-1]
    if scope != "root" and scmd[0] in var["local"][scope]:
        if var["local"][scope][scmd[0]]["mutable"]:
            if operator:
                var["local"][scope][scmd[0]]["val"] = ops[operator](var["local"][scope][scmd[0]]["val"], int(scmd[2]))
            else:
                var["local"][scope][scmd[0]]["val"] = scmd[2]
        else:
            void_error("const_immutable", arg_line_num=line_num)
    else:
        if var["global"][scmd[0]]["mutable"]:
            if operator:
                var["global"][scmd[0]]["val"] = ops[operator](var["global"][scmd[0]]["val"], int(scmd[2]))
            else:
                var["global"][scmd[0]]["val"] = scmd[2]
        else:
            void_error("const_immutable", arg_line_num=line_num)

def modify_seq(line_num, ccmd, inter_var, scope):
    seq = re.findall(".+?(?=\[)", ccmd)[0]
    locs = re.findall("(?<=\[)(.*?)(?=\])", re.findall(".+?(?=\=)", ccmd))
    new_val = parse_expr(re.sub(f"{seq}.*=", "", ccmd).lstrip(), inter_var, scope, ccmd)
    if scope != "root" and seq in var["local"][scope]:
        if var["local"][scope][seq]["type"] == list:
            try:
                var["local"][scope][seq]["val"][locs] = new_val
            except:
                void_error("loc_not_int", arg_line_num=line_num)
        else:
            var["local"][scope][seq]["val"][locs] = new_val
    else:
        if var["global"][seq]["type"] == list:
            try:
                for count, loc in enumerate(locs):
                    locs[count] = int(loc)
                var["global"][seq]["val"][locs] = new_val
            except:
                void_error("loc_not_int", arg_line_num=line_num)
        else:
            var["global"][seq]["val"][locs] = new_val

def handle_op(line_num, op, ccmd, scmd, global_var, scope_var):
    final_op = scmd
    for c, i in enumerate(final_op):
        if i in ops:
            print("wip")
        if i in scope_var:
            final_op[c] = scope_var[i]["val"]
    return

def exit_void(*args):
    global runtime
    global void_running
    void_running = False
    if "-t" in a:
        if len(args) > 0:
            print("\n")
        runtime = time() * 1000 - runtime
        print(f"time to setup: {round(setuptime, 5)}ms")
        print(f"time to run: {round(runtime, 5)}ms")
    _exit(0)

def void_error(err, **kw):
    if err in no_line_errors:
        print("ERROR |", no_line_errors[err])
    else:
        print(f"ERROR (line {kw.get('arg_line_num')+file_offset}) |", errors[err])
    _exit(1)

def merge_var_dicts(dict1, dict2):
    if dict2 not in empty:
        merged_dict = dict1
        for d2var in dict2:
            merged_dict[d2var] = dict2[d2var]["val"]
        return merged_dict
    else:
        return dict1

def get_var(scope):
    if scope != "root":
        merge_var_dicts(var["global"], var["local"][scope])
    else:
        return var["global"]

def parse_expr(expr, inter_var, scope, ccmd, **kw):
    args = kw.get("args")
    try:
        return int(expr)
    except:
        try:
            return float(expr)
        except:
            if inter_var not in empty and expr in inter_var:
                val = inter_var[expr]["val"]
                try:
                    return int(val)
                except:
                    try:
                        return float(val)
                    except:
                        return val
            elif args not in empty and expr in args:
                return args[expr]["val"]
            elif expr[0] == "[":
                return get_data(line_num, "[", "]", ccmd)
            elif expr[0] == "{":
                data = get_data(line_num, "{", "}", ccmd)
                final_dict = {}
                for pair in data:
                    split_pair = pair.split(": ")
                    final_dict[split_pair[0]] = split_pair[1]
                return final_dict
            elif re.match('\".*\"', expr) or re.match("\'.*\'", expr):
                return parse_string(expr, scope)
            else:
                interpret_func(expr, inter_var, scope)

# interpret
def interpret(data, **kw):
    for data_line_num, line in enumerate(data):
        line_num = file.index(data[data_line_num])
        returned = ""
        ccmd = line.lstrip()
        scmd = ccmd.split()
        scope = get_scope(line_num, ccmd)
        inter_var = get_var(scope)
        args = kw.get("args")
        full_var = merge_var_dicts(inter_var, args)
        if ccmd == "pass":
            pass
        elif scmd[0] == "var":
            returned = init_var(line_num, "var", ccmd, inter_var, scope) 
        elif scmd[0] == "const":
            returned = init_var(line_num, "const", ccmd, inter_var, scope)
        elif ccmd == "break":
            returned = "break"
        elif ccmd == "continue":
            returned = "continue"
        elif len(scmd) > 1 and scmd[1][-1] == "=" and "[" not in scmd[0]:
            returned = modify_var(line_num, scmd, inter_var, scope)
        elif ccmd == "forever {":
            forever_data = get_data(line_num, "{", "}", ccmd)
            while void_running:
                forever_returned = interpret(forever_data, args=kw.get("args"))
                if forever_returned == "break":
                    return
                elif forever_returned == "continue":
                    continue
        elif scmd[0] == "repeat" and scmd[-1] == "{":
            repeat_data = get_data(line_num, "{", "}", ccmd)
            repeat_times = parse_expr("".join(scmd[1:-1]), inter_var, scope, ccmd)
            for _ in range(repeat_times-1):
                repeat_returned = interpret(repeat_data, args=kw.get("args"))
                if repeat_returned == "break":
                    return
                elif repeat_returned == "continue":
                    continue
        elif "[" in ccmd and "]" in ccmd and re.findall(".+?(?=\[)", ccmd)[0] in inter_var:
            returned = modify_seq(line_num, ccmd, inter_var, scope)
        elif ccmd[-1] == ")" and re.findall("[^\s]+\(.*\)", ccmd)[0] == ccmd:
            returned = interpret_func(ccmd, inter_var, scope, args=kw.get("args"))
        if returned:
            return returned

# variables
keywords = [
    "if",
    "elif",
    "else"
]
libs = {}
funcs = {}
scopes = {}
runtime = 0
setuptime = 0
void_running = False
no_line_errors = {
    "file_not_specified": "a file must be specified with the -f flag.",
    "file_not_found": "the specified file cannot be found.",
    "file_not_void": "the specified file must have the .void file extension.",
}
errors = {
    "cannot_concat_int": "integers and strings cannot be concatenated.",
    "const_immutable": "variables defined with const are immutable; they cannot be modified.",
    "loc_not_int": "if changing an item in a list, the index must be an integer."
}
root = {
    "temp": {}
}
var = {
    "global": {},
    "local": {},
}
ops = {
    "+": operator.add,
    "-": operator.sub,
    "*": operator.mul,
    "/": operator.truediv,
    "%": operator.mod,
    "//": operator.floordiv,
    "**": operator.pow,
    "&": operator.and_,
    "|": operator.or_,
    "^": operator.xor,
    ">>": operator.rshift,
    "<<": operator.lshift,
}
break_chars = ("}", "]", ")")
start_chars = ("{", "[", "(")
empty = (None, [], (), {}, "")
file_offset = 1

# process args
a = []
if argv:
    a = argv
else:
    void_error("file_not_specified")

if "-f" in a:
    file_loc = a[a.index("-f")+1]
    if file_loc[-5:] != ".void":
        void_error("file_not_void")
    if not path.isfile(file_loc):
        void_error("file_not_found")
    with open(file_loc, "r") as file_raw:
        file = file_raw.readlines()
        old_file = file
else:
    void_error("file_not_specified")

if "-t" in a:
    setuptime = time() * 1000

# clean up step 1 + handle imports
for line_num, line in enumerate(file):
    file[line_num] = line.rstrip()
    new_line = line.strip()
    if "include " in new_line:
        to_import = new_line.replace('include ', "").replace('"', "").replace("'", "").split(",")
        for item in to_import:
            s_item = item.strip()
            libs[s_item] = import_module(f"lib.{s_item}", package=None)
            funcs[s_item] = {
                "type": "included",
                "code": getattr(libs[s_item], s_item.split(".")[-1])
            }
        file[line_num] = ""
    elif new_line[0:2] == "//":
        file[line_num] = ""
    elif new_line[0:2] == "/*":
        current_line_num = line_num
        current_line = line
        while current_line[-2:] != "*/":
            current_line_num += 1
            current_line = file[current_line_num]
            file[current_line_num] = ""

# clean up step 2
while "" in file:
    file_offset += 1
    file.remove("")

# process funcs and scopes
for line_num, line in enumerate(file):
    s_line = line.lstrip()
    scope = get_scope(line_num, s_line)
    if scope != "root" and scope not in var["local"] and scope not in scopes:
        scopes[scope] = { "temp": {} }
        var["local"][scope] = {}
    if s_line[0:5] == "func ":
        func_name = re.sub("\(.*", "", s_line.replace("func ", ""))
        func_data = get_data(line_num, "{", "}", s_line)
        funcs[func_name] = {
            "type": "defined",
            "code": func_data
        }

if "-t" in a:
    setuptime = time() * 1000 - setuptime
    runtime = time() * 1000

# sigint handler
signal(SIGINT, exit_void)

# actually interpret
void_running = True
interpret(file)

# time
exit_void()