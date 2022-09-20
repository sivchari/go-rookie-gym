package user

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	groupd "github.com/sivchari/go-rookie-gym/domain/group"
	gmock "github.com/sivchari/go-rookie-gym/domain/mock/group"
	umock "github.com/sivchari/go-rookie-gym/domain/mock/user"
	userd "github.com/sivchari/go-rookie-gym/domain/user"
	"github.com/sivchari/go-rookie-gym/infrastructure"
)

func TestPutUser(t *testing.T) {
	t.Parallel()

	type args struct {
		userName string
	}

	tests := []struct {
		name    string
		args    args
		gmock   func(m *gmock.MockRepository)
		umock   func(m *umock.MockRepository)
		want    *UserOutput
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				userName: "sample name",
			},
			gmock: func(m *gmock.MockRepository) {
				m.EXPECT().Transaction(gomock.Any()).Return(m)
				m.EXPECT().Group(gomock.Any(), &groupd.Group{
					UserID: 1,
					Name:   "default",
				}).Return(int64(1), nil)
			},
			umock: func(m *umock.MockRepository) {
				m.EXPECT().Transaction(gomock.Any()).Return(m)
				m.EXPECT().PutUser(gomock.Any(), &userd.User{
					Name: "sample name",
				}).Return(int64(1), nil)
			},
			want: &UserOutput{
				GroupID: 1,
				UserID:  1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			gm := gmock.NewMockRepository(ctrl)
			tt.gmock(gm)
			um := umock.NewMockRepository(ctrl)
			tt.umock(um)
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Errorf("failed to create a sqlmock err = %s", err.Error())
				return
			}
			mock.ExpectBegin()
			mock.ExpectCommit()
			txm := infrastructure.NewTxManager(db)
			uc := NewUsecase(um, gm, txm)
			got, err := uc.User(context.Background(), &UserInput{
				Name: "sample name",
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("failed to uc err = %s", err.Error())
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(UserOutput{}, "JWT")); diff != "" {
				t.Errorf("PutUser differs: (-got +want)\n%s", diff)
				return
			}
		})
	}
}
