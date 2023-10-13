package repo

// import (
// 	"fmt"
// 	"log"
// 	"testing"

// 	sqlmock "github.com/DATA-DOG/go-sqlmock"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
// 	"gopkg.in/go-playground/assert.v1"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func TestSaveUserOnDatabase(t *testing.T) {
// 	testCases := []struct {
// 		name        string
// 		input       request.SignUpData
// 		beforeTest  func(sqlmock.Sqlmock)
// 		want        response.UserData
// 		expectedErr error
// 	}{
// 		{
// 			name: "success sign up",
// 			input: request.SignUpData{
// 				UserName: "Anas",
// 				Email:    "anazibinurasheed@gmail.com",
// 				Phone:    8590138151,
// 				Password: "password123",
// 			},
// 			beforeTest: func(mockSQL sqlmock.Sqlmock) {
// 				expectedQuery := `^INSERT INTO users (.+)$`
// 				mockSQL.ExpectQuery(expectedQuery).WithArgs("Anas", "anazibinurasheed@gmail.com", 8590138151, "password123").
// 					WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email", "phone"}).AddRow(1, "Anas", "anazibinurasheed@gmail.com", 8590138151))
// 			},
// 			want: response.UserData{
// 				ID:       1,
// 				UserName: "Anas",
// 				Email:    "anazibinurasheed@gmail.com",
// 				Phone:    8590138151,
// 			},
// 			expectedErr: nil,
// 		},

// 		{
// 			name: "existing user sign up",
// 			input: request.SignUpData{
// 				UserName: "Anas",
// 				Email:    "anazibinurasheed@gmail.com",
// 				Phone:    8590138151,
// 				Password: "password123",
// 			},
// 			beforeTest: func(mockSQL sqlmock.Sqlmock) {
// 				expectedQuery := `^INSERT INTO users(.+)$`
// 				mockSQL.ExpectQuery(expectedQuery).WithArgs("Anas", "anazibinurasheed@gmail.com", 8590138151, "password123").WillReturnError(fmt.Errorf("user already exist"))
// 			},
// 			want:        response.UserData{},
// 			expectedErr: fmt.Errorf("user already exist"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		// t.Parallel()
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockDB, mockSQL, err := sqlmock.New()
// 			if err != nil {
// 				t.Fatalf("failed to create mock database: %v", err)
// 			}
// 			defer mockDB.Close()

// 			gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 				Conn: mockDB,
// 			}), &gorm.Config{})

// 			if err != nil {
// 				t.Fatalf("failed to create GORM database: %v", err)
// 			}

// 			tc.beforeTest(mockSQL)

// 			ud := NewUserRepository(gormDB)

// 			got, err := ud.CreateUser(tc.input)

// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, got, tc.want)
// 			// if !reflect.DeepEqual(got, tc.want) {
// 			// 	t.Errorf("Expected UserData: %v, but got: %v", tc.want, got)
// 			// }

// 		})
// 	}
// }

// func TestFindUserByEmail(t *testing.T) {
// 	testCases := []struct {
// 		name        string
// 		input       string
// 		beforeTest  func(sqlmock.Sqlmock)
// 		want        response.UserData
// 		expectedErr error
// 	}{{
// 		name:  "finding user by email",
// 		input: "anazibinurasheed@gmail.com",
// 		beforeTest: func(mockSQL sqlmock.Sqlmock) {
// 			expectedQuery := `SELECT \* FROM users WHERE  email \= \$1`
// 			mockSQL.ExpectQuery(expectedQuery).WithArgs("anazibinurasheed@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email", "phone"}).AddRow(1, "anaz", "anazibinurasheed@gmail.com", 8590138151))
// 		},

// 		want: response.UserData{
// 			ID:       1,
// 			UserName: "anaz",
// 			Email:    "anazibinurasheed@gmail.com",
// 			Phone:    8590138151,
// 		},
// 		expectedErr: nil,
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockDB, mockSQL, err := sqlmock.New()
// 			if err != nil {
// 				log.Fatalf("failed to initialize mockDB connection %v", err)
// 			}
// 			defer mockDB.Close()

// 			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
// 				Conn: mockDB,
// 			}), &gorm.Config{})

// 			tc.beforeTest(mockSQL)

// 			ud := NewUserRepository(gormDB)
// 			got, err := ud.FindUserByEmail(tc.input)
// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, tc.want, got)
// 		})
// 	}
// }

// func TestFindUserById(t *testing.T) {
// 	type args struct {
// 		ID int
// 	}
// 	testCases := []struct {
// 		name        string
// 		input       args
// 		beforeTest  func(sqlmock.Sqlmock)
// 		want        response.UserData
// 		expectedErr error
// 	}{{
// 		name:  "success, find user by id",
// 		input: args{ID: 1},
// 		beforeTest: func(mockSQL sqlmock.Sqlmock) {
// 			expectedQuery := `SELECT \* FROM users WHERE  ID\= \$1`
// 			mockSQL.ExpectQuery(expectedQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email", "phone"}).AddRow(1, "anaz", "anazibinurasheed@gmail.com", 8590138151))

// 		},
// 		want: response.UserData{
// 			ID:       1,
// 			UserName: "anaz",
// 			Email:    "anazibinurasheed@gmail.com",
// 			Phone:    8590138151,
// 		},
// 		expectedErr: nil,
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockDB, mockSQL, err := sqlmock.New()
// 			if err != nil {
// 				log.Fatalf("failed to initialize mockDB connection %#v", err)
// 			}
// 			gormDB, _ := gorm.Open(postgres.New(
// 				postgres.Config{Conn: mockDB},
// 			), &gorm.Config{})

// 			tc.beforeTest(mockSQL)
// 			ud := NewUserRepository(gormDB)
// 			got, err := ud.FindUserByID(tc.input.ID)

// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, got, tc.want)

// 		})

// 	}
// }
