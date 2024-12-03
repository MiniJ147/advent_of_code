import collections


cnt = collections.defaultdict(int)
it = []

with open("input.txt") as f:
    for line in f:
        line = line.strip()
        line = line.split(" ")

        it.append(int(line[0]))
        cnt[int(line[3])] += 1

print(sum(i * cnt[i] for i in it))
