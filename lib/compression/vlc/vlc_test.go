package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr       string
		chunksSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "Base test",
			args: args{
				bStr:       "0010000001101000011",
				chunksSize: 8,
			},
			want: BinaryChunks{"00100000", "01101000", "01100000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunksSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		encodedData []byte
	}
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "base test",
			data: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want: "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			if got := decoder.Decode(tt.data); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
