from math import factorial

def get_factorial_sum(num):
    if len(num) == 1:
        return factorial(num)
    _sum = 0
    
    for val in str(num):
        _sum += int(factorial(val))
    return _sum

t = int(input().strip())

for _ in range(t):
    num, chain_len = map(int, input().split())
    
    for i in range(1,num):
        val = get_factorial_sum(i)
        
        
    