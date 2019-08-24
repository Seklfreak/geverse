package geverse

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeverse_GetCommunity(t *testing.T) {
	httpClient := &http.Client{}

	weverseClient := NewGeverse(httpClient, os.Getenv("WEVERSE_TOKEN"))

	resp, err := weverseClient.GetCommunity(context.Background(), 3)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotZero(t, resp.ID)
	assert.Condition(t, func() (success bool) {
		return len(resp.Artists) > 0
	})
	assert.Condition(t, func() (success bool) {
		return len(resp.Tabs) > 0
	})
}
