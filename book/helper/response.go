package helper

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type ResponseModel struct {
	Error          error       `json:"-"`    // low-level runtime error
	HTTPStatusCode int         `json:"code"` // http response status code
	HTTPStatusText string      `json:"-"`    // http response status code
	Success        bool        `json:"success"`
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
}

func (e *ResponseModel) Render(w http.ResponseWriter, r *http.Request) error {
	if !e.Success {
		e.Message = fmt.Sprintf("%s: %s", e.HTTPStatusText, e.Message)
	}

	render.Status(r, e.HTTPStatusCode)

	return nil
}

func statusText(code int) string {
	statusText := "Unknown Error"

	switch code {
	case 422:
		statusText = "Error Rendering Response"
	case 400:
		statusText = "Bad Request"
	case 401:
		statusText = "Unauthorized"
	case 403:
		statusText = "Forbidden"
	case 404:
		statusText = "Not Found"
	case 405:
		statusText = "Method Not Allowed"
	case 500:
		statusText = "Internal Server Error"
	case 502:
		statusText = "Bad Gateway"
	case 503:
		statusText = "Server Unavailable"
	}

	return statusText
}

func ResponseSuccess(data interface{}) render.Renderer {
	res := &ResponseModel{
		Data:           data,
		Success:        true,
		HTTPStatusCode: 200,
	}

	return res
}

func ResponseError(errCode int, err error) render.Renderer {
	return &ResponseModel{
		Success:        false,
		Data:           nil,
		HTTPStatusCode: errCode,
		HTTPStatusText: statusText(errCode),
		Error:          err,
		Message:        err.Error(),
	}
}
