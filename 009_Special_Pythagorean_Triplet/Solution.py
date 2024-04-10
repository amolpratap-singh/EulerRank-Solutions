
def get_max_product(num):
    product = -1
    b = 0
    c = 0
    for a in range(1, num//3+1):
        b = (num * num -  2 * num * a) // (2 * num - 2 * a)
        c = num - b - a
        
        if c * c == (a * a + b * b):
            temp = a * b * c
            if temp > product:
                product = temp
    return product

t = int(input().strip())

for _ in range(t):
    n = int(input().strip())
    print(get_max_product(n))