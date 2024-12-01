def part1():
  file = open("input.txt", "r")
  left = []
  right = []
  total = 0
  for line in file:
    left.append(line[0:5])
    right.append(line[8:13])
  left = sorted(left)
  right = sorted(right)

  for i in range(len(left)):
    total += abs(int(left[i]) - int(right[i]))

  print(total)
  file.close()

def part2():
  file = open("input.txt", "r")
  left = []
  right = []
  total = 0
  for line in file:
    left.append(line[0:5])
    right.append(line[8:13])

  for i in range(len(left)):
    leftScore = right.count(left[i]) * int(left[i])
    total += leftScore

  print(total)
  file.close()

part2()
