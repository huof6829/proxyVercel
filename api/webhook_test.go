package handler

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseUrl(t *testing.T) {
	require := require.New(t)

	targetURL := "http://ec2-52-77-241-219.ap-southeast-1.compute.amazonaws.com:8443"
	// targetURL := "http://localhost:8888"

	url, err := url.Parse(targetURL)
	require.NoError(err)

	fmt.Printf("url.Host: %+v\n", url.Host)
}
