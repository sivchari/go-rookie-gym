package group

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	domain "github.com/sivchari/go-rookie-gym/domain/group"
	mock "github.com/sivchari/go-rookie-gym/domain/mock/group"
)

func TestGroups(t *testing.T) {
	t.Parallel()

	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		mockfn  func(m *mock.MockRepository)
		want    []*domain.Group
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				id: "1",
			},
			mockfn: func(m *mock.MockRepository) {
				m.EXPECT().Groups(gomock.Any(), 1).Return([]*domain.Group{
					{
						ID:     1,
						UserID: 1,
						Name:   "default",
					},
				}, nil)
			},
			want: []*domain.Group{
				{
					ID:     1,
					UserID: 1,
					Name:   "default",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				id: "aaa",
			},
			mockfn:  func(m *mock.MockRepository) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid",
			args: args{
				id: "1",
			},
			mockfn: func(m *mock.MockRepository) {
				m.EXPECT().Groups(gomock.Any(), 1).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			m := mock.NewMockRepository(ctrl)
			tt.mockfn(m)
			uc := NewUsecase(m)
			got, err := uc.Groups(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("failed to uc err = %s", err.Error())
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Groups differs: (-got +want)\n%s", diff)
				return
			}
		})
	}
}

func TestPutGroup(t *testing.T) {
	t.Parallel()

	type args struct {
		userID int
		name   string
	}

	tests := []struct {
		name    string
		args    args
		mockfn  func(m *mock.MockRepository)
		want    int64
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				userID: 1,
				name:   "name",
			},
			mockfn: func(m *mock.MockRepository) {
				m.EXPECT().Group(gomock.Any(), &domain.Group{
					UserID: 1,
					Name:   "name",
				}).Return(int64(1), nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				userID: 1,
				name:   "name",
			},
			mockfn: func(m *mock.MockRepository) {
				m.EXPECT().Group(gomock.Any(), &domain.Group{
					UserID: 1,
					Name:   "name",
				}).Return(int64(0), errors.New("error"))
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			m := mock.NewMockRepository(ctrl)
			tt.mockfn(m)
			uc := NewUsecase(m)
			got, err := uc.Group(context.Background(), &GroupInput{
				UserID: tt.args.userID,
				Name:   tt.args.name,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("failed to uc err = %s", err.Error())
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PutGroup differs: (-got +want)\n%s", diff)
				return
			}
		})
	}
}
