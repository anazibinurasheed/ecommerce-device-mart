package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	mock "github.com/anazibinurasheed/project-device-mart/pkg/mock/usecaseMock"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
// 	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
// 	"github.com/gin-gonic/gin"
// 	gomock "github.com/golang/mock/gomock"
// 	"gopkg.in/go-playground/assert.v1"
// )

// func TestUserSignup(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockUseCase := mock.NewMockUserUseCase(ctrl)
// 	userHandler := NewUserHandler(mockUseCase)
// 	testCases := []struct {
// 		name        string
// 		response    response.Response
// 		beforeTest  func(userUseCase *mock.MockUserUseCase)
// 		want        error
// 		expectedErr error
// 	}{{
// 		name:     "success sign up",
// 		response: response.Response{StatusCode: 200, Message: "Success, account created", Data: nil, Error: nil},
// 		beforeTest: func(userUseCase *mock.MockUserUseCase) {
// 			userUseCase.EXPECT().SignUp(request.SignUpData{
// 				UserName: "Anaz",
// 				Email:    "anazibinurasheed@gmail.com",
// 				Phone:    8590138151,
// 				Password: "123456789",
// 				UUID:     "key",
// 			}).Return(nil).Times(1)

// 		},
// 		want:        nil,
// 		expectedErr: nil,
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			router := gin.Default()
// 			URL := "/sign-up"
// 			router.POST(URL, userHandler.UserSignUp)

// 			phoneDataMutex.Lock()
// 			uid := "key"
// 			phoneDataMap[uid] = "8590138151"
// 			phoneDataMutex.Unlock()

// 			data := `{"username":"Anaz","email":"anazibinurasheed@gmail.com",
// 			"password":"123456789","uuid":"key"}`

// 			req, _ := http.NewRequest("POST", URL, strings.NewReader(data))
// 			req.Header.Set("Content-Type", "application/json")

// 			responseRecorder := httptest.NewRecorder()

// 			tc.beforeTest(mockUseCase)

// 			router.ServeHTTP(responseRecorder, req)

// 			var actual response.Response
// 			json.Unmarshal(responseRecorder.Body.Bytes(), &actual)

// 			assert.Equal(t, tc.response.StatusCode, actual.StatusCode)
// 			assert.Equal(t, tc.response.Message, actual.Message)
// 			assert.Equal(t, tc.response.Error, actual.Error)
// 		})
// 	}
// }

// func TestEditUserName(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockUseCase := mock.NewMockUserUseCase(ctrl)
// 	userHandler := NewUserHandler(mockUseCase)

// 	testCases := []struct {
// 		name        string
// 		response    response.Response
// 		beforeTest  func(userUseCase *mock.MockUserUseCase)
// 		expectedErr error
// 	}{{
// 		name:     "success, username changed",
// 		response: response.ResponseMessage(200, "success, username has been changed", nil, nil),
// 		beforeTest: func(userUserCase *mock.MockUserUseCase) {
// 			userUserCase.EXPECT().UpdateUserName("anaz", 1).Return(nil).Times(1)
// 		},
// 		expectedErr: nil,
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			router := gin.Default()
// 			URL := "/profile/edit-username"

// 			router.POST(URL, userHandler.EditUserName)
// 			data := `{"name:"anaz"}`
// 			req, _ := http.NewRequest("POST", URL, strings.NewReader(data))

// 			req.Header.Set("Content-Type", "application/json")
// 			responseRecorder := httptest.NewRecorder()
// 			tc.beforeTest(mockUseCase)
// 			router.ServeHTTP(responseRecorder, req)

// 			var actual response.Response

// 			json.Unmarshal(responseRecorder.Body.Bytes(), &actual)

// 			assert.Equal(t, tc.response.StatusCode, actual.StatusCode)

// 			assert.Equal(t, tc.response.Message, actual.Message)
// 			assert.Equal(t, tc.response.Error, actual.Error)

// 		})

// 	}

// }
