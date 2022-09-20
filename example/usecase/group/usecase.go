package group

import (
	"context"
	"strconv"

	domain "github.com/sivchari/go-rookie-gym/domain/group"
)

type Usecase interface {
	Groups(ctx context.Context, id string) ([]*domain.Group, error)
	Group(ctx context.Context, gi *GroupInput) (int64, error)
}

type usecase struct {
	repository domain.Repository
}

func NewUsecase(gr domain.Repository) Usecase {
	return &usecase{
		repository: gr,
	}
}

func (u *usecase) Groups(ctx context.Context, id string) ([]*domain.Group, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	groups, err := u.repository.Groups(ctx, intID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

type GroupInput struct {
	UserID int
	Name   string
}

func (u *usecase) Group(ctx context.Context, gi *GroupInput) (int64, error) {
	g := domain.NewGroup(gi.UserID, gi.Name)
	id, err := u.repository.Group(ctx, g)
	if err != nil {
		return 0, err
	}
	return id, nil
}
