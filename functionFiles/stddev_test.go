package functionFiles

import (
	"testing"
)

func TestVariance(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Single element",
			args: args{nums: []float64{10.0}},
			want: 0,
		},
		{
			name: "Multiple elements",
			args: args{nums: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want: 2.0,
		},
		{
			name: "All same numbers",
			args: args{nums: []float64{2.0, 2.0, 2.0, 2.0}},
			want: 0,
		},
		{
			name: "Mixed values",
			args: args{nums: []float64{-1.0, 0.0, 1.0}},
			want: 0.6666666666666666,
		},
		{
			name: "Large range of values",
			args: args{nums: []float64{-100.0, 0.0, 50.0, 100.0}},
			want: 5468.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.nums); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStddev(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Zero variance",
			args: args{variance: 0},
			want: 0,
		},
		{
			name: "Positive variance",
			args: args{variance: 4.0},
			want: 2.0,
		},
		{
			name: "Large variance",
			args: args{variance: 24850.0},
			want: 157.63882770434446,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stddev(tt.args.variance); got != tt.want {
				t.Errorf("Stddev() = %v, want %v", got, tt.want)
			}
		})
	}
}
