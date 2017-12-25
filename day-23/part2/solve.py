def is_prime(a):
    return all(a % i for i in xrange(2, a))

b = (67 * 100) + 100000
c = b + 17000
nonprime = 0
while b <= c:
    nonprime += 1 - int(is_prime(b))
    b += 17

print(nonprime)
