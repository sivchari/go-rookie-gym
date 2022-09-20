//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/$GOPACKAGE/$GOFILE
package group

import (
	"context"

	"github.com/sivchari/go-rookie-gym/infrastructure"
)

type Repository interface {
	Group(ctx context.Context, group *Group) (int64, error)
	Groups(ctx context.Context, id int) ([]*Group, error)
	Transaction(txm infrastructure.TxManager) Repository
}
