import math
t = int(input().strip())

for _ in range(t):
    N = int(input().strip())
    # Using Quadratic eqn
    a = 1
    b = 1
    c = -2
    
    D = int(math.sqrt(b**2 - 4 * a * c * N))
    
    rt_1 = (-b + D) // 2 * a
    rt_2 = (-b - D) // 2 * a
    
    if rt_1 > 0 and N == (rt_1 * ( rt_1 +1) // 2):
        print(rt_1)
    elif rt_2 > 0 and N == (rt_2 * ( rt_2 +1) // 2):
        print(rt_2)
    else:
        print(-1)