package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"fgh",
			args{
				i: "test",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.i); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
