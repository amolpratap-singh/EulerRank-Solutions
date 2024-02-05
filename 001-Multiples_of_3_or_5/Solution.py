def sum_of_muliples(n, k):
    d = n // k
    return k * (d * (d + 1)) // 2

if __name__ == "__main__":
    t = int(input().strip())
    for a0 in range(t):
        n = int(input().strip())
        
        print(sum_of_muliples(n-1, 3) + 
              sum_of_muliples(n-1, 5) - 
              sum_of_muliples(n-1, 15))
        
    