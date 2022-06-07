package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

// RandSlice returns a new slice with a random collection
// of "count" items from the original slice.
func RandSlice[T any](slice []T, count int) []T {
	if count > len(slice) {
		count = len(slice)
	}
	idxs := rand.Perm(len(slice))[:count]
	randSlice := make([]T, 0, count)
	for _, idx := range idxs {
		randSlice = append(randSlice, slice[idx])
	}
	return randSlice
}

// FilterSlice is a fast item slice filter based on keys.
//
// keyFunc is used to extract the key from the items.
// If the key is present in the key slice, the item will be part of the result.
func FilterSlice[I, K comparable](items []I, keys []K, keyFunc func(I) K) []I {
	kSet := NewSet(keys...)
	result := make([]I, 0, len(keys))
	for _, item := range items {
		if kSet.Contains(keyFunc(item)) {
			result = append(result, item)
		}
	}
	return result
}
