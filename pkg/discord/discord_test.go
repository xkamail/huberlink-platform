package discord

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {

	t.Run("error should return", func(t *testing.T) {
		c := NewClient("", "", "")
		token, err := c.GetAccessToken(context.Background(), "xxx")
		assert.Error(t, err)
		assert.Emptyf(t, token, "token should be empty")
	})
	t.Run("should return invalid client setting", func(t *testing.T) {
		c := NewClient("123123123123", "xxxx", "")
		token, err := c.GetAccessToken(context.Background(), "xxx")
		assert.Error(t, err)
		assert.Equal(t, "discord: invalid_client", err.Error())
		assert.Emptyf(t, token, "token should be empty")
	})
	t.Run("empty code should return", func(t *testing.T) {
		c := NewClient("xxxx", "xxxx", "https://google.com")
		token, err := c.GetAccessToken(context.Background(), "")
		assert.Error(t, err)
		assert.Equal(t, "code is empty", err.Error())
		assert.Emptyf(t, token, "token should be empty")
	})
}
