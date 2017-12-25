import sys

a = 1
d = e = f = g = h = 0

b = 67 * 100 + 100000
c = b + 17000

while True:
    f = True
    d = 2
    while True:
        e = 2
        print(g)
        while True:
            g = (d * e) - b
            if g == 0:
                f = False

            e += 1
            g = e - b
            if g == 0:
                break

        d += 1
        g = d - b
        if g == 0:
            break

    if not f:
        h += 1
        print(h)

    g = b - c
    if g == 0:
        break

    b += 17
