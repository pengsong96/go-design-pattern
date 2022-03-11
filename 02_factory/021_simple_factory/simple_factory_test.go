package factory

import (
	"reflect"
	"testing"
)

/**
简单工厂模式
通过工厂类决定具体实例化那个
*/
func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want Factory
	}{
		{
			name: "1",
			args: args{t: "1"},
			want: implement1{},
		},
		{
			name: "2",
			args: args{t: "2"},
			want: implement2{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewIRuleConfigParser(tt.args.t)
			if !reflect.DeepEqual(parser, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", parser, tt.want)
			}
		})
	}
}
