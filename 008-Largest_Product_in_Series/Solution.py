#!/bin/python3

def get_k_product(num):
    product = 1
    for var in num:
        product *= int(var)
    return product
        

t = int(input().strip())
for a0 in range(t):
    n,k = input().strip().split(' ')
    n,k = [int(n),int(k)]
    num = input().strip()
    
    sol = [get_k_product(num[i:i+k]) for i in range(len(num)-k+1)]
    print(max(sol))