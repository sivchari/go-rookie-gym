package main

import "testing"

func TestParallel(t *testing.T) {
	defer func() {
		t.Log("complete top level test")
	}()
	t.Parallel()
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
		{
			name: "invalid",
			args: args{
				i: 1,
				j: -1,
			},
			want:    0,
			wantErr: true,
		},
	}

	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("i is %d, tt.name is %s", i, tt.name)
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

func TestInvalidParallel(t *testing.T) {
	t.Parallel()
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
		{
			name: "invalid",
			args: args{
				i: 1,
				j: -1,
			},
			want:    0,
			wantErr: true,
		},
	}

	for i, tt := range tests {
		// i := i
		// tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("i is %d, tt.name is %s", i, tt.name)
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
