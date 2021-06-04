package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariefsn/book-store/auth/controllers"
	"github.com/stretchr/testify/assert"
)

func buildClaims() string {
	id := 0
	email := "administrator@mail.com"

	claimsString := fmt.Sprintf("%v*%s", id, email)
	data := []byte(claimsString)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded
}

type TestCase struct {
	name       string
	method     string
	headers    map[string]string
	want       string
	statusCode int
	message    string
}

func TestRootHandler(t *testing.T) {
	assert := assert.New(t)

	tc := []TestCase{
		{
			name:       "should return api info",
			method:     http.MethodGet,
			want:       `{"code":200,"success":true,"data":"Hi, Welcome to Auth Service Version 1","message":""}`,
			statusCode: 200,
			message:    "response body not contains api info",
		},
		{
			name:   "should have claims",
			method: http.MethodGet,
			headers: map[string]string{
				"Claims": buildClaims(),
			},
			want:       buildClaims(),
			statusCode: 200,
			message:    "claims should be exists",
		},
	}

	for i, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(c.method, "/", nil)
			res := httptest.NewRecorder()

			if c.headers["Claims"] != "" {
				req.Header.Add("Claims", c.headers["Claims"])
			}

			controllers.NewAuthController().Hi(res, req)

			assert.Equal(c.statusCode, res.Result().StatusCode, fmt.Sprintf("status code should be %v instead of %v", c.statusCode, res.Result().StatusCode))

			if i == 0 {
				assert.Equal(c.want, strings.TrimSpace(res.Body.String()), c.message)
			} else {
				assert.Equal(c.want, req.Header.Get("Claims"), c.message)
			}
		})
	}
}
