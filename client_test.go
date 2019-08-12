package geverse

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeverse_GetArtistPosts(t *testing.T) {
	httpClient := &http.Client{}

	weverseClient := NewGeverse(httpClient, os.Getenv("WEVERSE_TOKEN"))

	resp, err := weverseClient.GetArtistPosts(context.Background(), 3)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Condition(t, func() (success bool) {
		return len(resp.Posts) > 0
	})
}
