digits = [0, 1]
f1 = f2 = 1
c = 10
for i in range(3, 25000):
    (f1, f2) = (f2, f1 + f2)
    if f2 >= c:
        c *= 10
        digits.append(i)

t = int(input().strip())

for a0 in range(t):
    n = int(input().strip())
    print(digits[n])
