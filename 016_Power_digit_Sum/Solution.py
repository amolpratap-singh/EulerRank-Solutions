t = int(input().strip())

for var in range(t):
    n = int(input().strip())
    pow_val = pow(2, n)
    sum = 0
    for val in str(pow_val):
        sum += int(val)
    print(sum)