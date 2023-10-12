package shared

import (
	"testing"
)

func TestArgon2Manager_HashPassword(t *testing.T) {
	manager := NewArgon2Manager()

	password := "mypassword"
	hash, err := manager.HashPassword(password)

	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hash) == 0 {
		t.Error("Expected non-empty hash but got empty string")
	}
}

func TestArgon2Manager_VerifyPassword(t *testing.T) {
	manager := NewArgon2Manager()

	password := "mypassword"
	hash, err := manager.HashPassword(password)

	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	matched, err := manager.VerifyPassword(password, hash)

	if err != nil {
		t.Fatalf("Failed to verify password: %v", err)
	}

	if !matched {
		t.Error("Expected password to match hash, but they did not match")
	}
}

func TestArgon2Manager_VerifyPasswordWithInvalidHash(t *testing.T) {
	manager := NewArgon2Manager()

	// Providing invalid hash format
	matched, err := manager.VerifyPassword("mypassword", "invalid$hash")

	if err == nil {
		t.Error("Expected error due to invalid hash format but got nil")
	}

	if matched {
		t.Error("Expected password not to match hash, but they matched")
	}
}

func TestArgon2Manager_VerifyPasswordWithIncorrectPassword(t *testing.T) {
	manager := NewArgon2Manager()

	password := "mypassword"
	hash, _ := manager.HashPassword(password)

	// Try to verify with incorrect password
	matched, err := manager.VerifyPassword("wrongpassword", hash)

	if err == nil {
		t.Fatal("Expected an error due to incorrect password but got nil")
	}

	if matched {
		t.Error("Expected password not to match hash, but they matched")
	}
}
