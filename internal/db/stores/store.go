package stores

import (
	"context"

	"github.com/google/uuid"
	"github.com/vxdiazdel/rest-api/models"
)

type IStore interface {
	// users
	CreateUser(ctx context.Context, email, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
