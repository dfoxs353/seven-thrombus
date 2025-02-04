package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct {
	tokenSecret string
	accessTTL   time.Duration
	refreshTTL  time.Duration
}

func New(
	tokenSecret string,
	accessTTL time.Duration,
	refreshTTL time.Duration,
) TokenManager {
	return TokenManager{
		tokenSecret: tokenSecret,
		accessTTL:   accessTTL,
		refreshTTL:  refreshTTL,
	}
}

type TokenPair struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn int    `json:"-"`
}

func (t *TokenManager) GenerateTokens(uid int, roles []string) (TokenPair, error) {
	generator := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   uid,
		"roles": roles,
		"exp":   time.Now().Add(t.accessTTL).Unix(),
	})
	accessToken, err := generator.SignedString([]byte(t.tokenSecret))
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          uuid.New().String(),
		RefreshTokenExpiresIn: int(time.Now().Add(t.refreshTTL).Unix()),
	}, nil
}

type Payload struct {
	Uid   int
	Roles []string
}

func (tg *TokenManager) ParseToken(token string) (Payload, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method alg: %v", t.Header["alg"])
		}

		return []byte(tg.tokenSecret), nil
	})
	if err != nil {
		return Payload{}, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return Payload{}, errors.New("unexpected jwt claims type")
	}

	if val, ok := claims["exp"].(float64); time.Now().After(time.Unix(int64(val), 0)) && ok {
		return Payload{}, errors.New("token expired")
	}

	var (
		uid   int
		roles []string
	)

	if id, ok := claims["uid"].(float64); ok && id > 0 {
		uid = int(id)
	}

	if rol, ok := claims["roles"].([]any); ok {
		for _, r := range rol {
			roles = append(roles, r.(string))
		}
	}

	if uid < 1 || len(roles) < 1 {
		return Payload{}, errors.New("invalid token claims")
	}

	return Payload{
		Uid:   uid,
		Roles: roles,
	}, nil
}
