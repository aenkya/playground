package algo

import "testing"

func Test_containsDuplicateII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains duplicate",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "does not contain duplicate",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicateII(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsDuplicateII() = %v, want %v", got, tt.want)
			}
		})
	}
}
