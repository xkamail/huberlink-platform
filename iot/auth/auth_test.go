package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/pkg/discord"
	"github.com/xkamail/huberlink-platform/pkg/tm"
)

func TestSignInWithDiscord(t *testing.T) {
	testm, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testm.Cleanup()
	assert.NoError(t, testm.CreateTable())
	ctx := testm.Ctx()
	mockProfile := discord.Profile{
		ID:            "395561779368951811",
		Username:      "xKamaiL",
		Avatar:        nil,
		Discriminator: "4882",
		Email:         "doamin@email.com",
	}
	mockClient := mockDiscord{
		&mockProfile,
	}
	t.Parallel()
	t.Run("should return error when code is empty", func(t *testing.T) {
		_, err := SignInWithDiscord(ctx, mockClient, &SignInWithDiscordParam{})
		assert.Error(t, err)
	})
	t.Run("should return error when code is invalid", func(t *testing.T) {
		_, err := SignInWithDiscord(ctx, mockClient, &SignInWithDiscordParam{
			Code: "invalid",
		})
		assert.Error(t, err)
	})
	t.Run("sign in success", func(t *testing.T) {
		res, err := SignInWithDiscord(ctx, mockClient, &SignInWithDiscordParam{
			Code: "correct",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.Token)
		assert.NotEmpty(t, res.RefreshToken)
		t.Run("login again", func(t *testing.T) {
			res, err := SignInWithDiscord(ctx, mockClient, &SignInWithDiscordParam{
				Code: "correct",
			})
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res.Token)
			assert.NotEmpty(t, res.RefreshToken)
		})
	})
}

type mockDiscord struct {
	mockProfile *discord.Profile
}

func (m mockDiscord) GetAccessToken(ctx context.Context, code string) (string, error) {
	if code == "correct" {
		return "accessToken", nil
	}
	return "", discord.Error{
		Code:    0,
		Errors:  nil,
		Message: "invalid code",
	}
}

func (m mockDiscord) GetProfile(ctx context.Context, accessToken string) (*discord.Profile, error) {
	if accessToken == "accessToken" {
		return m.mockProfile, nil
	}
	return nil, discord.Error{
		Code:    0,
		Errors:  nil,
		Message: "invalid access token",
	}
}

var _ discord.Client = (*mockDiscord)(nil)
