from ..console.style import text_color

def notice(*args):
    print(*text_color("blue", "NOTICE |", *args))