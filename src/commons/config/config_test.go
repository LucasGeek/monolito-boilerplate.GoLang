package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	os.Setenv("PORT", "4000")

	config := LoadConfig()

	if config.JWTSecret != "test_secret" {
		t.Errorf("Expected JWTSecret to be 'test_secret', but got %s", config.JWTSecret)
	}

	if config.Port != 4000 {
		t.Errorf("Expected Port to be 4000, but got %d", config.Port)
	}
}

func TestGetEnvWithDefaultValue(t *testing.T) {
	os.Unsetenv("TEST_ENV")

	value := getEnv("TEST_ENV", "default_value")
	if value != "default_value" {
		t.Errorf("Expected default_value, but got %s", value)
	}
}

func TestGetEnvWithSetValue(t *testing.T) {
	os.Setenv("TEST_ENV", "set_value")

	value := getEnv("TEST_ENV", "default_value")
	if value != "set_value" {
		t.Errorf("Expected set_value, but got %s", value)
	}
}

func TestGetEnvAsIntWithDefaultValue(t *testing.T) {
	os.Unsetenv("TEST_INT_ENV")

	value := getEnvAsInt("TEST_INT_ENV", 1234)
	if value != 1234 {
		t.Errorf("Expected 1234, but got %d", value)
	}
}

func TestGetEnvAsIntWithSetValue(t *testing.T) {
	os.Setenv("TEST_INT_ENV", "5678")

	value := getEnvAsInt("TEST_INT_ENV", 1234)
	if value != 5678 {
		t.Errorf("Expected 5678, but got %d", value)
	}
}

func TestGetEnvAsIntWithInvalidValue(t *testing.T) {
	os.Setenv("TEST_INT_ENV", "invalid_int")

	value := getEnvAsInt("TEST_INT_ENV", 1234)
	if value != 1234 {
		t.Errorf("Expected 1234, but got %d due to invalid int conversion", value)
	}
}
