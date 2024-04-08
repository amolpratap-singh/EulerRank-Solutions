def fibonacci_index(digits):
    a, b = 1, 1
    index = 2
    while len(str(b)) < digits:
        a, b = b, a + b
        index += 1
    return index 

t = int(input().strip())

for a0 in range(t):
    n = int(input().strip())
    print(fibonacci_index(n))
