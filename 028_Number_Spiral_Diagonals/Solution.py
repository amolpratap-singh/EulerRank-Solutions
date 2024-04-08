t = int(input().strip())

for _ in range(t):
    n = (int(input().strip()) - 1) // 2
    result = (16 * n ** 3 + 30 * n ** 2 + 26 * n + 3) // 3 % (10 ** 9 + 7)
    print(result)