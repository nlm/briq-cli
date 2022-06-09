package main

import (
	"github.com/nlm/briq-cli/briq"
)

// FilterOnUsername provides the keyFunc to filter on Briq Username
func BriqUserKey(user briq.User) string {
	return user.Username
}
