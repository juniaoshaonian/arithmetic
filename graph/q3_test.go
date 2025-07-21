package graph

import (
	"fmt"
	"testing"
)

func TestGetGraphSide(t *testing.T) {
	testcases := []struct {
		name    string
		input   [][]int
		wantRes [][]int
	}{
		{
			name: "正常案例",
			input: [][]int{
				{
					5, 5,
				},
				{
					1, 3,
				},
				{
					3, 5,
				},
				{
					1, 2,
				},
				{
					2, 4,
				},
				{
					4, 5,
				},
			},
			wantRes: [][]int{
				{
					1, 3, 5,
				},
				{
					1, 2, 4, 5,
				},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetGraphSide(tc.input)
			fmt.Println(res)
		})

	}
}
