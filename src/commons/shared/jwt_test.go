package shared

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTManager_Generate(t *testing.T) {
	manager := NewJWTManager("testSecret", 30*time.Minute, 24*time.Hour)

	userID := uuid.New()
	token, refreshToken, err := manager.Generate(userID)

	if err != nil {
		t.Fatalf("Failed to generate tokens: %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty JWT token but got empty string")
	}

	if refreshToken == "" {
		t.Error("Expected non-empty refresh token but got empty string")
	}
}

func TestJWTManager_Verify(t *testing.T) {
	manager := NewJWTManager("testSecret", 30*time.Minute, 24*time.Hour)

	userID := uuid.New()
	token, _, _ := manager.Generate(userID)

	claims, err := manager.Verify(token)

	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID to be %s but got %s", userID, claims.UserID)
	}
}

func TestJWTManager_VerifyRefreshToken(t *testing.T) {
	manager := NewJWTManager("testSecret", 30*time.Minute, 24*time.Hour)

	userID := uuid.New()
	_, refreshToken, _ := manager.Generate(userID)

	claims, err := manager.VerifyRefreshToken(refreshToken)

	if err != nil {
		t.Fatalf("Failed to verify refresh token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID to be %s but got %s", userID, claims.UserID)
	}
}

func TestJWTManager_InvalidToken(t *testing.T) {
	manager := NewJWTManager("testSecret", 30*time.Minute, 24*time.Hour)

	_, err := manager.Verify("invalidToken")

	if err == nil {
		t.Error("Expected error due to invalid token but got nil")
	}
}

func TestJWTManager_InvalidRefreshToken(t *testing.T) {
	manager := NewJWTManager("testSecret", 30*time.Minute, 24*time.Hour)

	_, err := manager.VerifyRefreshToken("invalidToken")

	if err == nil {
		t.Error("Expected error due to invalid refresh token but got nil")
	}
}
