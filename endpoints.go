package geverse

import (
	"fmt"
)

var (
	endpointBaseURL     = "https://weverseapi.weverse.io/api/"
	endpointArtistPosts = func(communityID int64) string {
		return endpointBaseURL + fmt.Sprintf("v1/stream/community/%d/artistTab", communityID)
	}
	endpointMediaPosts = func(communityID int64) string {
		return endpointBaseURL + fmt.Sprintf("v1/stream/community/%d/mediaTab", communityID)
	}
)
