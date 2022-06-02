package main

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func RandSlice[T any](slice []T, count int) ([]T, error) {
	if count > len(slice) {
		return nil, errors.New("slice too small")
	}
	idxs := rand.Perm(len(slice))[:count]
	randSlice := make([]T, 0, count)
	for _, idx := range idxs {
		randSlice = append(randSlice, slice[idx])
	}
	return randSlice, nil
}
