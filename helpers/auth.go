package helpers

import (
	"os"
	"strconv"
	"time"

	logger "myturbogarage/loggers"

	"github.com/golang-jwt/jwt"
)

var log = logger.NewLogger()

// getAuthSecrete ...
func getAuthSecrete() string {
	return GetEnv("AUTH_SECRET_KEY", "secrete")
}

// Token ...
type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// AuthClaims is the struct that will be encoded to a JWT.
type AuthClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateAuthTokens generates both the detailed token and refresh token
// tokenExpiry is optional and defaults to 3 hours for access token and 7 days for refresh token
func GenerateAuthTokens(AuthClaims *AuthClaims, tokenExpiry ...time.Duration) (*Token, error) {
	accessTokenExpiry := time.Now().Add(time.Hour * 3).Unix()       // 3 hours
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days
	if len(tokenExpiry) > 0 {
		accessTokenExpiry = time.Now().Add(tokenExpiry[0]).Unix()
	}

	if len(tokenExpiry) > 1 {
		refreshTokenExpiry = time.Now().Add(tokenExpiry[1]).Unix()
	}

	AuthClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: accessTokenExpiry,
		Issuer:    "ochom",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims).SignedString([]byte(getAuthSecrete()))
	if err != nil {
		return nil, err
	}

	AuthClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: refreshTokenExpiry,
		Issuer:    "ochom",
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims).SignedString([]byte(getAuthSecrete()))
	if err != nil {
		return nil, err
	}

	return &Token{AccessToken: token, RefreshToken: refreshToken}, nil
}

// ValidateToken validates the token
func ValidateToken(token string) (*AuthClaims, error) {
	AuthClaims := &AuthClaims{}
	tkn, err := jwt.ParseWithClaims(token, AuthClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(getAuthSecrete()), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return AuthClaims, nil
}

// GetEnv ...
func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Warn("Environment variable `%s` not found, returning default value `%s`\n", key, defaultValue)
		return defaultValue
	}

	return value
}

// GetEnvInt ...
func GetEnvInt(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Warn("Environment variable `%s` not found, returning default value `%v`\n", key, defaultValue)
		return defaultValue
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		log.Warn("Environment variable `%s` error: %s, returning default value `%v`\n", key, err.Error(), defaultValue)
		return defaultValue
	}

	return val
}

// GetEnvBool ...
func GetEnvBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Warn("Environment variable `%s` not found, returning default value `%v`\n", key, defaultValue)
		return defaultValue
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		log.Warn("Environment variable `%s` error: %s, returning default value `%v`\n", key, err.Error(), defaultValue)
		return defaultValue
	}

	return val
}

// GetEnvFloat ...
func GetEnvFloat(key string, defaultValue float64) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Warn("Environment variable `%s` not found, returning default value `%v`\n", key, defaultValue)
		return defaultValue
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Warn("Environment variable `%s` error: %s, returning default value `%v`\n", key, err.Error(), defaultValue)
		return defaultValue
	}

	return val
}
