package shared

import (
	"errors"
	"github.com/google/uuid"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	// Constants for error messages
	errUnexpectedSigningMethod = "método inesperado de assinatura de token"
	errUnexpectedTokenClaims   = "reivindicações de token inesperadas"
)

// JWTManager é responsável por gerar e validar tokens JWT.
type JWTManager struct {
	secretKey       string
	tokenDuration   time.Duration
	refreshDuration time.Duration
}

// UserClaims representa as informações personalizadas contidas no token JWT.
type UserClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"ID"`
}

// RefreshTokenClaims representa as informações personalizadas contidas no refresh token JWT.
type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"ID"`
}

// NewJWTManager cria e retorna um novo JWTManager.
func NewJWTManager(secretKey string, tokenDuration time.Duration, refreshDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenDuration: tokenDuration, refreshDuration: refreshDuration}
}

// Generate cria e retorna um novo token JWT.
func (manager *JWTManager) Generate(UserID uuid.UUID) (string, string, error) {
	token, err := manager.generateToken(UserID, manager.tokenDuration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := manager.generateToken(UserID, manager.refreshDuration)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (manager *JWTManager) generateToken(UserID uuid.UUID, duration time.Duration) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "server",
		},
		UserID: UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify analisa o token JWT, valida-o e retorna as informações do usuário contidas no token.
func (manager *JWTManager) Verify(tokenStr string) (*UserClaims, error) {
	return manager.verify(tokenStr)
}

// VerifyRefreshToken analisa o refresh token JWT, valida-o e retorna as informações do usuário contidas no token.
func (manager *JWTManager) VerifyRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	return manager.verifyRefreshToken(tokenStr)
}

func (manager *JWTManager) verifyToken(tokenStr string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New(errUnexpectedSigningMethod)
			}
			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func (manager *JWTManager) verifyRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	claims, err := manager.verifyToken(tokenStr, &RefreshTokenClaims{})
	if err != nil {
		return nil, err
	}

	userClaims, ok := claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New(errUnexpectedTokenClaims)
	}

	return userClaims, nil
}

func (manager *JWTManager) verify(tokenStr string) (*UserClaims, error) {
	claims, err := manager.verifyToken(tokenStr, &UserClaims{})
	if err != nil {
		return nil, err
	}

	userClaims, ok := claims.(*UserClaims)
	if !ok {
		return nil, errors.New(errUnexpectedTokenClaims)
	}

	return userClaims, nil
}
