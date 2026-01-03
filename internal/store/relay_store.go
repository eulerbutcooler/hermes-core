package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eulerbutcooler/hermes-core/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RelayStore struct {
	db *pgxpool.Pool
}

func NewRelayStore(db *pgxpool.Pool) *RelayStore {
	return &RelayStore{db: db}
}

func (s *RelayStore) CreateRelay(ctx context.Context, req models.CreateRelayRequest) (string, error) {
	relayID := uuid.New()
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)
	queryRelay := `INSERT INTO relays (id, user_id, name) VALUES($1,$2,$3)`
	_, err = tx.Exec(ctx, queryRelay, relayID, req.UserID, req.Name)
	if err != nil {
		return "", fmt.Errorf("Failed to insert relay: %w", err)
	}
	configJSON, err := json.Marshal(req.Config)
	if err != nil {
		return "", err
	}

	queryAction := `INSERT INTO relay_actions (relay_id, action_type, config) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, queryAction, relayID, req.ActionType, configJSON)
	if err != nil {
		return "", fmt.Errorf("failed to insert action: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return relayID.String(), nil
}
