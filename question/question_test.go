package question

import (
	"reflect"
	"testing"
)

func Test_getDescriptionsOfOptions(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "return correctly",
			args: args{
				options: []Option{
					{
						Value:       "a",
						Description: "aa",
					},
					{
						Value:       "b",
						Description: "bb",
					},
				},
			},
			want: []string{"aa", "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescriptionsOfOptions(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDescriptionsOfOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValuesOfOptions(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "return correctly",
			args: args{
				options: []Option{
					{
						Value:       "a",
						Description: "aa",
					},
					{
						Value:       "b",
						Description: "bb",
					},
				},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValuesOfOptions(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValuesOfOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueFromSelect(t *testing.T) {
	type args struct {
		options       []Option
		selectedValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return correctly when matched",
			args: args{
				options: []Option{
					{
						Value:       "a",
						Description: "aa",
					},
					{
						Value:       "b",
						Description: "bb",
					},
				},
				selectedValue: "aa",
			},
			want: "a",
		},
		{
			name: "return correctly when not matched",
			args: args{
				options: []Option{
					{
						Value:       "a",
						Description: "aa",
					},
					{
						Value:       "b",
						Description: "bb",
					},
				},
				selectedValue: "afa",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueFromSelect(tt.args.options, tt.args.selectedValue); got != tt.want {
				t.Errorf("GetValueFromSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
