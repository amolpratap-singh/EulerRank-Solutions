def area(x1, y1, x2, y2, x3, y3):
    return abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)

n = int(input().strip())
count = 0
for _ in range(n):
    
    x1, y1, x2, y2, x3, y3 = map(int, input().split())
    triangle1_area = area(0, 0, x2, y2, x3, y3)
    triangle2_area = area(x1, y1, 0, 0, x3, y3)
    triangle3_area = area(x1, y1, x2, y2, 0, 0)
    
    triangle_area = area(x1, y1, x2, y2, x3, y3)
    
    if triangle_area == triangle1_area + triangle2_area + triangle3_area:
        count += 1

print(count)

