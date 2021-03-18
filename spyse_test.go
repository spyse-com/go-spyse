package spyse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient_WithOutBaseURL(t *testing.T) {
	c, err := NewClient("", "", nil)
	require.NoError(t, err)
	require.Equal(t, c.token, "", "http status not equal")
}
