package algo

import (
	"reflect"
	"sort"
	"testing"
)

func sortGroup(group [][]string) {
	for _, g := range group {
		sort.Strings(g)
	}

	sort.Slice(group, func(i, j int) bool {
		if len(group[i]) == 0 || len(group[j]) == 0 {
			return len(group[i]) < len(group[j])
		}

		return group[i][0] < group[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "Example 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:  "Empty input",
			input: []string{},
			want:  [][]string{},
		},
		{
			name:  "Single word",
			input: []string{"abc"},
			want:  [][]string{{"abc"}},
		},
		{
			name:  "All anagrams",
			input: []string{"abc", "bca", "cab"},
			want:  [][]string{{"abc", "bca", "cab"}},
		},
		{
			name:  "No anagrams",
			input: []string{"abc", "def", "ghi"},
			want:  [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:  "Mixed case",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.input)
			sortGroup(got)
			sortGroup(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
