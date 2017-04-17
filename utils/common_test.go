package utils

import "testing"
import "fmt"

func TestFloat2Str(t *testing.T) {
	type args struct {
		fv float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"floattostring",
			args{
				12.34,
			},
			"12.34",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float2Str(tt.args.fv)
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("Float2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
