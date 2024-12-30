import re


def parse(filename):
    robots = {}
    with open(filename) as f:
        for y, line in enumerate(f):
            px, py, vx, vy = map(int, re.findall(r"-?\d+", line))
            robots[(px, py)] = (vx, vy)
    return robots
