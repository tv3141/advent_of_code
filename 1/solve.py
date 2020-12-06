with open('input', 'r') as fh:
    numbers = fh.readlines()

numbers = [int(n.strip()) for n in numbers]
print(numbers)

print('Part 1')
for i1, n1 in enumerate(numbers):
    for n2 in numbers[i1:]:
        if n1 + n2 == 2020:
            print(n1, n2, n1*n2)

print('Part 2')
for i1, n1 in enumerate(numbers):
    for i2, n2 in enumerate(numbers[i1:]):
        for n3 in numbers[i1+i2:]:
            if n1 + n2 + n3 == 2020:
                print(n1, n2, n3, n1*n2*n3)
