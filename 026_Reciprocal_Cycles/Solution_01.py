            
def get_long_division_cycle_length(denominator: int) -> int:
    numerator = 10
    digits = []
    remainders = []
    cur = numerator
    while True:
        digit = cur // denominator
        cur = cur % denominator
        if cur in remainders:
            break
        digits.append(digit)
        remainders.append(cur)
        cur *= numerator
        if cur == 0:
            return 0
        while cur < denominator:
            cur *= 10
            digits.append(0)
            
    while remainders[0] != cur:
        remainders.pop(0)
    return len(remainders)


test_case  = int(input().strip())

for var in range(test_case):
    number = int(input().strip())
    max_length = 0
    num = 1
    for denominator in range(1, number):
        rem_length = get_long_division_cycle_length(denominator)
        if max_length < rem_length:
            max_length = rem_length
            num = denominator
    print(num)
        
        
