import math

def pentagon_num(num):
    return num * (3 * num - 1) // 2

def is_pentagon_num(num_p):
    # Using Quadratic eqn
    a = 3
    b = -1
    c = -2
    D = int(math.sqrt(b**2 - 4 * a * c * num_p))
    x = 2 * a
    rt_1 = (-b + D) // x
    rt_2 = (-b - D) // x
    
    if rt_1 > 0 and num_p == pentagon_num(rt_1):
        return True
    elif rt_2 > 0 and num_p == pentagon_num(rt_2):
        return True
    else:
        return False

n, k = map(int, input().strip().split())

for num in range(k +1 , n):
    if is_pentagon_num(pentagon_num(num) - pentagon_num(num - k)) or is_pentagon_num(pentagon_num(num) + pentagon_num(num - k)):
        print(pentagon_num(num))
    
    