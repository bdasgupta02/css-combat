from random import randint
from .div.div import Div
from .div.outer_div import OuterDiv
from .gen_color import generate_pair

max_level = 3
st_h = "height: {}%".format(randint(60, 90))
st_w = "width: {}%".format(randint(60, 90))
center_style = [
    "display: flex",
    "align-items: center",
    "justify-content: center"
]


def generate_flexbox_rec(level, is_vert, is_reverse):
    pair = generate_pair()
    if level == 3:
        return Div([])

    bg_col_i = 0 if level % 2 == 0 else 1
    if is_reverse and bg_col_i == 1:
        bg_col_i = 0
    elif is_reverse:
        bg_col_i = 1

    bg_col = pair[bg_col_i]
    bg_style = "" if level != 0 and randint(
        1, 4) < 2 else "background-color: {}".format(bg_col)

    flex_dir = "flex-direction: column" if randint(
        1, 2) < 2 else "flex-direction: row"
    h = "height: 100%" if level == 0 else st_h
    w = "width: 100%" if level == 0 else st_w
    margin = "" if level == 0 else "margin: 5px 0px 5px" if is_vert else "margin: 0px 5px 0px"

    is_new_vert = randint(1, 2) < 2
    if bg_style == "":
        is_reverse = not is_reverse

    children = [
        generate_flexbox_rec(level + 1, is_new_vert, is_reverse)
        for i in range(randint(2, 4))]

    outer = Div([
        h,
        w,
        flex_dir,
        bg_style,
        margin,
    ] + center_style, children)

    return outer


def generate_flexbox():
    flex = generate_flexbox_rec(0, False, False)
    outer = OuterDiv(flex)
    outer.render()
