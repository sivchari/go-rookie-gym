//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/$GOPACKAGE/$GOFILE
package user

import (
	"context"

	"github.com/sivchari/go-rookie-gym/infrastructure"
)

type Repository interface {
	User(ctx context.Context, user *User) (int64, error)
	Transaction(txm infrastructure.TxManager) Repository
}
