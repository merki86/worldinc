title = "Title"
buttons = (("Q", "Quit", 4), ("R", "Reset", 5))
h = "━"
space = " "
width = 40  # ширина всей панели

# Построение кнопочной строки
button_strs = [f"[{b[0]}] {b[1]}" for b in buttons]
buttons_text = "  ".join(button_strs)
padding = width - len(title) - len(buttons_text)

# Если места мало — обрежем
if padding < 1:
    buttons_text = buttons_text[:width - len(title) - 1]
    padding = 1

legend = (
    f"┏{h * width}┓\n"
    f"┃{title}{space * padding}{buttons_text}┃\n"
    f"┗{h * width}┛"
)

print(legend)