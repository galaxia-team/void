from ..console.style import text_color

def warn(*args):
    print(*text_color("yellow", "WARN |", *args))