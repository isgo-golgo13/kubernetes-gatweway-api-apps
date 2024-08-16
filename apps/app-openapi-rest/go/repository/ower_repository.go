package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type OwnerRepository struct {
	redisClient *redis.Client
}

func NewOwnerRepository(redisClient *redis.Client) *OwnerRepository {
	return &OwnerRepository{redisClient: redisClient}
}

// AddDataSequenceToOwner adds a data sequence ID to an owner's list of sequences.
func (r *OwnerRepository) AddDataSequenceToOwner(ctx context.Context, ownerID string, dataSequenceID string) error {
	// Ensure the owner exists
	ownerExists, err := r.redisClient.Exists(ctx, ownerID).Result()
	if err != nil {
		return err
	}
	if ownerExists == 0 {
		return fmt.Errorf("owner with ID %s does not exist", ownerID)
	}

	// Add the DataSequence ID to the owner's list of sequences
	_, err = r.redisClient.RPush(ctx, ownerID+":sequences", dataSequenceID).Result()
	return err
}

// GetOwnerByID retrieves an owner by their ID from Redis.
func (r *OwnerRepository) GetOwnerByID(ctx context.Context, ownerID string) (map[string]interface{}, error) {
	owner, err := r.redisClient.HGetAll(ctx, ownerID).Result()
	if err != nil {
		return nil, err
	}

	if len(owner) == 0 {
		return nil, fmt.Errorf("owner with ID %s not found", ownerID)
	}

	// Convert map[string]string to map[string]interface{}
	ownerInterface := make(map[string]interface{}, len(owner))
	for k, v := range owner {
		ownerInterface[k] = v
	}

	return ownerInterface, nil
}

// GetAllOwners retrieves all owners from Redis.
func (r *OwnerRepository) GetAllOwners(ctx context.Context) ([]map[string]interface{}, error) {
	var owners []map[string]interface{}

	// Assuming owners are stored with a specific key pattern
	keys, err := r.redisClient.Keys(ctx, "owner:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		owner, err := r.GetOwnerByID(ctx, key)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}

	return owners, nil
}
