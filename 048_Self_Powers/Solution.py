n = int(input().strip())

total_sum = 0
modulus = 10**10

for i in range(1, n+1):
    total_sum += pow(i, i, modulus)

last_ten_digits = total_sum % modulus
print(last_ten_digits)
