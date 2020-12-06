import re
from typing import Tuple

def parse(s: str) -> Tuple[str, dict]:
    m = re.match(r"^(?P<min>\d+)-(?P<max>\d+) (?P<letter>\S): (?P<pw>\S+)$", s)
    pw = m.group("pw")

    policy = {}
    policy['min'] = int(m.group("min"))
    policy['max'] = int(m.group("max"))
    policy['letter'] = m.group("letter")
    return pw, policy

def check_policy(pw: str, policy: dict) -> bool:
    letter_count = sum([l == policy['letter'] for l in pw])
    return policy['min'] <= letter_count <= policy['max']


with open('input') as fh:
    data = [l.strip() for l in fh.readlines()]


print('Part 1')
print(sum([check_policy(*parse(l)) for l in data]))

################

def parse(s: str) -> Tuple[str, dict]:
    m = re.match(r"^(?P<pos1>\d+)-(?P<pos2>\d+) (?P<letter>\S): (?P<pw>\S+)$", s)
    pw = m.group("pw")

    policy = {}
    policy['pos1'] = int(m.group("pos1"))
    policy['pos2'] = int(m.group("pos2"))
    policy['letter'] = m.group("letter")
    return pw, policy

def check_policy(pw: str, policy: dict) -> bool:
    return bool(pw[policy['pos1'] - 1] == policy['letter']) != bool(pw[policy['pos2'] - 1] == policy['letter'])  # xor
assert check_policy('abc', {'letter': 'a', 'pos1': 1, 'pos2': 3}) == True
assert check_policy('cba', {'letter': 'a', 'pos1': 1, 'pos2': 3}) == True
assert check_policy('aba', {'letter': 'a', 'pos1': 1, 'pos2': 3}) == False
assert check_policy('bbb', {'letter': 'a', 'pos1': 1, 'pos2': 3}) == False

print('Part 2')
print(sum([check_policy(*parse(l)) for l in data]))