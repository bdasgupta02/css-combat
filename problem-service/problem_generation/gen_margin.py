from random import randint
from .div.div import Div
from .div.outer_div import OuterDiv
from .gen_color import generate_pair

pair = generate_pair()
max_level = randint(3, 5)
st_h = "height: {}%".format(randint(60, 80))
st_w = "width: {}%".format(randint(60, 80))

margin_prefs = [
    "margin-left:",
    "margin-top:",
    "margin-right:",
    "margin-bottom:"
]

pad_prefs = [
    "padding-left:",
    "padding-top:",
    "padding-right:",
    "padding-bottom:"
]

center_style = [
    "display: flex",
    "align-items: center",
    "justify-content: center"
]

def generate_margin_div(level):
    if level == max_level:
        return Div([], "")

    bg_col = pair[0] if level % 2 == 0 else pair[1]
    bg_style = "background-color: {}".format(bg_col)
    h = "height: 100%" if level == 0 else st_h
    w = "width: 100%" if level == 0 else st_w

    margin_pref_idx = randint(0, 3)
    margin_pref = margin_prefs[margin_pref_idx]
    margin = "" if level == 0 else "{} {}px".format(margin_pref, randint(20, 40))

    pad_pref_idx = randint(0, 3)
    pad_pref = pad_prefs[pad_pref_idx]
    padding = "" if level == 0 else "{} {}px".format(pad_pref, randint(20, 40))

    child = generate_margin_div(level + 1)
    div = Div([
        h,
        w,
        bg_style,
        margin,
        padding,
        "overflow: hidden"
    ] + center_style, "")
    print(div.generate())
    div.children = child
    return div

def generate_margin():
    div = generate_margin_div(0)
    outer = OuterDiv(div)
    outer.render()