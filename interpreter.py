# imports
from sys import argv
import os
import importlib
import re
from ast import literal_eval
from time import time
import operator
from signal import signal, SIGINT
from lib.std.console.style import text_color

# functions
def get_scope(line_num, ccmd):
    if ccmd[-1] == "{":
        scope = list(re.findall(".*{", ccmd)[0])
        del scope[-1]
        return f'{"".join(scope).strip().replace(" ", "_")}_{line_num}'
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
            scope = list(re.findall(".*{", current_line)[0])
            del scope[-1]
            return f'{"".join(scope).strip().replace(" ", "_")}_{current_line_num-1}'

def get_full_multiline_data(line_num, break_char, ccmd):
    if ccmd[-1] == break_char:
        data = list(re.sub(".*{", "", ccmd))
        del data[0]
        return ["".join(data)]
    else:
        current_line_num = line_num
        current_line = ccmd
        data = [current_line]
        while current_line and current_line[-1] != break_char:
            current_line_num += 1
            current_line = file[current_line_num]
            data.append(current_line)
        del data[0]
        del data[-1]
    return data

def parse_input(unparsed_input, scope):
    parsed_input = unparsed_input
    
    if re.match('f\".*\"', parsed_input) or re.match("f\'.*\'", parsed_input):
        parsed_input = list(parsed_input)
        del parsed_input[0]
        parsed_input = "".join(parsed_input)
        fstring_vars = re.findall("{.*}", parsed_input)
        if scope != "root":
            parsing_variables = { **variables["globals"], **variables["locals"][scope] }
        else:
            parsing_variables = variables["globals"]
        for fstring_var in fstring_vars:
            val = parsing_variables[fstring_var.replace("{", "").replace("}", "")]["val"]
            parsed_input = parsed_input.replace(fstring_var, val)     
    elif re.match('\".*\"', parsed_input) or re.match("\'.*\'", parsed_input):
        parsed_input = list(parsed_input)
        del parsed_input[0]
        del parsed_input[-1]
        parsed_input = "".join(parsed_input)
    
    return parsed_input

def parse_list(unparsed_list, scope):
    if unparsed_list == "()":
        return []

    parsed_list = unparsed_list
    if scope != "root":
        parsing_variables = { **variables["globals"], **variables["locals"][scope] }
        scopes[scope]["arg_processing_tmp"]["strings"] = [ *re.findall('\".*\",', parsed_list), *re.findall("\'.*\',", parsed_list) ]
        for num, string in enumerate(scopes[scope]["arg_processing_tmp"]["strings"]):
            parsed_list = parsed_list.replace(string, f'"str{num}",')
    else:
        parsing_variables = variables["globals"]
        root_data["arg_processing_tmp"]["strings"] = [ *re.findall('\".*\,"', parsed_list), *re.findall("\'.*\',", parsed_list) ]
        for num, string in enumerate(root_data["arg_processing_tmp"]["strings"]):
            parsed_list = parsed_list.replace(string, f'"str{num}",')
    
    for variable in parsing_variables:
        val = parsing_variables[variable]["val"]
        if f", {variable})" in parsed_list or f",{variable})" in parsed_list:
            parsed_list = parsed_list.replace(f", {variable})", f", {val})").replace(f",{variable})", f", {val})")
        if f"({variable}," in parsed_list:
            parsed_list = parsed_list.replace(f"({variable}, ", f"({val}, ")
        if parsed_list == f"({variable})":
            return [f"{val}"]
        while f", {variable}," in parsed_list or f",{variable}," in parsed_list:
            parsed_list = parsed_list.replace(f", {variable},", f", {val},").replace(f",{variable},", f",{val},")

    if scope != "root":
        scopes[scope]["arg_processing_tmp"]["fstrings"] = [ *re.findall('f\".*\"', parsed_list), *re.findall("f\'.*\'", parsed_list) ]
        for fstring in scopes[scope]["arg_processing_tmp"]["fstrings"]:
            fstring_vars = re.findall("{.*}", fstring)
            for fstring_var in fstring_vars:
                val = parsing_variables[fstring_var.replace("{", "").replace("}", "")]["val"]
                parsed_list = parsed_list.replace(fstring, fstring.replace(fstring_var, val).replace("f", "", 1))
        for num, string in enumerate(scopes[scope]["arg_processing_tmp"]["strings"]):
            parsed_list = parsed_list.replace(f"'str{num}',", string)

    else:
        root_data["arg_processing_tmp"]["fstrings"] = [ *re.findall('f\".*\"', parsed_list), *re.findall("f\'.*\'", parsed_list) ]
        for fstring in root_data["arg_processing_tmp"]["fstrings"]:
            fstring_vars = re.findall("{.*}", fstring)
            for fstring_var in fstring_vars:
                val = parsing_variables[fstring_var.replace("{", "").replace("}", "")]["val"]
                parsed_list = parsed_list.replace(fstring, fstring.replace(fstring_var, val).replace("f", "", 1))
        for num, string in enumerate(root_data["arg_processing_tmp"]["strings"]):
            parsed_list = parsed_list.replace(f"'str{num}',", string)
    
    parsed_list = literal_eval(parsed_list)
    
    if not isinstance(parsed_list, tuple):
        parsed_list = [parsed_list]

    return parsed_list

def initialize_variable(line_num, var_type, ccmd, scope):
    variable_name = re.sub("=.*", "", ccmd.replace(f"{var_type} ", "")).strip()
    variable_value = parse_input(re.sub(f"{var_type} {variable_name}.*=", "", ccmd).strip(), scope)

    if variable_value[0] == "[":
        variable_value = get_full_multiline_data(line_num, "]", ccmd)

    elif variable_value[0] == "{":
        data = get_full_multiline_data(line_num, "}", ccmd)
        variable_value = {}
        for pair in data:
            split_pair = pair.split(": ")
            variable_value[split_pair[0]] = split_pair[1]

    if scope == "root":
        variables["globals"][variable_name] = {
            "type": var_type,
            "val": variable_value
        }
    else:
        variables["locals"][scope][variable_name] = {
            "type": var_type,
            "val": variable_value
        }

def interpret_function(ccmd, scope):
    full_func = re.findall("[^\s]+\(.*\)", ccmd)[0]
    func_name = full_func.replace("(", "")
    args = parse_list(re.findall("\(.*\)", full_func)[0], scope)
    func_name = ccmd.replace(re.findall("\(.*\)", full_func)[0], "")
    if functions[func_name]["type"] == "included":
        return functions[func_name]["code"](*args)
    else:
        return interpret(functions[func_name]["code"], args)

def reassign_variable(line_num, scmd, global_vars, scope_vars, scope):
    global variables
    operator = scmd[1][0]
    if operator != "=":
        operator
    else:     
        if scmd[0] in scope_vars:
            variables["locals"][scope][scmd[0]] = scmd[2]
        else:
            variables["globals"][scmd[0]] = scmd[2]

def interpret_operation(line_num, operator, ccmd, scmd, global_vars, scope_vars):
    final_operation = scmd
    for c, i in enumerate(final_operation):
        if i in operators:
            print("wip")
        if i in scope_vars:
            final_operation[c] = scope_vars[i]["val"]
    return

def close_interpreter(*args):
    global runtime
    global interpreter_running
    interpreter_running = False
    if "-t" in a:
        if len(args) > 0:
            print("\n")
        runtime = time() * 1000 - runtime
        print(f"time to setup: {round(setuptime, 4)}ms")
        print(f"time to run: {round(runtime, 4)}ms")
    os._exit(0)

# interpret
def interpret(data, args, **kw):
    for data_line_num, line in enumerate(data):
        line_num = file.index(data[data_line_num])
        returned = ""
        ccmd = line.strip()
        scmd = ccmd.split()
        scope = get_scope(line_num, ccmd)
        global_vars = variables["globals"]
        if scope != "root":
            scope_vars = variables["locals"][scope]
        else:
            scope_vars = global_vars
        if operators in scmd:
            returned = interpret_operation(line_num, scmd[1], ccmd, scmd, global_vars, scope_vars)
        elif scmd[0] == "let":
            returned = initialize_variable(line_num, "let", ccmd, scope)
        elif scmd[0] == "const":
            returned = initialize_variable(line_num, "const", ccmd, scope)
        elif ccmd == "break" and "scope" in kw:
            returned = f"{kw.get('scope')}_break"
        elif len(scmd) > 1 and scmd[1][-1] == "=":
            returned = reassign_variable(line_num, scmd, global_vars, scope_vars, scope)
        elif ccmd == "forever {":
            forever_data = get_full_multiline_data(line_num, "}", ccmd)
            while interpreter_running:
                breakb = interpret(forever_data, {}, scope=scope)
                if breakb == f"{scope}_break":
                    return
        elif ccmd[-1] == ")" and re.findall("[^\s]+\(.*\)", ccmd)[0] == ccmd:
            returned = interpret_function(ccmd, scope)
        if returned:
            return returned

# variables
keywords = [
    "if",
    "elif",
    "else"
]
imported_libraries = {}
functions = {}
scopes = {}
runtime = 0
setuptime = 0
interpreter_running = False
root_data = {
    "arg_processing_tmp": {
        "strings": []
    }
}
variables = {
    "globals": {},
    "locals": {},
}
operators = ("+", "-", "/", "*", "@", "**", "%")

# process args
a = []
if argv:
    a = argv
else:
    print("must specify a file.")
    os._exit(1)

if "-f" in a:
    with open(a[a.index("-f")+1], "r") as file_raw:
        file = file_raw.readlines()
        old_file = file
else:
    print("must specify a file.")
    os._exit(1)

if "-t" in a:
    setuptime = time() * 1000

# clean up step 1 + handle imports
for line_num, line in enumerate(file):
    file[line_num] = line.rstrip()
    new_line = file[line_num]
    if 'include "' in new_line:
        to_import = new_line.replace('include "', "").replace('"', "").split(",")
        for item in to_import:
            imported_libraries[item.strip()] = importlib.import_module(f"lib.{item.strip()}", package=None)
            functions[item.strip()] = {
                "type": "included",
                "code": getattr(imported_libraries[item.strip()], item.strip().split(".")[-1])
            }
        file[line_num] = ""
    elif ' //' in new_line or new_line.strip()[0:2] == "//":
        file[line_num] = re.sub("//.*", "", new_line)
        if new_line.strip() == "":
            file[line_num] = ""

# clean up step 2
file = list(filter(("").__ne__, file))

# process functions and scopes
for line_num, line in enumerate(file):
    scope = get_scope(line_num, line)
    if scope != "root" and scope not in variables["locals"] and scope not in scopes:
        scopes[scope] = {
            "line_num": line_num,
            "arg_processing_tmp": {"strings": []},
        }
        variables["locals"][scope] = {}
    if line.strip()[0:9] == "function ":
        func_name = line.strip().replace("function ", "").replace("(", "").replace("", "").strip()
        functions[func_name] = {
            "type": "defined",
            "code": get_full_multiline_data(line_num, "}", line)
        }

if "-t" in a:
    setuptime = time() * 1000 - setuptime
    runtime = time() * 1000

# sigint handler
signal(SIGINT, close_interpreter)

# actually interpret
interpreter_running = True
interpret(file, {})

# time
close_interpreter()