package models

import "testing"

func TestUserRegister(t *testing.T) {
	var (
		username = "test_user"
		password = "test_password"
		email    = "test_email"
	)

	u, err := NewUser(username, password, email)
	if err != nil {
		t.Error("Creating user", err)
		t.FailNow()
	}

	if u.Username != username {
		t.Errorf("Username not the same (expected:%v, got:%v)", username, u.Username)
	}

	if u.Password != password {
		t.Errorf("Password not the same (expected:%v, got:%v)", password, u.Password)
	}

	if u.Email != email {
		t.Errorf("Email not the same (expected:%v, got:%v)", email, u.Email)
	}
}
