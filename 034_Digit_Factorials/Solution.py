from math import factorial

n = int(input().strip())
sum_num = 0
for i in range(10, n):
    sum = 0
    num = i
    while num > 0:
        sum += factorial(num % 10)
        num //= 10

    if sum % i == 0:
       sum_num += i
print(sum_num)