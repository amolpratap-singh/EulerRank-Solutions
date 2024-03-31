#count of n digit numbers
fac = [ i*9*10**(i-1) for i in range(1,18)]

#aggregate of the first 1 to n digit number
fac1 = [fac[0]]

for i in range(1,len(fac)):
    fac1.append(fac[i] + fac1[i-1])
   
def find_digit(n):
    if n < 9:
        return(n)
    for j in range(len(fac1)):
        if n == fac1[j]:
            return(9)
        if n < fac1[j]:
            break
    n, r = (n - fac1[j-1])//(j+1), (n - fac1[j-1])%(j+1)
    n = n-1 if r == 0 else n
    return(int(str(10**j+n)[j if r == 0 else r-1]))
    
for a0 in range(int(input())):
    n = map(int, input().split())
    prod = 1
    for i in n:
        prod *= find_digit(i)
    print(prod)