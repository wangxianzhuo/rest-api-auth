package signature

import (
	"reflect"
	"testing"
)

func TestHmacSha1Sign_Signature(t *testing.T) {
	type args struct {
		rawParams []string
		key       string
	}
	tests := []struct {
		name string
		s    HmacSha1Sign
		args args
		want string
	}{
		{
			name: "t1",
			s: HmacSha1Sign{
				rawStringGenerater: CommonRawStringGenerate{},
			},
			args: args{
				rawParams: []string{
					"lakjflad;",
					"asda;f",
					"j;asfjw",
					"ier234",
				},
				key: "123",
			},
			want: "jJj8L8WInhfxzvtwz8QZSrDYSg0=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Signature(tt.args.rawParams, tt.args.key); got != tt.want {
				t.Errorf("HmacSha1Sign.Signature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSha1Sign_GetRawStringGenerater(t *testing.T) {
	tests := []struct {
		name    string
		s       HmacSha1Sign
		want    RawStringGenerater
		wantErr bool
	}{
		{
			name:    "t1",
			s:       HmacSha1Sign{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "t2",
			s: HmacSha1Sign{
				rawStringGenerater: CommonRawStringGenerate{},
			},
			want:    CommonRawStringGenerate{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRawStringGenerater()
			if (err != nil) != tt.wantErr {
				t.Errorf("HmacSha1Sign.GetRawStringGenerater() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HmacSha1Sign.GetRawStringGenerater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSha1Sign_Check(t *testing.T) {
	type args struct {
		rawParams     []string
		key           string
		encodedString string
	}
	tests := []struct {
		name string
		s    HmacSha1Sign
		args args
		want bool
	}{
		{
			name: "t1",
			s: HmacSha1Sign{
				rawStringGenerater: CommonRawStringGenerate{},
			},
			args: args{
				rawParams: []string{
					"lakjflad;",
					"asda;f",
					"j;asfjw",
					"ier234",
				},
				key:           "123",
				encodedString: "jJj8L8WInhfxzvtwz8QZSrDYSg0=",
			},
			want: true,
		},
		{
			name: "t2",
			s: HmacSha1Sign{
				rawStringGenerater: CommonRawStringGenerate{},
			},
			args: args{
				rawParams: []string{
					"lakjflad;",
					"asda;f",
					"j;asfjw",
					"sdf",
				},
				key:           "123",
				encodedString: "jJj8L8WInhfxzvtwz8QZSrDYSg0=",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Check(tt.args.rawParams, tt.args.key, tt.args.encodedString); got != tt.want {
				t.Errorf("HmacSha1Sign.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
