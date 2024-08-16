package repository

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type DataSequence struct {
	ID             string
	Data           []byte
	DataOffsetStart int
	DataOffsetEnd   int
	ResequenceCount int
	OwnerID        string
	Timestamp      string
	PurgeExpiry    string
}

type DataSequenceRepository struct {
	redisClient *redis.Client
	ownerRepo   *OwnerRepository
}

func NewDataSequenceRepository(redisClient *redis.Client, ownerRepo *OwnerRepository) *DataSequenceRepository {
	return &DataSequenceRepository{redisClient: redisClient, ownerRepo: ownerRepo}
}

func (r *DataSequenceRepository) CreateDataSequence(ctx context.Context, ds DataSequence) error {
	encodedData := base64.StdEncoding.EncodeToString(ds.Data)

	err := r.redisClient.HSet(ctx, "data_sequence:"+ds.ID, map[string]interface{}{
		"ID":              ds.ID,
		"Data":            encodedData,
		"DataOffsetStart": ds.DataOffsetStart,
		"DataOffsetEnd":   ds.DataOffsetEnd,
		"ResequenceCount": ds.ResequenceCount,
		"OwnerID":         ds.OwnerID,
		"Timestamp":       ds.Timestamp,
		"PurgeExpiry":     ds.PurgeExpiry,
	}).Err()
	if err != nil {
		return err
	}

	return r.ownerRepo.AddDataSequenceToOwner(ctx, ds.OwnerID, ds.ID)
}

func (r *DataSequenceRepository) GetDataSequenceByID(ctx context.Context, id string) (*DataSequence, error) {
	data, err := r.redisClient.HGetAll(ctx, "data_sequence:"+id).Result()
	if err != nil || len(data) == 0 {
		return nil, err
	}

	decodedData, _ := base64.StdEncoding.DecodeString(data["Data"])

	return &DataSequence{
		ID:              data["ID"],
		Data:            decodedData,
		DataOffsetStart: r.intFromString(data["DataOffsetStart"]),
		DataOffsetEnd:   r.intFromString(data["DataOffsetEnd"]),
		ResequenceCount: r.intFromString(data["ResequenceCount"]),
		OwnerID:         data["OwnerID"],
		Timestamp:       data["Timestamp"],
		PurgeExpiry:     data["PurgeExpiry"],
	}, nil
}

func (r *DataSequenceRepository) GetAllDataSequences(ctx context.Context) ([]DataSequence, error) {
	keys, err := r.redisClient.Keys(ctx, "data_sequence:*").Result()
	if err != nil {
		return nil, err
	}

	var sequences []DataSequence
	for _, key := range keys {
		ds, err := r.GetDataSequenceByID(ctx, key[14:]) // remove "data_sequence:" prefix
		if err == nil {
			sequences = append(sequences, *ds)
		}
	}

	return sequences, nil
}

func (r *DataSequenceRepository) intFromString(s string) int {
	// Utility function to convert a string to an int, assuming valid input
	i, _ := strconv.Atoi(s)
	return i
}
