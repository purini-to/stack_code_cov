package msg

import (
	"reflect"
	"testing"
)

func TestProvideHello(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *Hello
	}{
		{
			name: "引数の文字列でインスタンス生成できること",
			args: args{text: "test"},
			want: &Hello{Text: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProvideHello(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProvideHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHello_GetMessage(t *testing.T) {
	type fields struct {
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Textを引数にした場合は'Hello Text.'を取得できること",
			fields: fields{Text: "Text"},
			want:   "Hello Text.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hello{
				Text: tt.fields.Text,
			}
			if got := h.GetMessage(); got != tt.want {
				t.Errorf("Hello.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
