import sys
import math

with open(sys.argv[1]) as f: s = f.read()
lines = s.split("\n")

waymarks = ['|', '-']

# +----- 1 x
# |
# |
# 1 y
direction = [0, 1]
position = [0, 0]

# Find start
position[0] = lines[0].index('|')

found_letters = []
steps = 0

while True:
    # walk straight until a cross
    while lines[position[1] + direction[1]][position[0] + direction[0]] not in ['+', ' ']:
        position[0] += direction[0]
        position[1] += direction[1]
        steps += 1
        if lines[position[1]][position[0]].isalpha():
            found_letters.append(lines[position[1]][position[0]])

    # save for later
    previous_mark = lines[position[1]][position[0]]

    # step into the cross
    steps += 1
    position[0] += direction[0]
    position[1] += direction[1]

    # decide where to go next
    found_route = False
    for degree in [90, -90]:
        possible = [
                int((direction[0] * math.cos(degree * math.pi/180)) - (direction[1] * math.sin(degree * math.pi/180))),
                int((direction[0] * math.sin(degree * math.pi/180)) - (direction[1] * math.cos(degree * math.pi/180)))
        ]
        if position[1] + possible[1] < 0 or position[0] + possible[0] < 0:
            next
        if (position[1] + possible[1] > len(lines)) or (position[0] + possible[0] > len(lines[position[1] + possible[1]])):
            next

        next_mark = lines[position[1] + possible[1]][position[0] + possible[0]]
        if next_mark in waymarks and next_mark != previous_mark:
            direction = possible
            found_route = True
            break

    if not found_route:
        break

print(''.join(found_letters))
print(steps)
