def convert_to_base(decimal_number, base):
    if decimal_number == 0:
        return 0

    digits = "0123456789"

    result = ""
    while decimal_number > 0:
        remainder = decimal_number % base
        result = digits[remainder] + result
        decimal_number //= base

    return result

def check_palindrome(num):
    return str(num) == str(num)[::-1]

n, k = map(int, input().strip().split())

_sum = 0

for num in range(1, n):
    base_repr = convert_to_base(num, k)
    if check_palindrome(num) and check_palindrome(base_repr):
        _sum += num
print(_sum)
