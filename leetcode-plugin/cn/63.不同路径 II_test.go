package cn

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{{1}}},
			0,
		},
		{
			"2",
			struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{{0}}},
			1,
		},
		{
			"3",
			struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			2,
		},
		{
			"4",
			struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{{0}, {0}}},
			1,
		},
		{
			"5",
			struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
