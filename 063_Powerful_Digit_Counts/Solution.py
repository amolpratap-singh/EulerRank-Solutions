
n = int(input().strip())

i = 1

while True:
    num = pow(i,n)
    if len(str(num)) == n:
        print(num)
    if len(str(num)) > n:
        break
    i += 1