package algo

// To execute Go code, please declare a func main() in a package "main"
// Implement Atoi

// Given a string s, the objective is to convert it into integer format without utilizing any built-in functions. Refer the below steps to know about atoi() function.

// Cases for atoi() conversion:

// Skip any leading whitespaces.
// Check for a sign (‘+’ or ‘-‘), default to positive if no sign is present.
// Read the integer by ignoring leading zeros until a non-digit character is encountered or end of the string is reached. If no digits are present, return 0.
// If the integer is greater than 2^31 – 1, then return 2^31 – 1 and if the integer is smaller than -231, then return -231.
// Examples:

// Input: s = "-123"
// Output: -123
// Explanation: It is possible to convert -123 into an integer so we returned in the form of an integer
// Input: s = "  -"
// Output: 0
// Explanation: No digits are present, therefore the returned answer is 0.
// Input: s = " 1231231231311133"
// Output: 2147483647
// Explanation: The converted number will be greater than 231 – 1, therefore print 231 – 1 = 2147483647.
// Input: s = "-999999999999"
// Output: -2147483648
// Explanation: The converted number is smaller than -231, therefore print -231 = -2147483648.
// Input: s = "  -0012gfg4"
// Output: -12
// Explanation: After ignoring leading zeros nothing is read after -12 as a non-digit character ‘g’ was encountered.

/*
brute force:
- converting string to byte array
- check first character,
- is_negative is false
if negative, set is_negative to true. set starting point to 1
- init digit_sum at 0
-  looping through characters
- for each character, subtract zero to it to convert to integer
- digit_sum += 1 * 10 ^ i
- make negative if negative
- return digit value

2 *10^32 > given_max return given_max

// Examples
-123, 0012gfg4, 4444g4444 => 4444, '453  2345' => 453, '   4 453'
is_negative: true
start = 1
deduct = len(input)
current_total = 0

non_numeric_encountered
loop from last to start:
- if 231 - current_total < val, return 231
-  if is first value and negative, is_negative to true. break
- is_not_numeric
  - non_numeric_encountered

- else if non_numeric_encountered true?:
  - reset power, start, non_numeric_encountered

2, 12, 12, 12,
- increment power, start

- if non_numeric_encountered is true, return 0
*/
func atoi(s string) int {
	const (
		maxInt = 1<<31 - 1
		minInt = -1 << 31
	)

	bytes := []byte(s)
	if len(bytes) == 0 {
		return 0
	}

	isNegative := false
	start := 0

	for start < len(bytes) && bytes[start] == ' ' {
		start++
	}

	if start < len(bytes) && (bytes[start] == '-' || bytes[start] == '+') {
		if bytes[start] == '-' {
			isNegative = true
		}

		start++
	}

	digitSum := 0

	for i := start; i < len(bytes); i++ {
		if bytes[i] < '0' || bytes[i] > '9' {
			break
		}

		digit := int(bytes[i] - '0')

		if digitSum > (maxInt-digit)/10 {
			if isNegative {
				return minInt
			}

			return maxInt
		}

		digitSum = digitSum*10 + digit
	}

	if isNegative {
		digitSum = -digitSum
	}

	return digitSum
}
