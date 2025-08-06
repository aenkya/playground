'''
You are given an array of words. For each word, find the shortest prefix that uniquely identifies it in the list.

The goal is to return an array of unique prefixes in the same order as the input.

Example:
Input:  ["zebra", "dog", "doggo", "duck", "dove"]
Output: ["z", "dog", "du", "dov"]

You can assume all strings are lowercase and non-empty

'''

# def say_hello(arr):
#     j = 0
#     result = []
#     while j < len(arr):
#         curr = ""
#         for i in range(len(arr[j])):
#             curr = curr+arr[j][i]

#             check = False
#             for val in arr:
#                 if val == arr[j]:
#                     continue
#                 if curr == val[0:i+1]: # O(L^2 * n^2) - time | O(n) - space
#                     check = True
#                     break
#             if not check or i == len(arr[j])-1:
#                 result.append(curr)
#                 break
#         j += 1
#     return result

# arr = ["zebra", "dog", "doggo", "duck", "dove"]
# print(say_hello(arr))

def unique_prefixes(words):
    class TrieNode:
        def __init__(self):
            self.children = {}
            self.count = 0

    root = TrieNode()

    # Build the trie
    for word in words:
        node = root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
            node.count += 1

    # Find unique prefixes
    result = []
    for word in words:
        node = root
        prefix = ""
        for char in word:
            prefix += char
            node = node.children[char]
            if node.count == 1:  # Unique prefix found
                result.append(prefix)
                break

    return result