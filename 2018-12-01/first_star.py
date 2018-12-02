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


def calculate(directions):
    num = 0
    for operation, change in directions:
        print(num, operation, change)
        num = get_number(num, operation, int(change))
    return num


def main():
    directions = get_file('first_star_input.csv')
    final_number = calculate(directions)
    print(final_number)


if __name__ == '__main__':
    main()
