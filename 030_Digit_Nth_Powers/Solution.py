nth_pow = int(input().strip())

output = set()

for i in range(100, 1000000):
    _sum = 0
    num = i
    while num != 0:
        _sum += (num %10)**nth_pow
        num //= 10
    if _sum == i:
        output.add(i)

print(sum(output))