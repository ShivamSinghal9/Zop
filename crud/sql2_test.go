package testP

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_DEL(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testCases := []struct {
		id    int
		ind   info
		query interface{}
		err   error
	}{
		// success
		{
			id:    1,
			ind:   info{1, "shivam", "singhal@gmail.com", "sde"},
			query: mock.ExpectExec("DELETE FROM employee WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			err:   nil},

		// error
		{
			id:    0,
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("errror while deleting")},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := Delete(db, tc.id)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}

}

func Test_UPD(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testCases := []struct {
		k     int
		id    int
		ind   info
		query interface{}
		err   error
	}{
		// success
		{
			k:     0,
			id:    1,
			ind:   info{1, "SSSS", "singhal@gmail.com", "sde"},
			query: mock.ExpectExec("update employee set name=? where id=?").WithArgs("SSSS", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			err:   nil},

		// success
		{
			k:     1,
			id:    1,
			ind:   info{1, "SSSS", "SSSS", "sde"},
			query: mock.ExpectExec("update employee set email=? where id=?").WithArgs("SSSS", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			err:   nil},

		// success

		{
			k:     2,
			id:    1,
			ind:   info{1, "SSSS", "SSSS", "SSSS"},
			query: mock.ExpectExec("update employee set role=? where id=?").WithArgs("SSSS", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			err:   nil},

		// error
		{
			k:     0,
			id:    0,
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("name not changed")},

		// error
		{
			k:     1,
			id:    0,
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("email not changed")},

		// error
		{
			k:     2,
			id:    0,
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("role not changed")},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := Update(db, tc.id, tc.k, tc.ind)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}

}
