package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock "github.com/anazibinurasheed/project-device-mart/pkg/mock/usecaseMock"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockUserUseCase(ctrl)
	userHandler := NewUserHandler(mockUseCase)

	testCases := []struct {
		name        string
		response    response.Response
		beforeTest  func(userUseCase *mock.MockUserUseCase)
		want        error
		expectedErr error
	}{{
		name:     "success signup",
		response: response.Response{StatusCode: 200, Message: "Success, account created", Data: nil, Error: nil},
		beforeTest: func(userUseCase *mock.MockUserUseCase) {
			userUseCase.EXPECT().SignUp(request.SignUpData{
				UserName: "Anaz",
				Email:    "anazibinurasheed@gmail.com",
				Phone:    8590138151,
				Password: "123456789",
				UUID:     "key",
			}).Return(nil).Times(1)

		},
		want:        nil,
		expectedErr: nil,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			router := gin.Default()
			URL := "/sign-up"
			router.POST(URL, userHandler.UserSignUp)

			phoneDataMutex.Lock()
			uid := "key"
			phoneDataMap[uid] = "8590138151"
			phoneDataMutex.Unlock()

			data := `{"username":"Anaz","email":"anazibinurasheed@gmail.com",
			"password":"123456789","uuid":"key"}`

			req, _ := http.NewRequest("POST", URL, strings.NewReader(data))
			req.Header.Set("Content-Type", "application/json")

			responseRecorder := httptest.NewRecorder()

			tc.beforeTest(mockUseCase)

			router.ServeHTTP(responseRecorder, req)

			var actual response.Response
			json.Unmarshal(responseRecorder.Body.Bytes(), &actual)

			assert.Equal(t, tc.response.StatusCode, actual.StatusCode)
			assert.Equal(t, tc.response.Message, actual.Message)
			assert.Equal(t, tc.response.Error, actual.Error)
		})
	}
}

func TestEditUserName(t *testing.T) {
	// Creating new controller
	ctrl := gomock.NewController(t)
	// Closing the ctrl not necessary
	defer ctrl.Finish()

	//creating new mockUseCase instance by passing the ctrl to the constructor and the returning value has all the methods of userUseCase
	mockUseCase := mock.NewMockUserUseCase(ctrl)
	// Passing it into userHandler constructor
	userHandler := NewUserHandler(mockUseCase)

	testCases := []struct {
		name        string
		response    response.Response
		beforeTest  func(userUseCase *mock.MockUserUseCase)
		expectedErr error
	}{{
		name:     "success, username changed",
		response: response.ResponseMessage(200, "success, username has been changed", nil, nil),
		// Setting up expectations
		beforeTest: func(userUserCase *mock.MockUserUseCase) {
			userUserCase.EXPECT().UpdateUserName("anaz", 1).Return(nil).Times(1)
		},
		expectedErr: nil,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Creating a new Gin engine instance
			router := gin.Default()

			// Defining the URL path for the request
			URL := "/profile/edit-username"

			// Setting up the route path and handler
			router.POST(URL, func(c *gin.Context) {
				c.Set("userId", "1")
				c.Next()
			}, userHandler.EditUserName)

			// Input data for the request in JSON format
			data := `{"name":"anaz"}`

			// Creating a new HTTP request "req"
			req, _ := http.NewRequest("POST", URL, strings.NewReader(data))

			// Setting the request header for JSON content
			req.Header.Set("Content-Type", "application/json")

			// Creating a recorder to capture the HTTP response
			responseRecorder := httptest.NewRecorder()

			// Setting up expectations using the "beforeTest" function
			tc.beforeTest(mockUseCase)

			// Handling the request with the Gin router
			router.ServeHTTP(responseRecorder, req)

			var actual response.Response

			// Parsing the response body to make check
			err := json.Unmarshal(responseRecorder.Body.Bytes(), &actual)

			if err != nil {
				log.Fatalf("failed to parse the response body : %s", err)
			}
			assert.Equal(t, tc.response.StatusCode, actual.StatusCode)
			assert.Equal(t, tc.response.Message, actual.Message)
			assert.Equal(t, tc.response.Error, actual.Error)
		})
	}

}
