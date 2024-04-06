test_case =int(input().strip())

names_list = list()

for i in range(test_case):
    names_list.append(input().strip())

names_list.sort()
q = int(input().strip())

for i in range(q):
    name_q=input().strip()
    _sum = 0
    for count,name in enumerate(names_list):
        if name_q == name:
            for k in name:
                _sum = _sum + (ord(k)-ord('A')+1)
            _sum = (count+1) * _sum
    print(_sum)