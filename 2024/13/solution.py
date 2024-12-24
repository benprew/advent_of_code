#!/usr/bin/env python3

import sys
import re


def solve(px, py, bx, by, ax, ay):
    count_a = (px * by - py * bx) / (ax * by - ay * bx)
    count_b = (px - ax * count_a) / bx
    if count_a % 1 == count_b % 1 == 0:
        return count_a, count_b, int(count_a * 3 + count_b)
    return -1, -1, 0


filename = sys.argv[1] if len(sys.argv) > 1 else "input"


problems = []
with open(filename) as f:
    for line in f:
        if "Button A" in line:
            # parse a_x and a_y from line
            match = re.search(r"X([+-]?\d+), Y([+-]?\d+)", line)
            if match:
                a_x, a_y = int(match.group(1)), int(match.group(2))
        elif "Button B" in line:
            # parse a_x and a_y from line
            match = re.search(r"X([+-]?\d+), Y([+-]?\d+)", line)
            if match:
                b_x, b_y = int(match.group(1)), int(match.group(2))
        elif "Prize" in line:
            # parse x and y from line
            # Prize: X=4300, Y=2204
            match = re.search(r"X=(\d+), Y=(\d+)", line)
            if match:
                x, y = int(match.group(1)), int(match.group(2))
            problems.append((x, y, b_x, b_y, a_x, a_y))

sum = 0
for i, problem in enumerate(problems):
    x, y, b_x, b_y, a_x, a_y = problem

    ans = solve(x, y, b_x, b_y, a_x, a_y)
    if ans[0] == -1:
        continue

    print(x, y, b_x, b_y, a_x, a_y, "=>", ans)
    sum += ans[2]
print("Part 1 Sum:", sum)

sum = 0
for i, problem in enumerate(problems):
    x, y, b_x, b_y, a_x, a_y = problem
    x += 10000000000000
    y += 10000000000000
    ans = solve(x, y, b_x, b_y, a_x, a_y)
    if ans[0] == -1:
        continue

    # print(x, y, b_x, b_y, a_x, a_y, "=>", ans)
    sum += ans[2]
print("Part2 Sum:", sum)
