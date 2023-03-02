package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

const baseURL = `https://discord.com/api/v10`

var httpClient = &http.Client{Timeout: 2 * time.Second}

type Error struct {
	Code         int         `json:"code"`
	Errors       interface{} `json:"errors"`
	Message      string      `json:"error_description"`
	ErrorMessage string      `json:"error"`
}

func (e Error) Error() string {
	if e.ErrorMessage != "" {
		return "discord: " + e.ErrorMessage
	}
	return "discord: " + e.Message
}

type Profile struct {
	ID            snowid.ID `json:"id"`
	Username      string    `json:"username"`
	Avatar        *string   `json:"avatar"`
	Discriminator string    `json:"discriminator"`
	Email         string    `json:"email"`
}

func (p Profile) AvatarURL() string {
	if p.Avatar == nil {
		return ""
	}
	return fmt.Sprintf(`https://cdn.discordapp.com/avatars/%d/%s.webp`, p.ID, *p.Avatar)
}

type Response struct {
	ErrorDescription string `json:"error_description"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type Client interface {
	GetAccessToken(ctx context.Context, code string) (string, error)
	GetProfile(ctx context.Context, accessToken string) (*Profile, error)
}

func NewClient(clientID, clientSecret, redirectUri string) Client {
	return &client{clientID, clientSecret, redirectUri}
}

type client struct {
	clientID     string
	clientSecret string
	redirectUri  string
}

func (c client) GetAccessToken(ctx context.Context, code string) (string, error) {
	if len(code) == 0 || strings.TrimSpace(code) == "" {
		return "", uierr.BadInput("code", "code is empty")
	}
	q := url.Values{
		"client_id":     {c.clientID},
		"client_secret": {c.clientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {c.redirectUri},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+`/oauth2/token`, strings.NewReader(q.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		slog.Error("discord: http client get access token", err)
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		var disErr Error
		if err := json.NewDecoder(res.Body).Decode(&disErr); err != nil {
			slog.Error("discord: decode access token response", err)
			return "", err
		}
		return "", disErr
	}
	var result AccessTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		slog.Error("discord: decode access token response", err)
		return "", err
	}
	return result.AccessToken, nil
}

func (c client) GetProfile(ctx context.Context, accessToken string) (*Profile, error) {
	// https://discord.com/developers/docs/resources/user#get-current-user
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+`/users/@me`, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		slog.Error("discord: http client get profile", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		var disErr Error
		if err := json.NewDecoder(res.Body).Decode(&disErr); err != nil {
			slog.Error("discord: decode profile response", err)
			return nil, err
		}
		return nil, disErr

	}
	var profile Profile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		slog.Error("discord: decode profile response", err)
		return nil, err
	}
	return &profile, nil
}

var _ Client = (*client)(nil)
