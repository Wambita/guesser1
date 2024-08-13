package functionFiles

import "testing"

func TestRange(t *testing.T) {
	type args struct {
		numslice []float64
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 int
	}{
		{
			name:  "Single element",
			args:  args{numslice: []float64{10.0}},
			want1: 10,
			want2: 10,
		},
		{
			name:  "Multiple elements with normal distribution",
			args:  args{numslice: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want1: 1,
			want2: 6,
		},
		{
			name:  "Mixed positive and negative values",
			args:  args{numslice: []float64{-1.0, -2.0, 3.0, 4.0}},
			want1: -4,
			want2: 6,
		},
		{
			name:  "All negative numbers",
			args:  args{numslice: []float64{-1.0, -2.0, -3.0, -4.0, -5.0}},
			want1: -6,
			want2: -1,
		},
		{
			name:  "All same numbers",
			args:  args{numslice: []float64{2.0, 2.0, 2.0, 2.0, 2.0}},
			want1: 2,
			want2: 2,
		},
		{
			name:  "Floating point numbers",
			args:  args{numslice: []float64{1.5, 2.5, 3.5, 4.5}},
			want1: 1,
			want2: 5,
		},
		{
			name:  "Large range of values",
			args:  args{numslice: []float64{-100.0, 0.0, 50.0, 100.0}},
			want1: -121,
			want2: 146,
		},
		{
			name:  "Small range of values",
			args:  args{numslice: []float64{1.0, 2.0, 2.5, 3.0}},
			want1: 1,
			want2: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Range(tt.args.numslice)
			if got != tt.want1 {
				t.Errorf("Range() got = %v, want %v", got, tt.want1)
			}
			if got1 != tt.want2 {
				t.Errorf("Range() got1 = %v, want %v", got1, tt.want2)
			}
		})
	}
}
