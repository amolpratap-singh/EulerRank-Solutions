# Time Complexity Exceeded 
def get_divisor_sum(num):
    _sum = 0 
    for i in range(1, num):
        if num % i == 0:
            _sum += i
    return _sum


t = int(input().strip())

for _ in range(t):
    n = int(input().strip())
    _sum_of_divisors = 0
    for i in range(1, n+1):
        for j in range(i, n+1):
            if get_divisor_sum(i) == j and get_divisor_sum(j) == i \
                and i != j:
                    _sum_of_divisors += (i + j)
                    
    print(_sum_of_divisors)