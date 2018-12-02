def get_file(filename):
    with open(filename) as f:
        return [(line[0], line[1:].strip()) for line in f]


FUNCTIONS = {
    '+': lambda num, change: num + change,
    '-': lambda num, change: num - change
}


def get_number(num, operation, change):
    f = FUNCTIONS[operation]
    if not f:
        raise Exception('Unknown operation {}'.format(operation))
    return f(num, change)


numbers = set()


def num_is_dupe(num):
    if num in numbers:
        return True
    numbers.add(num)
    return False


def get_directions(directions, i):
    return directions[i % len(directions)]


def calculate(directions):
    num = 0
    i = 0
    while not num_is_dupe(num):
        operation, change = get_directions(directions, i)
        i += 1
        num = get_number(num, operation, int(change))
    return num


def main():
    directions = get_file('second_star_input.csv')
    final_number = calculate(directions)
    print(final_number)


if __name__ == '__main__':
    main()
