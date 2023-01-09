package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrInternalServer = errors.New("internal server error")
var ErrEmailEmpty = errors.New("email is empty")
var ErrEmailNotFoundInFB = errors.New("email is not existed in firebase")
var ErrEmailNotFoundInDB = errors.New("email is not existed in DB")
var ErrEmailAlreadyRegistered = errors.New("email is already registered")
var ErrPathParameter = errors.New("wrong path parameter(s)")
var ErrQueryParameter = errors.New("wrong query parameter(s)")
var ErrBodyParameter = errors.New("wrong body parameter(s)")
var ErrGetFirestore = errors.New("read error in firestore")
var ErrSetFirestore = errors.New("write error in firestore")
var ErrGetLocalDB = errors.New("read error in mariadb")
var ErrSetLocalDB = errors.New("write error in mariadb")
var ErrDataNotFoundInDB = errors.New("data not exist in db")
var ErrDBNotFound = errors.New("db instance not found")
var ErrFirestoreAuth = errors.New("firestore Auth Error")
var ErrFirestoreContent = errors.New("firestore Content Error")
var ErrDBContent = errors.New("db Content Error")
var ErrGroupIdCapability = errors.New("parameter groupId invalid")
var ErrCustomerIdCapability = errors.New("parameter customerId invalid")
var ErrSinceCapability = errors.New("parameter since invalid")
var ErrUntilCapability = errors.New("parameter until invalid")
var ErrSinceUntilDurationCapability = errors.New("duration between since and until is too long. must be less than 12 weeks")
var ErrUserNotFound = errors.New("user not found")
var ErrAiEngineNotOnline = errors.New("ai engine not online")
var ErrDataAlreadyExist = errors.New("data already exists")

type (
	httpErrorHandler struct {
		statusCodes map[error]int
	}
)

func NewErrorStatusCodeMaps() map[error]int {
	var errorStatusCodeMaps = make(map[error]int)
	errorStatusCodeMaps[ErrInternalServer] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrEmailEmpty] = http.StatusBadRequest
	errorStatusCodeMaps[ErrEmailNotFoundInFB] = http.StatusNotFound
	errorStatusCodeMaps[ErrEmailNotFoundInDB] = http.StatusNotFound
	errorStatusCodeMaps[ErrEmailAlreadyRegistered] = http.StatusBadRequest
	errorStatusCodeMaps[ErrPathParameter] = http.StatusBadRequest
	errorStatusCodeMaps[ErrBodyParameter] = http.StatusBadRequest
	errorStatusCodeMaps[ErrQueryParameter] = http.StatusBadRequest
	errorStatusCodeMaps[ErrGetFirestore] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrSetFirestore] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrGetLocalDB] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrSetLocalDB] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrDataNotFoundInDB] = http.StatusNotFound
	errorStatusCodeMaps[ErrDBNotFound] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrFirestoreAuth] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrFirestoreContent] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrDBContent] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrGroupIdCapability] = http.StatusBadRequest
	errorStatusCodeMaps[ErrCustomerIdCapability] = http.StatusBadRequest
	errorStatusCodeMaps[ErrSinceCapability] = http.StatusBadRequest
	errorStatusCodeMaps[ErrUntilCapability] = http.StatusBadRequest
	errorStatusCodeMaps[ErrSinceUntilDurationCapability] = http.StatusBadRequest
	errorStatusCodeMaps[ErrUserNotFound] = http.StatusNotFound
	errorStatusCodeMaps[ErrAiEngineNotOnline] = http.StatusServiceUnavailable
	errorStatusCodeMaps[ErrDataAlreadyExist] = http.StatusBadRequest

	return errorStatusCodeMaps
}

func NewHttpErrorHandler(errorStatusCodeMaps map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		statusCodes: errorStatusCodeMaps,
	}
}

func (self *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range self.statusCodes {
		if errors.Is(err, key) {
			return value
		}
	}

	return http.StatusInternalServerError
}

func unwrapRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}

func (self *httpErrorHandler) ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: unwrapRecursive(err).Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
