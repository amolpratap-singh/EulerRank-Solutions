import sys


def gcd(a,b):
    if b == 0:
        return a
    return gcd(b, a%b)

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    val = 1
    
    for i in range(1, n+1):
        val = int(val * i // gcd(val, i))
        
    print(val)