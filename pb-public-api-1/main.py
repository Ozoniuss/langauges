import requests
from collections import Counter

response = requests.get('https://api.publicapis.org/entries')
response_dict = dict(response.json())

count = response_dict['count']
entries = response_dict['entries']

# check that count is equal to the entries length
assert(count == len(entries))

all_keys = ([tuple(sorted(entries[i].keys())) for i in range(count)])
key_counter = Counter(all_keys)


# check that all keys across all entries are the same
assert(len(key_counter) == 1)

categories = [e['Category'] for e in entries]
print(Counter(categories))