package getenv

import (
	"testing"
)

func TestGetBool(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Negative: empty .env key bool",
			args: args{
				key:        "",
				defaultVal: true,
			},
			want: true,
		},
		{
			name: "Negative: failed parse .env key bool",
			args: args{
				key:        "WRONG_IS_SECURE",
				defaultVal: true,
			},
			want: true,
		},
		{
			name: "Positive: valid .env boolean",
			args: args{
				key:        "IS_SECURE",
				defaultVal: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBool(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Negative: empty integer key value",
			args: args{
				key:        "",
				defaultVal: 1,
			},
			want: 1,
		},
		{
			name: "Negative: wrong env integer  value",
			args: args{
				key:        "WRONG_INTEGER",
				defaultVal: 2,
			},
			want: 2,
		},
		{
			name: "Positivie: valid integer value",
			args: args{
				key:        "VALID_INTEGER",
				defaultVal: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Negative: empty key",
			args: args{
				key:        "",
				defaultVal: "oke",
			},
			want: "oke",
		},
		{
			name: "Positive: valid key and value",
			args: args{
				key:        "VALID_STRING",
				defaultVal: "oke-oke",
			},
			want: "oke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
