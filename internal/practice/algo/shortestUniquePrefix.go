package algo

import "strings"

/*
Prompt:
You are given an array of words. For each word, find the shortest prefix that uniquely identifies it in the list.

The goal is to return an array of unique prefixes in the same order as the input.

✏️ Example:
Input:  ["zebra", "dog", "duck", "dove"]
Output: ["z", "dog", "du", "dov"]

	•	You can assume all strings are lowercase and non-empty.
	•	Try to solve this efficiently (not brute force).
	•	Optimize for time complexity if you can.
*/

/*
Brute force:
- loop through words
-  for each word, loop through increasing subsequences from 0 to i
- look for the subsequence in the rest of the array
- if present, increment i and try again.
- if absent, add word substring to uniquePrefix array
- if we reach the end of the word, add the entire word

attempting using trie:
- with a trie, create a node of character and count
- create empty root
- loop through array
- for each word, loop through characters
- insert character in trie with the next character being a child of the current character
- if character already exists, increase count by 1
- if not, add to children
				{}
			    /\
			z:1.  d:3
		   /    /   \
		  e:1  u:1    o:2
		/       /       /  \
	   b:1     c:1     v:1  g:1
	  /        |        |
	r:1        k:1     e:1
	/
   a:1

- add children of root to traverse list
- traverse list
- for each element in the traverse list,
- loop through every available prefix in prefix map
- add char to prefix: e.g. z:'z'
- if char count is 1, remove that string from map, and add to res
- continue
- add children for non 1 count chars to traverse list
-
*/

func ShortestUniquePrefixes(words []string) []string {
	up := make([]string, len(words))

	for i := 0; i < len(words); i++ {
		found := false

		for j := 1; j <= len(words[i]); j++ {
			subseq := words[i][:j]
			unique := true

			for k := 0; k < len(words); k++ {
				if i != k && strings.HasPrefix(words[k][:j], subseq) {
					unique = false
					break
				}
			}

			if unique {
				up[i] = subseq
				found = true

				break
			}
		}

		if !found {
			up[i] = words[i]
		}
	}

	return up
}
