from typing import List

with open('input') as fh:
    data = [l.strip() for l in fh.readlines()]

grid = [[col for col in row] for row in data]
grid_height = len(grid)
grid_width = len(grid[0])
for i, row in enumerate(grid):
    print(''.join(row))
    if i > 10: 
        print('[...]')
        break

class Coord():
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return f'Coord({self.x}, {self.y})'

    def __add__(self, other):
        return Coord(self.x + other.x, self.y + other.y)

coord1 = Coord(1, 1)
coord2 = Coord(2, 3)
result = coord1 + coord2
assert result.x == 3
assert result.y == 4


def traverse(grid: List[List], start: Coord, step: Coord) -> List:
    pos = start
    locations = []
    grid_height = len(grid)
    grid_width = len(grid[0])
    while pos.y <= grid_height - 1:
        locations.append(grid[pos.y][pos.x % grid_width])
        #if grid[pos.y][pos.x % grid_width] == '#':
        #    grid[pos.y][pos.x % grid_width] = 'X'
        #else:
        #    grid[pos.y][pos.x % grid_width] = '0'
        pos += step
    return locations

locations = traverse(grid, start=Coord(0, 0), step=Coord(3, 1))
assert len(locations) == grid_height

# print solution grid
#for i, row in enumerate(grid):
#    print(''.join(row))
#    if i > 10: break

print('Part 1')
print(sum([loc =='#' for loc in locations]))

print('Part 2')
solutions = [
    sum([loc =='#' for loc in traverse(grid, start=Coord(0, 0), step=Coord(1, 1))]),
    sum([loc =='#' for loc in traverse(grid, start=Coord(0, 0), step=Coord(3, 1))]),
    sum([loc =='#' for loc in traverse(grid, start=Coord(0, 0), step=Coord(5, 1))]),
    sum([loc =='#' for loc in traverse(grid, start=Coord(0, 0), step=Coord(7, 1))]),
    sum([loc =='#' for loc in traverse(grid, start=Coord(0, 0), step=Coord(1, 2))]),
    ]
print(solutions)
result = 1
for s in solutions:
    result *= s
print(f'Result: {result}')