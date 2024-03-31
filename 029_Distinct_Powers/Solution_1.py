# Efficient Solution
def distinct_power_terms(n: int) -> int:
    count = 0
    test = [0]*(n+1)
    for i in range(2, n+1):
        if test[i]:
            continue
        combined = set()
        pw = 2
        while (i**pw <= n):
            test[i**pw] = True
            s = set([b*pw for b in range(2, n+1) if b*pw > n])
            combined.update(s)
            pw += 1
        count += len(combined) + n-1
    return count

n = int(input().strip())
print(distinct_power_terms(n))