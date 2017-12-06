import sys

with open(sys.argv[1]) as f: s = f.read()

stacks = [int(x) for x in s.strip().split("\t")]

visited = [stacks[:]]
redistribution_cycles = 0
last_index = len(stacks) - 1

while True:
    redistribution_cycles += 1

    m = max(stacks)
    index = stacks.index(m)
    stacks[index] = 0

    while m > 0:
        index += 1
        if index > last_index:
            index = 0
        stacks[index] += 1; m -= 1

    if stacks[:] in visited:
        break

    visited.append(stacks[:])

print("Redistribution cycles: %d" % (redistribution_cycles))
