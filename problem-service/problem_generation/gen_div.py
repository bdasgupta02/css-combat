from random import randint
from .div.div import Div
from .div.outer_div import OuterDiv
from .gen_color import generate_light_color, generate_dark_color

h = "height: 80%"
w = "width: 25%"

center_style = [
    "display: flex",
    "flex-direction: row",
    "align-items: center",
    "justify-content: center",
]

max_level = 2


def generate_div_unit(level):
    if level == max_level:
        return Div([], "")

    bg_col = "background-color: {}".format(generate_dark_color(
    )) if level % 2 == 0 else "background-color: {}".format(generate_light_color())
    bg_grad = "background-image: linear-gradient({}deg, {}, {})".format(
        randint(0, 359), generate_dark_color(), generate_light_color())
    bg = bg_col if randint(1, 3) > 1 else bg_grad

    border_style = "border: {}px solid {}".format(
        randint(1, 5), generate_light_color())
    border = "" if randint(1, 3) < 2 else border_style

    shadow_style = "box-shadow: {} {}px {}px {}px".format(
        generate_dark_color(), randint(2, 8), randint(2, 8), randint(6, 10))
    print(shadow_style)
    shadow = "" if randint(1, 3) < 2 else shadow_style

    children = []
    num_child = randint(0, 3)
    has_children = randint(1, 2) < 2
    if has_children and num_child > 0:
        gap = Div(["width: 5px", "height: 1px"], "")
        for i in range(num_child):
            children = children + [gap, generate_div_unit(level + 1), gap]

    return Div([
        h,
        w,
        bg,
        border,
        shadow
    ] + center_style, children)


def generate_div():
    gap = Div(["width: 20px", "height: 1px"], "")
    outer = OuterDiv(
        Div([
            "height: 100%",
            "width: 100%",
            "background-color: {}".format(generate_light_color()),
        ] + center_style, [
            generate_div_unit(0),
            gap,
            generate_div_unit(0),
            gap,
            generate_div_unit(0),
        ])
    )
    outer.render()
