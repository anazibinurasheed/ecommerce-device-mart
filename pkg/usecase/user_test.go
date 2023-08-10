package usecase

// import (
// 	"fmt"
// 	"testing"

// 	mockRepo "github.com/anazibinurasheed/project-device-mart/pkg/mock/repoMock"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
// 	"github.com/golang/mock/gomock"
// )

// func TestSignup(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	fmt.Println("ctrl in usecase testing:", ctrl)
// 	defer ctrl.Finish()

// 	mockRepository := mockRepo.NewMockUserRepository(ctrl)
// 	fmt.Println("mock repository in usecase testing:", mockRepository)

// 	userUseCase := NewUserUseCase(mockRepository)
// 	fmt.Println("mock userUseCase in the usecase testing:", userUseCase)

// 	type args struct {
// 		args request.SignUpData
// 	}

// testCases:=map[string]struct{
// 	args args
// 	beforeTest func(userRepo *mockRepo.MockUserRepository)
// 	want       error
// }{
// 	"success sign up":{
// 	args: request.SignUpData{
// 		UserName:"Anas" ,
// 		Email: "anazibinurasheed@gmail.com",
// 		Phone: 8590138151,
// 	},

// 	}
// }
