result = [2]
for x in range(3, 110000, 2):
    m=True
    for i in range(3, int(x**0.5)+1, 2):
        if x % i == 0:
            m=False
            break
    if m:
        result.append(x)

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(result[n-1])