text_colors = {
    "black": 30,
    "red": 31,
    "green": 32,
    "yellow": 33,
    "blue": 34,
    "purple": 35,
    "cyan": 36,
    "white": 37,
}

bg_colors = {
    "black": 40,
    "red": 41,
    "green": 42,
    "yellow": 43,
    "blue": 44,
    "purple": 45,
    "cyan": 46,
    "white": 47,
}

styles = {
    "none": 0,
    "bold": 1,
    "underline": 2,
    "negative1": 3,
    "negative2": 5,
}

def style(style, *args):
    styled_args = []
    if len(args) > 1:
        for arg in args:
            styled_args.append(f"\033[{styles[style]}m{arg}")
        return styled_args
    else:
        return(f"\033[{styles[style]}m{args[0]}")

def text_color(text_color, *args):
    colored_args = []
    if len(args) > 1:
        for arg in args:
            colored_args.append(f"\033[{text_colors[text_color]}m{arg}")
        return colored_args
    else:
        return(f"\033[{text_colors[text_color]}m{args[0]}")

def bg_color(bg_color, *args):
    colored_args = []
    if len(args) > 1:
        for arg in args:
            colored_args.append(f"\033[{bg_colors[bg_color]}m{arg}")
        return colored_args
    else:
        return(f"\033[{bg_colors[bg_color]}m{args[0]}")