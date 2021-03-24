package signature

import "testing"

func TestCommonRawStringGenerate_RawString(t *testing.T) {
	type args struct {
		params []string
	}
	tests := []struct {
		name string
		g    CommonRawStringGenerate
		args args
		want string
	}{
		{
			name: "t1",
			g:    CommonRawStringGenerate{},
			args: args{
				params: []string{
					"78912fhaksjfh",
					"jkj",
				},
			},
			want: "78912fhaksjfhjkj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RawString(tt.args.params); got != tt.want {
				t.Errorf("CommonRawStringGenerate.RawString() = %v, want %v", got, tt.want)
			}
		})
	}
}
