LIMIT = 2_000_000

def get_prime_num(n):
    primes = [True] * (n + 1)
    primes[0], primes[1] = False, False

    p = 2
    while p * p <= n:
        if primes[p] == True:
            for i in range(p * p, n + 1, p):
                primes[i] = False
        p += 1

    prime_numbers = [i for i in range(n + 1) if primes[i]]
    return prime_numbers

prime_num = get_prime_num(LIMIT)

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(sum(num for num in prime_num if num <= n))
