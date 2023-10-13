package usecase

// import (
// 	"fmt"
// 	"testing"

// 	mockRepo "github.com/anazibinurasheed/project-device-mart/pkg/mock/repoMock"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
// 	"github.com/golang/mock/gomock"
// 	"gopkg.in/go-playground/assert.v1"
// )

// // func TestSignUp(t *testing.T) {
// // 	ctrl := gomock.NewController(t)
// // 	defer ctrl.Finish()

// // 	mockRepository := mockRepo.NewMockUserRepository(ctrl)
// // 	userUseCase := NewUserUseCase(mockRepository)

// // 	testCases := []struct {
// // 		name        string
// // 		input       request.SignUpData
// // 		beforeTest  func(userRepo *mockRepo.MockUserRepository)
// // 		expectedErr error
// // 	}{
// // 		{name: "success sign up",
// // 			input: request.SignUpData{
// // 				UserName: "Anaz",
// // 				Email:    "anazibinurasheed@gmail.com",
// // 				Phone:    8590138151,
// // 				Password: "123456789",
// // 			},

// // 			beforeTest: func(userRepo *mockRepo.MockUserRepository) {
// // 				userRepo.EXPECT().FindUserByPhone(8590138151).Return(response.UserData{}, nil)
// // 				userRepo.EXPECT().FindUserByEmail("anazibinurasheed@gmail.com").Return(response.UserData{}, nil)
// // 				userRepo.EXPECT().CreateUser(request.SignUpData{
// // 					UserName: "Anaz",
// // 					Email:    "anazibinurasheed@gmail.com",
// // 					Phone:    8590138151,
// // 					Password: "123456789",
// // 				}).Return(response.UserData{
// // 					ID:       1,
// // 					UserName: "Anaz",
// // 					Email:    "anazibinurasheed@gmail.com",
// // 					Phone:    8590138151,
// // 				}, nil)
// // 			},
// // 			expectedErr: nil},
// // 	}

// // 	for _, tc := range testCases {
// // 		t.Run(tc.name, func(t *testing.T) {
// // 			tc.beforeTest(mockRepository)
// // 			err := userUseCase.SignUp(tc.input)
// // 			assert.Equal(t, tc.expectedErr, err)
// // 		})
// // 	}
// // }

// func TestUpdateUserName(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepository := mockRepo.NewMockUserRepository(ctrl)
// 	userUseCase := NewUserUseCase(mockRepository)

// 	type args struct {
// 		name string
// 		ID   int
// 	}
// 	testCases := []struct {
// 		name        string
// 		input       args
// 		beforeTest  func(*mockRepo.MockUserRepository)
// 		expectedErr error
// 	}{{
// 		name: "success, username",
// 		input: args{
// 			name: "anaz",
// 			ID:   1,
// 		},
// 		beforeTest: func(userRepo *mockRepo.MockUserRepository) {
// 			userRepo.EXPECT().UpdateUserName("anaz", 1).Return(response.UserData{
// 				ID:       1,
// 				UserName: "anaz",
// 				Email:    "anazibinurasheed@gmail.com",
// 			}, nil).Times(1)

// 		},
// 		expectedErr: nil,
// 	}, {
// 		name: "success, username",
// 		input: args{
// 			name: "anaz",
// 			ID:   1,
// 		},
// 		beforeTest: func(userRepo *mockRepo.MockUserRepository) {
// 			userRepo.EXPECT().UpdateUserName("anaz", 1).Return(response.UserData{
// 				ID:       1,
// 				UserName: "Anaz",
// 				Email:    "anazibinurasheed@gmail.com",
// 			}, nil).Times(1)

// 		},
// 		expectedErr: fmt.Errorf("Failed to update username"),
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			tc.beforeTest(mockRepository)
// 			err := userUseCase.UpdateUserName(tc.input.name, tc.input.ID)

// 			assert.Equal(t, err, tc.expectedErr)

// 		})
// 	}

// }
