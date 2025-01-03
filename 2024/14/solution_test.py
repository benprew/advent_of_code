import sys
import os

# Add the directory containing your script to the Python path
sys.path.append(os.path.dirname(os.path.abspath(__file__)))

import solution

robots = solution.parse("test")
print("start", robots)
print("finish", solution.run_robots(robots, (11, 7), 2))
print(solution.solve(solution.parse("test"), (11, 7), 100))
