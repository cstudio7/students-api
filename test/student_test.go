package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const BASE_URL = "http://localhost:9000"

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/students")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostStudent(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"FirstName": "Jack", "LastName": "Smith", "Age": 28, "School": "UCLA"}`).
		Post(BASE_URL + "/api/students")

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
