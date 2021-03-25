package access

import (
	"reflect"
	"testing"
)

func TestFilePermissionLoader_PermissioLoad(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		l       FilePermissionLoader
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success test",
			l: FilePermissionLoader{
				fileName: "./fileloader_test.json",
			},
			args: args{
				id: "a",
			},
			want: []string{
				"/1",
				"/2",
				"/3",
				"/4",
			},
			wantErr: false,
		},
		{
			name: "not exist id test",
			l: FilePermissionLoader{
				fileName: "./fileloader_test.json",
			},
			args: args{
				id: "b",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "file load error test",
			l: FilePermissionLoader{
				fileName: "",
			},
			wantErr: true,
		},
		{
			name: "json parse error test",
			l: FilePermissionLoader{
				fileName: "./fileloader_error_test.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PermissioLoad(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilePermissionLoader.PermissioLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilePermissionLoader.PermissioLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
