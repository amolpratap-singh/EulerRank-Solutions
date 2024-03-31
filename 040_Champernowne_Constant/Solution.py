# 50 % Test got passed cause Runtime error

def champernowne_contsant(n):
    
    champernowne = ""
    i = 1
    while len(champernowne) < n:
        champernowne += str(i)
        i +=1
        
    return champernowne

t = int(input().strip())


for i in range(t):
    input_string = input().strip()
    num = list(map(int, input_string.split()))
    mul = 1
    for var in num:
        res = champernowne_contsant(var)
        mul *= int(res[var-1])
    print(mul)