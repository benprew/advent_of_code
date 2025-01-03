import re


def solve(robots, grid_dimensions, seconds):
    robots = run_robots(robots, grid_dimensions, seconds)
    # print_grid(robots, grid_dimensions)
    return count_quadrants(robots, grid_dimensions)


def run_robots(robots, grid_dimensions, seconds):
    max_x, max_y = grid_dimensions
    lowest_score = float("inf")
    best_iteration = 0
    best_robots = []
    for i in range(seconds):
        new_robots = []
        for (px, py), (vx, vy) in robots:
            new_px, new_py = (px + vx), (py + vy)
            new_px %= max_x
            new_py %= max_y
            if new_px < 0:
                new_px = max_x + new_px
            if new_py < 0:
                new_py = max_y + new_py
            new_robots.append(((new_px, new_py), (vx, vy)))
        robots = new_robots
        sf = count_quadrants(robots, grid_dimensions)
        # christmas tree likely has lowest score because most robots are in
        # one quadrant
        if sf < lowest_score:
            lowest_score = sf
            best_iteration = i + 1
            best_robots = robots
    print_grid(best_robots, grid_dimensions)
    print("Best iteration at: ", best_iteration, lowest_score)
    return robots


def print_grid(robots, grid_dimensions):
    max_x, max_y = grid_dimensions
    grid = [["." for _ in range(max_x)] for _ in range(max_y)]
    for (px, py), (_, _) in robots:
        if grid[py][px] == ".":
            grid[py][px] = 1
        else:
            grid[py][px] += 1
    for row in grid:
        print("".join([str(x) for x in row]))


# Split grid into quandrants and count the # of robots in each
# then multpliy those counts together
def count_quadrants(robots, grid_dimensions):
    max_x, max_y = grid_dimensions
    q1, q2, q3, q4 = 0, 0, 0, 0
    for (px, py), (_, _) in robots:
        if px < max_x // 2 and py < max_y // 2:
            q1 += 1
        elif px > max_x // 2 and py < max_y // 2:
            q2 += 1
        elif px < max_x // 2 and py > max_y // 2:
            q3 += 1
        elif px > max_x // 2 and py > max_y // 2:
            q4 += 1
    return q1 * q2 * q3 * q4


def parse(filename):
    robots = []
    with open(filename) as f:
        for y, line in enumerate(f):
            px, py, vx, vy = map(int, re.findall(r"-?\d+", line))
            robots.append(((px, py), (vx, vy)))
    return robots


if __name__ == "__main__":
    print("Part 1")
    print(solve(parse("input.txt"), (101, 103), 100))
    print("Part 2")
    solve(parse("input.txt"), (101, 103), 101 * 103)
