def factorial_num(n):
    if n == 0 or n == 1:
        return 1
    return n * factorial_num(n-1)

t = int(input().strip())

for i in range(t):
    sum = 0
    n = int(input().strip())
    val = factorial_num(n)
    for var in str(val):
        sum += int(var)
    print(sum)