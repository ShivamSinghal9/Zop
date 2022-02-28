package testP

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_crud(t *testing.T) {
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
			id:  1,
			ind: info{1, "shivam", "singhal@gmail.com", "sde"},
			query: mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
				AddRow(1, "shivam", "singhal@gmail.com", "sde")),
			err: nil},

		// error
		{
			id:    -1,
			ind:   info{0, "", "", ""},
			query: nil,
			err:   errors.New("INVALID ID")},

		// error
		{
			id:  4,
			ind: info{4, "Changed", "Changed@gmail.com", "CCCCC"},
			query: mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
				AddRow("a", "Changed", "Changed@gmail.com", "CCCCC")),
			err: errors.New("error while scaning")},

		// error
		{
			id:    3,
			ind:   info{},
			query: mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(3).WillReturnError(sql.ErrNoRows),
			err:   sql.ErrNoRows},

		//succes
		// {
		// 	id:    2,
		// 	ind:   info{2, "ishan", "kochar@gmail.com", "sde"},
		// 	query: mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(2).WillReturnRows(rows),
		// 	err:   nil},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			fmt.Println(tc.id)
			w, e := Read(db, tc.id)
			if e != nil && e.Error() != tc.err.Error() {
				t.Errorf("expected error: %v, got: %v", tc.err, e)
			}
			if w != nil && !reflect.DeepEqual(&tc.ind, w) {
				t.Errorf("expected user: %v, got: %v", tc.ind, w)
			}
		})
	}

}
