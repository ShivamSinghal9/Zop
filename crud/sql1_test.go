package testP

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_INS(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testCases := []struct {
		ind   info
		query interface{}
		err   error
	}{
		// success
		{
			ind:   info{1, "shivam", "singhal@gmail.com", "sde"},
			query: mock.ExpectExec("INSERT INTO employee VALUES(id,name,email,role) VALUES(?,?,?,?)").WillReturnResult(sqlmock.NewResult(1, 1)),
			err:   nil},

		// error
		{
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("errror while inserting")},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := Insert(db, tc.ind)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}

}
