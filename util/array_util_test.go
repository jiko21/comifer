package util

import (
	"reflect"
	"testing"
)

func TestGetKeysOfMap(t *testing.T) {
	type args struct {
		maps map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Correctly return keys",
			args: args{
				maps: map[int]int{
					1: 2,
					2: 4,
					3: 6,
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeysOfMap(tt.args.maps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeysOfMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
