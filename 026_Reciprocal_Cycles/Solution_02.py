def recurring_digits(n): 
    while n%2==0:
        n=n//2
    while n%5==0:
        n=n//5
    if n==1:
        return 0
    digits=1
    rem=1
    while True:
        rem*=10 
        rem=rem%n
        if rem==1:
            break
        digits+=1
    return digits

answers=[0,0,0,0]
max_digits=0
for i in range(3,10**4):
    digits=recurring_digits(i)
    if digits>max_digits:
        max_digits=digits
        number=i
    answers.append(number)
    
for _ in range(int(input())):
    n=int(input())
    print(answers[n])