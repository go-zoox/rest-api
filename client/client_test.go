package client

import (
	"testing"
)

func TestClient(t *testing.T) {
	c := New("http://localhost:8080/")
	type User struct {
		nickname string
		age      int
	}
	users := CreateResource[User](c, "users")

	users.List(0, 10)
	users.Create(&User{
		nickname: "Zero",
		age:      18,
	})
	users.Retrieve("666")
	users.Update("666", &User{
		nickname: "Zero2",
		age:      16,
	})
	users.Delete("666")
}
