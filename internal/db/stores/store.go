package stores

import "context"

type IStore interface {
	Ping(ctx context.Context) error
}
