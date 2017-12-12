import sys
import re

with open(sys.argv[1]) as f: s = f.read()
lines = s.strip().split("\n")

programs = {}
for line in lines:
    id, links = re.search('^(\d+) <-> (.*)$', line).groups()
    links = links.split(", ")
    programs[id] = links

def build_group_for(target):
    programs_in_this_group = [target]
    queue = [target]
    while len(queue) > 0:
        item = queue.pop(0)
        links = programs[item]
        for link in links:
            if link not in programs_in_this_group:
                if link not in queue:
                    queue.append(link)
                programs_in_this_group.append(link)

    return programs_in_this_group

groups = []
for key in programs.keys():
    if len(set([key in g for g in groups])) > 1:
        continue
    groups.append(build_group_for(key))

print(len(groups))
