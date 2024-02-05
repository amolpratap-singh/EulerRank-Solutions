    
if __name__ == "__main__":
    t = int(input().strip())
    for a0 in range(t):
        n = int(input().strip())
        num = 2
        while num * num <= n:
            if n%num:
                num += 1
            else:
                n //= num
        print(n)
            
        
    