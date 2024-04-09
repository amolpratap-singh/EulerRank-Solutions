import math

n,k=map(int,input().strip().split())
count=0
for i in range(1,n+1):
    for r in range(n):
        if math.comb(i,r)>k:
            count+=1
print(count)