package point

import "testing"

type test struct {
	name        string
	p1          *Point
	p2          *Point
	expectedRes float64
}

func TestDistance(t *testing.T) {
	tests := []test{
		{
			name:        "points equals",
			p1:          New(1, 1),
			p2:          New(1, 1),
			expectedRes: 0,
		},
		{
			name:        "point not equal",
			p1:          New(0, 0),
			p2:          New(3, 4),
			expectedRes: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Distance(tt.p1, tt.p2)

			if res != tt.expectedRes {
				t.Errorf("res = %v, expectedRes = %v", res, tt.expectedRes)
			}
		})
	}
}
