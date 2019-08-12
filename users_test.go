package geverse

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeverse_GetMe(t *testing.T) {
	httpClient := &http.Client{}

	weverseClient := NewGeverse(httpClient, os.Getenv("WEVERSE_TOKEN"))

	resp, err := weverseClient.GetMe(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Condition(t, func() (success bool) {
		return len(resp.MyCommunities) > 0
	})
}
