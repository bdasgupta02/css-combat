from random import randint
from .div.div import Div
from .div.outer_div import OuterDiv
from .gen_color import generate_pair

# Notes:
# - template should have 1x 50x50 div

hori_start = 50
hori_end = 350

vert_start = 20
vert_end = 240

hori_split = (vert_end + vert_start) / 2
vert_split = (hori_end + hori_start) / 2

split_windows = [
    [
        [vert_start, vert_split],
        [hori_start, hori_split]
    ],
    [
        [vert_start, vert_split],
        [hori_split, hori_end]
    ],
    [
        [vert_split, vert_end],
        [hori_start, hori_split]
    ],
    [
        [vert_split, vert_end],
        [hori_split, hori_end]
    ]
]

# transform only

def generate_transform():
    size = randint(40, 80)
    d = 4
    n = 2
    divions = [randint(n - 1, n + 1) for i in range(d)]
    translated = []
    pair = generate_pair()

    for i in range(d):
        v1 = split_windows[i][0][0]
        v2 = split_windows[i][0][1]
        h1 = split_windows[i][1][0]
        h2 = split_windows[i][1][1]

        def create_translate():
          return "translate({}px, {}px)".format(
              randint(
                  h1,
                  h2
              ),
              randint(
                  v1,
                  v2
              )
          )

        def create_rotate():
          is_rotate = randint(1, 4) < 2
          if not is_rotate:
            return ""
          else:
            return "rotate({}deg)".format(randint(20, 320))

        num = divions[i]
        translated += [Div([
            "transform-origin: top left",
            f"height: {size}px",
            f"width: {size}px",
            "transform: {} {}".format(
                create_translate(),
                create_rotate()
            ),
            "background-color: {}{}".format(pair[1], "A4"),
            "position: absolute"
        ]) for i in range(num)]

    translated_outer = OuterDiv(
        Div([
            "height: 100%",
            "width: 100%",
            "position: relative",
            "background-color: {}".format(pair[0])
        ], translated)
    )

    translated_outer.render()

