# Time Complexity for below code: O(n^2)
n = int(input().strip())

num_sqr = list()

for a in range(2, n+1):
    for b in range(2, n+1):
        num_sqr.append(a ** b)
  
x = set(num_sqr)

print(len(x))