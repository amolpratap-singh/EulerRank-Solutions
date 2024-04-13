import math

expo_list = list()
for _ in range(int(input())):
    b, e = map(int, input().split())
    expo_list.append((e*math.log(b), b,e))

k = int(input())
x, a, b = sorted(expo_list)[k-1]
print(a, b)