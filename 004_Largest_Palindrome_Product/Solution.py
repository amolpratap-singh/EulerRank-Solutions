import sys

def check_palindrome(num):
    return str(num) == str(num)[::-1]


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    largest_plaindrome = 0

    while largest_plaindrome==0:
        n -= 1
        if check_palindrome(n):
            for j in range(100, 999):
                r = n/j
                if r % 1 == 0 and r < 1000:
                    largest_plaindrome = n
                    break
  
    print(largest_plaindrome)