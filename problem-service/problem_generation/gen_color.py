from random import choice

def generate_dark_color():
  return "#"+''.join([choice('0123456') for j in range(6)])

def generate_light_color():
  return "#"+''.join([choice('BCDEF') for j in range(6)])

def generate_pair():
  return [generate_dark_color(), generate_light_color()]
