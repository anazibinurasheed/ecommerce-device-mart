package helper

import (
	"net/http"
	"strconv"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	errJSONBindingMsg               = "Failed to bind JSON inputs from request"
	errBindingPageInfo              = "Failed to bind page info from request"
	errBindParamFailMsg             = "Failed to retrieve param from URL"
	errBindFormValueMsg             = "Failed to bind form values from request"
	errInputDoesntMeetValidationMsg = "Failed, input does not meet validation criteria"
	errFailedGetIDFromContextMsg    = "Failed to retrieve user id from context"
)

type SubHandler struct{}

// NewSubHandler return a subHandler which implements the
func NewSubHandler() *SubHandler {
	return &SubHandler{}
}

// ParamInt retrieves the param by name and return as int.
// If error occurred while retrieving the param, it will return false and writes appropriate response to the header.
//
// Swagger
//
//	@Failure	400	{object}	response.Response	"Failed to retrieve param from URL"
func (s *SubHandler) ParamInt(c *gin.Context, name string) (param int, ok bool) {
	val, err := strconv.Atoi(c.Param(name))

	if IsErr(err) {
		errParamResp(c, err)
		return -1, false
	}

	return val, true
}

// errParamResp writes the appropriate response header.
func errParamResp(c *gin.Context, err error) {

	response := response.ResponseMessage(http.StatusBadRequest, errBindParamFailMsg, nil, err.Error())
	c.JSON(http.StatusBadRequest, response)
}

// BindRequest binds and validates the input based on the Validate and Binding tags.
// Any error occurred while binding or validating it will return false and it writes appropriate response to the header.
//
// # Swagger
//
//	@Failure	400	{object}	response.Response	"Failed to bind JSON inputs from request"
//
//	@Failure	400	{object}	response.Response	"Failed, input does not meet validation criteria"
func (s *SubHandler) BindRequest(c *gin.Context, obj any) bool {

	err := c.ShouldBindJSON(obj)
	if IsErr(err) {
		response := response.ResponseMessage(http.StatusBadRequest, errJSONBindingMsg, nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, &response)
		return false
	}

	return validateStruct(c, obj)
}

// validateStruct will validate the struct fields according to validate tags and if ok it will return true else it will return false.
// And it will also write appropriate response to the header.
func validateStruct(c *gin.Context, obj any) bool {

	if err := validater(obj); err != nil {
		response := response.ResponseMessage(http.StatusNotAcceptable, errInputDoesntMeetValidationMsg, nil, err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, &response)
		return false
	}
	return true

}

// validater validate struct using validate tags, if the validation fails it will return false else true
func validater(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}

// errPageInfoResp writes the appropriate response header.
func errPageInfoResp(c *gin.Context, err error) {

	response := response.ResponseMessage(http.StatusBadRequest, errBindingPageInfo, nil, err.Error())
	c.JSON(http.StatusBadRequest, response)
}

// GetPageNCount retrieve page and count query from the request.
// It returns true if the process success.
// Error occurred while processing it writes appropriate response to the header.
//
// swagger
//
//	@Failure	400	{object}	response.Response	"Failed to bind page info from request"
func (s *SubHandler) GetPageNCount(c *gin.Context) (page int, count int, ok bool) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errPageInfoResp(c, err)
		return 0, 0, false
	}

	count, err = strconv.Atoi(c.Query("count"))
	if err != nil {
		errPageInfoResp(c, err)
		return 0, 0, false
	}

	ok = true
	return
}

// GetUserID retrieves the Users Id from the context.
// If any error occurred it will return appropriate response to header and return !ok.
// swagger
//	@Failure	500	{object}	response.Response	"Failed to retrieve user id from context"
func (s *SubHandler) GetUserID(c *gin.Context) (ID int, ok bool) {
	idStr := c.GetString("userID")

	ID, err := strconv.Atoi(idStr)

	if err != nil {
		errGetIDResp(c, err)
		return
	}

	ok = true
	return
}

// errPageInfoResp writes the appropriate response header.
func errGetIDResp(c *gin.Context, err error) {

	response := response.ResponseMessage(http.StatusInternalServerError, errFailedGetIDFromContextMsg, nil, err.Error())
	c.JSON(http.StatusBadRequest, response)
}
