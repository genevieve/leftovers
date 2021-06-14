package common

import "testing"


func TestMatchRegex(t *testing.T) {
	type args struct {
		resourceName string
		filter       string
		regex		 bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "doesn't match",
			args: args{
				resourceName: "matc",
				filter:       "match",
				regex: true,
			},
			want: false,
		},
		{
			name: "does match",
			args: args{
				resourceName: "please-match-me",
				filter: "match",
				regex: true,
			},
			want: true,
		},
		{
			name: "matches for empty filter",
			args: args{
				resourceName: "something-arbitrary",
				filter: "",
				regex: true,
			},
			want: true,
		},
		{
			name: "matches for empty filter",
			args: args{
				resourceName: "something-arbitrary",
				filter: "",
				regex: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchRegex(tt.args.resourceName, tt.args.filter, tt.args.regex); got != tt.want {
				t.Errorf("MatchRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}