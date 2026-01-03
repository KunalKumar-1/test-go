package users

import (
	"errors"
	"net/mail"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	testManager := NewManager()

	testFirstName := "jhon"
	testLastName := "smith"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Fatalf("failed to parse email: %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.Address)
	if err != nil {
		t.Fatalf("failed to add user: %v", err)
	}

	if len(testManager.users) != 1 {
		t.Fatalf("failed to add user: expected 1 user, got %v", len(testManager.users))
		if len(testManager.users) < 1 {
			t.Fatal()
		}
	}

	expectedUser := User{
		FirstName: testFirstName,
		LastName:  testLastName,
		Email:     *testEmail,
	}

	founduser := testManager.users[0]
	if !reflect.DeepEqual(expectedUser, founduser) {
		t.Fatalf("failed to add user: expected %v, got %v",
			expectedUser, founduser)
	}
}

func TestAddUserInvalidEmail(t *testing.T) {
	testManager := NewManager()

	testFirstName := "jhon"
	testLastName := "smith"
	testEmail := "foobar"

	err := testManager.AddUser(testFirstName, testLastName, testEmail)
	if err == nil {
		t.Errorf("no error returned when adding invalid email")
	} else {
		expectedErr := errors.New("invalid email: foobar")
		if err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch: expected %v, got %v", expectedErr, err)
		}
	}

	if len(testManager.users) > 0 {
		t.Fatalf("bad test manager count: expected 1 user, got %v", len(testManager.users))
	}
}

func TestAddUserFirstName(t *testing.T) {
	testManager := NewManager()

	testFirstName := ""
	testLastName := "smith"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Errorf("no error returned when adding invalid email %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err == nil {
		t.Errorf("no error returned or invalid email")
	} else {
		expectedErr := "invalid first name: \"\""
		if err.Error() != expectedErr {
			t.Errorf("error mismatch: expected %v, got %v", expectedErr, err)
		}
	}

	if len(testManager.users) > 0 {
		t.Fatalf("bad test manager count: expected 1 user, got %v", len(testManager.users))
	}
}
