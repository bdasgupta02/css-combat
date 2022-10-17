from html2image import Html2Image
from .canvas import canvas_front, canvas_back
from .div import Div
import os

class OuterDiv:
    def __init__(self, children: Div) -> None:
        self.children = children

    def generate(self):
        generated = '{}<div>{}</div>{}'.format(
            canvas_front,
            self.children.generate(),
            canvas_back
        )
        print("Generated HTML")
        return generated

    # modification: save as random hash string of 10 chars and return the hash string
    # need to be a bit careful to screenshot -> potentially it might look different
    # -> test it out heavily for all problem types using the string generated

    # need to screenshot and convert to base64 -> check if this can be direct
    def render(self):
        hi = Html2Image(size=(400,300))
        hi.screenshot(html_str=self.generate(), save_as='output.png')
