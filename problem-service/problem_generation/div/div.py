class Div:
    def __init__(self, style: list, children = "") -> None:
        self.style = ';'.join(style)
        self.children = children

    def generate(self):
        return '<div style="{}">{}</div>'.format(
            self.style,
            self.generate_children()
        )
    
    def generate_children(self):
        return self.children if (isinstance(self.children, str)) else self.generate_list_children() if (isinstance(self.children, list)) else self.children.generate()
    
    def generate_list_children(self):
        s = ''
        for c in self.children:
            s += c.generate()
        return s

