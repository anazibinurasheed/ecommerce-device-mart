package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveUserOnDatabase(t *testing.T) {
	testCases := []struct {
		name       string
		input      request.SignUpData
		beforeTest func(sqlmock.Sqlmock)
		want       response.UserData
		wantErr    error
	}{
		{
			name: "success sign up",
			input: request.SignUpData{
				UserName: "Anas",
				Email:    "anazibinurasheed@gmail.com",
				Phone:    8590138151,
				Password: "password123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Anas", "anazibinurasheed@gmail.com", 8590138151, "password123").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email", "phone"}).AddRow(1, "Anas", "anazibinurasheed@gmail.com", 8590138151))
			},
			want: response.UserData{
				Id:       1,
				UserName: "Anas",
				Email:    "anazibinurasheed@gmail.com",
				Phone:    8590138151,
			},
			wantErr: nil,
		},

		{
			name: "existing user sign up",
			input: request.SignUpData{
				UserName: "Anas",
				Email:    "anazibinurasheed@gmail.com",
				Phone:    8590138151,
				Password: "password123",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Anas", "anazibinurasheed@gmail.com", 8590138151, "password123").WillReturnError(fmt.Errorf("user already exist"))
			},
			want:    response.UserData{},
			wantErr: fmt.Errorf("user already exist"),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database: %v", err)
			}
			defer mockDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			if err != nil {
				t.Fatalf("Failed to create GORM database: %v", err)
			}

			if tc.beforeTest != nil {
				tc.beforeTest(mockSQL)
			}

			ud := NewUserRepository(gormDB)

			got, err := ud.SaveUserOnDatabase(tc.input)

			assert.Equal(t, tc.wantErr, err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected UserData: %v, but got: %v", tc.want, got)
			}

		})
	}
}
