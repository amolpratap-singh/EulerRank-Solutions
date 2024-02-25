grid = []
for grid_i in range(20):
    grid_t = [int(grid_temp) for grid_temp in input().strip().split(' ')]
    grid.append(grid_t)
    
max_product = 0

for row in range(20):
    for col in range(20):
        # Horizontal product
        if col + 3 < 20:
            product = grid[row][col] * grid[row][col + 1] * grid[row][col + 2] * grid[row][col + 3]
            max_product = max(max_product, product)
        
        # Vertical product
        if row + 3 < 20:
            product = grid[row][col] * grid[row + 1][col] * grid[row + 2][col] * grid[row + 3][col]
            max_product = max(max_product, product)
        
        # Diagonal products
        if row + 3 < 20 and col + 3 < 20:
            # Top-left to bottom-right diagonal
            product = grid[row][col] * grid[row + 1][col + 1] * grid[row + 2][col + 2] * grid[row + 3][col + 3]
            max_product = max(max_product, product)
            
            # Top-right to bottom-left diagonal
            product = grid[row][col + 3] * grid[row + 1][col + 2] * grid[row + 2][col + 1] * grid[row + 3][col]
            max_product = max(max_product, product)
            
print(max_product)
