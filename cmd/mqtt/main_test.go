package main_test

import (
	"testing"

	main "github.com/xkamail/huberlink-platform/cmd/mqtt"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/thing"
)

func TestExtractTopic(t *testing.T) {
	type args struct {
		_topic string
	}
	tests := []struct {
		name    string
		args    args
		want    snowid.ID
		want1   string
		wantErr bool
	}{
		{
			name: "valid topic",
			args: args{
				_topic: thing.PrefixTopic + "/123/thing/irremote/learning",
			},
			want:    snowid.ID(123),
			want1:   "irremote/learning",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := main.ExtractTopic(tt.args._topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractTopic() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractTopic() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
