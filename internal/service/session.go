package service

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/006lp/akashchat-api-go/pkg/client"
)

// SessionCache represents a cached session
type SessionCache struct {
	Token     string
	ExpiresAt time.Time
}

// SessionService manages session tokens
type SessionService struct {
	httpClient *client.HTTPClient
	cache      *SessionCache
	mutex      sync.RWMutex
}

// NewSessionService creates a new SessionService instance
func NewSessionService() *SessionService {
	return &SessionService{
		httpClient: client.NewHTTPClient(),
		cache:      nil,
	}
}

// GetSessionToken gets a valid session token (cached or new)
func (s *SessionService) GetSessionToken() (string, error) {
	s.mutex.RLock()
	if s.cache != nil && time.Now().Before(s.cache.ExpiresAt) {
		token := s.cache.Token
		s.mutex.RUnlock()
		return token, nil
	}
	s.mutex.RUnlock()

	// Need to get a new session token
	return s.refreshSessionToken()
}

// refreshSessionToken fetches a new session token from Akash
func (s *SessionService) refreshSessionToken() (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Double-check in case another goroutine already refreshed
	if s.cache != nil && time.Now().Before(s.cache.ExpiresAt) {
		return s.cache.Token, nil
	}

	headers := map[string]string{
		"Referer": "https://chat.akash.network/",
		"Accept":  "*/*",
	}

	resp, err := s.httpClient.Get("https://chat.akash.network/api/auth/session/", headers)
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("session request failed with status: %d", resp.StatusCode)
	}

	// Extract session token from Set-Cookie header
	setCookieHeader := resp.Header.Get("Set-Cookie")
	if setCookieHeader == "" {
		return "", fmt.Errorf("no Set-Cookie header found")
	}

	token, err := s.extractSessionToken(setCookieHeader)
	if err != nil {
		return "", fmt.Errorf("failed to extract session token: %w", err)
	}

	// Cache the token with 1-hour expiration (minus 5 minutes for safety)
	s.cache = &SessionCache{
		Token:     token,
		ExpiresAt: time.Now().Add(55 * time.Minute),
	}

	return token, nil
}

// extractSessionToken extracts session token from Set-Cookie header
func (s *SessionService) extractSessionToken(setCookieHeader string) (string, error) {
	// Example: session_token=0c647105a2175953f14b9f33c3e0100f405667b6c3e2507fb2cc6d0baff1e567; Path=/; ...
	parts := strings.Split(setCookieHeader, ";")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid Set-Cookie header format")
	}

	sessionPart := strings.TrimSpace(parts[0])
	if !strings.HasPrefix(sessionPart, "session_token=") {
		return "", fmt.Errorf("session_token not found in Set-Cookie header")
	}

	token := strings.TrimPrefix(sessionPart, "session_token=")
	if token == "" {
		return "", fmt.Errorf("empty session token")
	}

	return fmt.Sprintf("session_token=%s", token), nil
}
