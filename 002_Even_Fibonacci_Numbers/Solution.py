
if __name__ == "__main__":
    t = int(input().strip())
    for a0 in range(t):
        n = int(input().strip())
        a =1 
        b =2 
        c =3
        sum = 2
        while c<n:
            if c%2 == 0:
                sum +=c
            a, b, c = b, c, b+c
        print(sum)
            
            