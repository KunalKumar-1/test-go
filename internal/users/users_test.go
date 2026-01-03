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
		t.Errorf("no error returned when adding first name %v", err)
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

func TestAddUserLastName(t *testing.T) {
	testManager := NewManager()

	testFirstName := "jhon"
	testLastName := ""
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Errorf("no error returned when adding last name %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err == nil {
		t.Errorf("no error returned or invalid email")
	} else {
		expectedErr := "invalid last name: \"\""
		if err.Error() != expectedErr {
			t.Errorf("error mismatch: expected %v, got %v", expectedErr, err)
		}
	}

	if len(testManager.users) > 0 {
		t.Fatalf("bad test manager count: expected 1 user, got %v", len(testManager.users))
	}
}

func TestAddUserDuplicateName(t *testing.T) {
	testManager := NewManager()

	testFirstName := "jhon"
	testLastName := "smith"
	testEmail, err := mail.ParseAddress("foo@bar.com")
	if err != nil {
		t.Errorf("no error returned when adding duplicate name %v", err)
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err != nil {
		t.Errorf("error creating user")
	}

	err = testManager.AddUser(testFirstName, testLastName, testEmail.String())
	if err == nil {
		t.Errorf("error creating duplicate user")
	} else {
		expectedErr := "user already exists"
		if err.Error() != expectedErr {
			t.Errorf("error mismatch: expected %v, got %v", expectedErr, err)
		}
	}

	if len(testManager.users) != 1 {
		t.Errorf("bad test manager count: expected %d user, got %d", 1, len(testManager.users))
	}
}

func TestGetUserByName(t *testing.T) {
	testManager := NewManager()

	err := testManager.AddUser("foo", "bar", "f.foo@bar.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("bari", "foo", "bar@bar.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("barz", "foo", "barz@bar.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}
	err = testManager.AddUser("fozz", "foo", "fooz@bar.com")
	if err != nil {
		t.Fatalf("error adding test user: %v", err)
	}

	tests := map[string]struct {
		first         string
		last          string
		expected      *User
		expectedError error
	}{
		"simple lookup": {
			first:         "foo",
			last:          "bar",
			expected:      &testManager.users[0],
			expectedError: nil,
		},
		"last element lookup": {
			first:         "bari",
			last:          "foo",
			expected:      &testManager.users[3],
			expectedError: nil,
		},
		"no match lookup": {
			first:         "rgdf",
			last:          "rgter",
			expected:      nil,
			expectedError: ErrNoResultFound,
		},
		"partial match lookup": {
			first:         "fozz",
			last:          "fozz",
			expected:      nil,
			expectedError: ErrNoResultFound,
		},
		"empty first name": {
			first:         "",
			last:          "fozz",
			expected:      nil,
			expectedError: ErrNoResultFound,
		},
		"empty last name": {
			first:         "fozz",
			last:          "",
			expected:      nil,
			expectedError: ErrNoResultFound,
		},
	}

	for name, test := range tests {
		result, err := testManager.GetUserByName(test.first, test.last)
		if err != test.expectedError {
			t.Errorf("%s: invalid result:\nexpected: %v\ngot: %v", name, result, test.expected)
			return
		}
	}

}
