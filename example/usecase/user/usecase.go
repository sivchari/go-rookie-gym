package user

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	groupd "github.com/sivchari/go-rookie-gym/domain/group"
	userd "github.com/sivchari/go-rookie-gym/domain/user"
	"github.com/sivchari/go-rookie-gym/infrastructure"
)

type Usecase interface {
	User(ctx context.Context, ui *UserInput) (*UserOutput, error)
}

type usecase struct {
	ur  userd.Repository
	gr  groupd.Repository
	txm infrastructure.TxManager
}

func NewUsecase(ur userd.Repository, gr groupd.Repository, txm infrastructure.TxManager) Usecase {
	return &usecase{
		ur:  ur,
		gr:  gr,
		txm: txm,
	}
}

type UserInput struct {
	Name string
}

type UserOutput struct {
	GroupID int64
	UserID  int64
	JWT     string
}

func (u *usecase) User(ctx context.Context, ui *UserInput) (*UserOutput, error) {
	user := userd.NewUser(ui.Name)
	var (
		uid int64
		gid int64
	)
	var err error
	fn := func(txm infrastructure.TxManager) error {
		uid, err = u.ur.Transaction(txm).User(ctx, user)
		if err != nil {
			return err
		}
		group := groupd.NewGroup(int(uid), "default")
		gid, err = u.gr.Transaction(txm).Group(ctx, group)
		if err != nil {
			return err
		}
		return nil
	}
	if err := u.txm.BeginTransaction(fn); err != nil {
		return nil, err
	}
	secretKey := os.Getenv("jwt-secret")
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"name": ui.Name,
	}
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	var uo UserOutput
	uo.UserID = uid
	uo.GroupID = gid
	uo.JWT = tokenStr
	return &uo, nil
}
