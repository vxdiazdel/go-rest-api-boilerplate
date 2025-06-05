package stores

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/vxdiazdel/rest-api/models"
)

func (s *PostgresStore) CreateUser(
	ctx context.Context,
	email, password string,
) (*models.User, error) {
	q := `
		INSERT INTO users (
			email,
			password
		) VALUES (
			@email,
			@password 
		)
		RETURNING
			id,
			email,
			created_at,
			updated_at
	`
	args := pgx.NamedArgs{
		"email":    email,
		"password": password,
	}
	row := s.DB().QueryRow(ctx, q, args)

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, ErrUserEmailTaken
		}
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return &user, nil
}

func (s *PostgresStore) GetUserByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
	q := `
		SELECT
			id,
			email,
			created_at,
			updated_at
		FROM users
		WHERE id = (@id)
		AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"id": id,
	}
	row := s.DB().QueryRow(ctx, q, args)

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return &user, nil
}

func (s *PostgresStore) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	q := `
		SELECT
			id,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email = (@email)
		AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"email": email,
	}
	row := s.DB().QueryRow(ctx, q, args)

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return &user, nil
}
