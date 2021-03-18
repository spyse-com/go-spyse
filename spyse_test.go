package spyse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient_WithOutToken(t *testing.T) {
	c, err := NewClient(nil, "")
	require.NoError(t, err)
	require.Equal(t, c.token, "", "http status not equal")
}

func TestNewClient_WithServices(t *testing.T) {
	c, err := NewClient(nil, "test_API_token")
	if err != nil {
		t.Errorf("Got an error: %s", err)
	}
	if c.AS == nil {
		t.Error("Autonomous System service do not provided")
	}
}
