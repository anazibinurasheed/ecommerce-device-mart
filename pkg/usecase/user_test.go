package usecase

import (
	"testing"

	mockRepo "github.com/anazibinurasheed/project-device-mart/pkg/mock/repoMock"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mockRepo.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(mockRepository)

	testCases := []struct {
		name        string
		input       request.SignUpData
		beforeTest  func(userRepo *mockRepo.MockUserRepository)
		expectedErr error
	}{
		{name: "success sign up",
			input: request.SignUpData{
				UserName: "Anas",
				Email:    "anazibinurasheed@gmail.com",
				Phone:    8590138151,
				Password: gomock.Any().String(),
			},

			beforeTest: func(userRepo *mockRepo.MockUserRepository) {
				userRepo.EXPECT().FindUserByPhone(8590138151).Return(response.UserData{}, nil)
				userRepo.EXPECT().FindUserByEmail("anazibinurasheed@gmail.com").Return(response.UserData{}, nil)
				userRepo.EXPECT().SaveUserOnDatabase(request.SignUpData{
					UserName: "Anas",
					Email:    "anazibinurasheed@gmail.com",
					Phone:    8590138151,
					Password: gomock.Any().String(),
				}).Return(response.UserData{
					Id:       1,
					UserName: "Anas",
					Email:    "anazibinurasheed@gmail.com",
					Phone:    8590138151,
				}, nil)
			},
			expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.beforeTest(mockRepository)
			err := userUseCase.SignUp(tc.input)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
