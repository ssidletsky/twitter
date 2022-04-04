package entities_test

import (
	"testing"

	"github.com/ssidletsky/esportal-twitter/app/tweets/entities"
)

var (
	users = generateUsers()
)

func BenchmarkAtLeastTwice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entities.AtLeastTwice(users)
	}
}

func BenchmarkExactlyTwice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entities.ExactlyTwice(users)
	}
}

func BenchmarkConstrainedExactlyTwice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entities.ConstrainedExactlyTwice(users)
	}
}

func generateUsers() []entities.User {
	size := 720000
	var age uint8 = 18
	users := make([]entities.User, size)
	for i := range users {
		if i%10000 == 0 {
			age++
		}
		users[i].Age = age
	}
	return users
}
