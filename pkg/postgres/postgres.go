package postgres

import (
	"context"
	"fmt"

	"user_base/internal/domain"
	"user_base/pkg/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER"     required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT"     required:"true"`
	Host     string `envconfig:"POSTGRES_HOST"     required:"true"`
	DBName   string `envconfig:"POSTGRES_DB_NAME"  required:"true"`
}

type Pool struct {
	*pgxpool.Pool
}

func New(ctx context.Context, c Config) (*Pool, error) {
	dsn := fmt.Sprintf("user=%s password=%s port=%s host=%s dbname=%s",
		c.User, c.Password, c.Port, c.Host, c.DBName)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	return &Pool{Pool: pool}, nil
}

func (p *Pool) Close() {
	p.Pool.Close()
	log.Info().Msg("postgres: closed")
}

func (p *Pool) CreateProfile(ctx context.Context, profile domain.Profile) error {
	executor := transaction.TryExtractTX(ctx)

	const query = `
		INSERT INTO profiles (id, name, age, email, phone)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
		    age = EXCLUDED.age,
		    email = EXCLUDED.email,
		    phone = EXCLUDED.phone
	`

	if _, err := executor.Exec(
		ctx,
		query,
		profile.ID,
		profile.Name,
		profile.Age,
		profile.Contacts.Email,
		profile.Contacts.Phone,
	); err != nil {
		return fmt.Errorf("executor.Exec (profiles): %w", err)
	}

	return nil
}

func (p *Pool) CreateProperty(ctx context.Context, property domain.Property) error {
	executor := transaction.TryExtractTX(ctx)

	const query = `
		INSERT INTO properties (profile_id, tags)
		VALUES ($1, $2)
	`

	if _, err := executor.Exec(ctx, query, property.ProfileID, property.Tags); err != nil {
		return fmt.Errorf("executor.Exec (properties): %w", err)
	}

	return nil
}

func (p *Pool) GetProfile(ctx context.Context, profileID string) (domain.Profile, error) {
	executor := transaction.TryExtractTX(ctx)

	const query = `
		SELECT id, name, age, email, phone
		FROM profiles
		WHERE id = $1
	`

	var (
		profile domain.Profile
		email   string
		phone   string
	)

	if err := executor.QueryRow(ctx, query, profileID).Scan(
		&profile.ID,
		&profile.Name,
		&profile.Age,
		&email,
		&phone,
	); err != nil {
		return domain.Profile{}, fmt.Errorf("scan profile: %w", err)
	}

	profile.Contacts = domain.Contacts{
		Email: email,
		Phone: phone,
	}

	return profile, nil
}
