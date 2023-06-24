import os
import sys
__cwd = os.path.abspath(os.path.dirname(__file__))
__modules = os.path.abspath(__cwd)

sys.path.append(os.path.abspath(__modules))

from algorithms import is_valid_palindrome

pal = is_valid_palindrome.palindrome_pairs

def main():
    pairs = pal(["abcd", "dcba", "lls", "s", "sssll"])
    print(pairs)

if __name__ == "__main__":
    main()