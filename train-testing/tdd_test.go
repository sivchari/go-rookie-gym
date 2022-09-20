package main

import "testing"

func TestAddStruct(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				i: 1,
				j: 2,
			},
			want:    3,
			wantErr: false,
		},
		// invalidケースを書いてみよう！
	}

	for i, tt := range tests {
		t.Logf("i is %d, tt.name is %s", i, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: err = %s, wantErr = %v", err.Error(), tt.wantErr)
				return
			}
			if tt.want != got {
				t.Errorf("unexpected error: want = %d, got = %d", tt.want, got)
				return
			}
		})
	}
}

// struct/mapdで実行順がかわる
func TestAddMap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := map[string]struct {
		args    args
		want    int
		wantErr bool
	}{
		"success": {
			args: args{
				i: 1,
				j: 2,
			},
			want:    3,
			wantErr: false,
		},
		"invalid": {
			args: args{
				i: 1,
				j: -1,
			},
			want:    0,
			wantErr: true,
		},
	}

	for k, tt := range tests {
		t.Logf("tt.name is %s", k)
		t.Run(k, func(t *testing.T) {
			got, err := Add(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: err = %s, wantErr = %v", err.Error(), tt.wantErr)
				return
			}
			if tt.want != got {
				t.Errorf("unexpected error: want = %d, got = %d", tt.want, got)
				return
			}
		})
	}
}
