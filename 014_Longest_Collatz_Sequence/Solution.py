longest_numbers = [0, 1]
longest = 0
terms = [0, 1]
t = int(input().strip())
for var in range(t):
    n = int(input().strip())
    for i in range(len(longest_numbers), n+1):
        num = i
        count = 0
        while num != 1 and num >= i:
            if num % 2 == 0:
                num //=2
            else:
                num = 3 * num + 1
            count += 1
        length = count + terms[num]
        terms.append(length)
        if length >= longest:
            longest = length
            longest_numbers.append(i)
        else:
            longest_numbers.append(longest_numbers[-1])
    print(longest_numbers[n])