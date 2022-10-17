from .gen_div import generate_div
from .gen_flexbox import generate_flexbox
from .gen_grid import generate_grid
from .gen_margin import generate_margin
from .gen_transform import generate_transform
from .gen_lists import generate_lists
import base64

import threading
lock = threading.Lock()

def gen(type) -> str:
    with lock:
        if type == "301":
            generate_margin()
        elif type == "302":
            generate_grid()
        elif type == "303":
            generate_lists()
        elif type == "304":
            generate_transform()
        elif type == "305":
            generate_flexbox()
        elif type == "306":
            generate_div()
        else:
            return ""
        
        return "data:image/png;base64," + base64.b64encode(open("output.png", "rb").read()).decode("utf-8")
        
        
