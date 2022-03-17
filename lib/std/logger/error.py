from ..console.style import text_color

def error(*args):
    print(*text_color("red", "ERROR |", *args))