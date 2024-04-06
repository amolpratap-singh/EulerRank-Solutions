

def get_pow_val_sum(a, b):
    out = a ** b
    _sum = 0
    for var in str(out):
        _sum += int(var)
    return _sum

num = int(input().strip())



max_sum = 0


for a in range(1,num):
    for b in range(1, num):
        if max_sum < get_pow_val_sum(a,b):
            max_sum = get_pow_val_sum(a,b)
        
print(max_sum)