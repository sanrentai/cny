package cny

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "整百",
			args: args{202000.00},
			want: "贰拾万贰仟元整",
		},
		{
			name: "整元",
			args: args{2020004.00},
			want: "贰佰零贰万零肆元整",
		},
		{
			name: "整角",
			args: args{2020004.60},
			want: "贰佰零贰万零肆元陆角整",
		},
		{
			name: "正数",
			args: args{2020004.11},
			want: "贰佰零贰万零肆元壹角壹分",
		},
		{
			name: "负数",
			args: args{-2020004.11},
			want: "负贰佰零贰万零肆元壹角壹分",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.num); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
