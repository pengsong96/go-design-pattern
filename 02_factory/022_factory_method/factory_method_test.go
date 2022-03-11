package factory

import (
	"reflect"
	"testing"
)

/**
工厂方法模式
在工厂方法上再抽象一层
*/
func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want FactoryMethods
	}{
		{
			name: "1",
			args: args{t: "1"},
			want: FactoryMethodsImpl1{},
		},
		{
			name: "2",
			args: args{t: "2"},
			want: FactoryMethodsImpl2{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
