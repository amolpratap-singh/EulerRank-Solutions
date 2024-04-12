import sys

# Only 3 cases get pass out of 5 due to Runtime error for two cases

N = 10000
prime_num = [True for _ in range(N+1)]

p = 2

while p*p <= N:
    if (prime_num[p] == True):
        for i in range(p * p, N+1, p):
            prime_num[i] = False
    p += 1

prime = list()

for i in range(1, len(prime_num)):
    if (prime_num[i] == True):
        prime.append(i)

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(prime[n])
    