import sys

with open(sys.argv[1]) as f: s = f.read().strip()
lines = [x.split(" ") for x in s.split("\n")]

registers = {}
head = 0
mul_count = 0

def resolve(v):
    try:
        return int(v)
    except ValueError:
        if v not in registers:
            registers[v] = 0
        return registers[v]

while head < len(lines):
    command = lines[head][0]
    (a, b) = [lines[head][1], lines[head][2]]
    if command == "mul":
        mul_count += 1
        registers[a] = resolve(a) * resolve(b)
        head += 1
    elif command == "set":
        registers[a] = resolve(b)
    elif command == "sub":
        registers[a] = resolve(a) - resolve(b)
    elif command == "jnz":
        if resolve(a) != 0:
            head += (resolve(b) - 1)

    head += 1

print(mul_count)
