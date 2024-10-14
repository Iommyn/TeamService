package integration

import (
	"testing"

	"github.com/levigross/grequests"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAPI(t *testing.T) {
	resp, err := grequests.Get("https://www.google.com/", nil)

	if resp.StatusCode == 200 {
		assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")
	} else {
		assert.Equal(t, 204, resp.StatusCode, "Expected status code 204, got %d", resp.StatusCode)
	}

	assert.NoError(t, err)
}
