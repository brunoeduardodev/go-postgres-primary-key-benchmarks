package fakedata

import "github.com/go-faker/faker/v4"

type FakeUser struct {
	Name      string `faker:"name"`
	Email     string `faker:"email"`
	Birthdate string `faker:"date"`
}

func GenerateFakeUsers(amount int) []FakeUser {
	users := make([]FakeUser, amount)
	for i := 0; i < amount; i++ {
		user := FakeUser{}
		faker.FakeData(&user)

		users[i] = user
	}

	return users
}
