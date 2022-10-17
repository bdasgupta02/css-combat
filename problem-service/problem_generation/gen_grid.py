from random import randint
from .div.div import Div
from .div.outer_div import OuterDiv
from .gen_color import generate_pair

# try generating other types of images too

def get_forward_neighbors(pos, rows, cols):
    delta_rows = [0, 1]
    delta_cols = [1, 0]
    neighbors = []

    for i in range(2):
        row = pos[0] + delta_rows[i]
        col = pos[1] + delta_cols[i] 
        if row > -1 and row < rows and col > -1 and col < cols:
            neighbors.append([row, col])

    return neighbors

def generate_template_areas():
    area = 0

    rows = randint(3, 4)
    cols = randint(4, 6)
    areas = [[-1 for x in range(cols)] for y in range(rows)]

    for r in range(rows):
        for c in range(cols):
            if areas[r][c] == -1:
                area += 1
                areas[r][c] = area
            
            slot_area = areas[r][c]
            neighbors = get_forward_neighbors([r, c], rows, cols)

            is_top = r != 0 and areas[r-1][c] == slot_area
            is_left = c != 0 and areas[r][c-1] == slot_area

            alr_merged = False
            for n in neighbors:
                should_merge = randint(1, 4) < 2
                if not should_merge:
                    continue

                if not alr_merged and ((not is_top and r == n[0]) or (not is_left and c == n[1])):
                    areas[n[0]][n[1]] = slot_area
                    alr_merged = True

    return areas, area

def randomize_dots(areas):
    max_dots = (len(areas) * len(areas[0])) - 4
    for r in range(len(areas)):
        for c in range(len(areas[0])):
            slot_area = areas[r][c]
            if max_dots > 0 and randint(1, 6) < 3 and areas[r-1][c] != slot_area and areas[r][c-1] != slot_area:
                areas[r][c] = '.'
                max_dots -= 1

    return areas

def generate_grid():
    areas, area = generate_template_areas()
    areas = randomize_dots(areas)
    pair = generate_pair()

    templ = ""
    for r in range(len(areas)):
        row = "'"
        for c in range(len(areas[0])):
            row += "a{} ".format(areas[r][c]) if areas[r][c] != '.' else '. '
        row += "'"
        templ += row

    objects = []
    for i in range(1, area + 1):
        if "a{} ".format(i) in templ:
            objects.append(Div([
                "grid-area: a{}".format(i),
                "background-color: {}".format(pair[1]),
            ]))

    outer = OuterDiv(Div([
        "height: 100%",
        "display: grid",
        "grid-template-areas: {}".format(templ),
        "grid-gap: {}px".format(randint(5, 12)),
        "background-color: {}".format(pair[0]),
    ], objects))

    outer.render()