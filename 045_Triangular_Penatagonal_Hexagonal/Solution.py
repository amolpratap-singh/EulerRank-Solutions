def is_triangular(x):
    return True if (1+8*x)**0.5==int((1+8*x)**0.5) else False
def is_pentagonal(x):
    return True if (1+24*x)**0.5==int((1+24*x)**0.5) and int((1+24*x)**0.5+1)%6==0 else False

n,k,b = map(int, input().strip().split())

pen_max=int((1+24*n)**0.5+1)//6
hex_max=int((1+8*n)**0.5+1)//4
if k==3:
    for i in range(1,pen_max):
        p=i*(3*i-1)//2
        if is_triangular(p):
            print(p)
else:
    for i in range(1,hex_max):
        h=i*(2*i-1)
        if is_pentagonal(h):
            print(h)
        

