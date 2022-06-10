package main

import (
	"github.com/nlm/briq-cli/briq"
)

const (
	briqsMaxAmount uint = 5
)

// CapBriqAmount ensures the number of briqs stays sane
func CapBriqAmount(amount uint) uint {
	if amount > briqsMaxAmount {
		return briqsMaxAmount
	}
	return amount
}

// FilterOnUsername provides the keyFunc to filter on Briq Username
func BriqUserKey(user briq.User) string {
	return user.Username
}
