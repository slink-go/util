package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

// https://dev.to/techschoolguru/how-to-create-and-verify-jwt-paseto-token-in-golang-1l5j

// region - Jwt API

var ErrExpiredToken = errors.New("auth token expired")
var ErrInvalidToken = errors.New("auth token invalid")

const minKeySize = 32

// claims fields

type Claims interface {
	Get(key string) string
	GetId() uuid.UUID
	GetIssuer() string
	GetTenant() string
	Expired() bool
}

type Jwt interface {
	Generate(issuer, tenant string, duration time.Duration, fields ...string) (string, error)
	Validate(token string) (Claims, error)
	GetClaims(token string) (Claims, error)
}

func Init(secret string) (Jwt, error) {
	if len(secret) < minKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minKeySize)
	}
	return &jwtImpl{
		secret: []byte(secret),
	}, nil
}

// endregion
// region - token payload

type tokenPayload struct {
	ID        uuid.UUID         `json:"id"`
	Issuer    string            `json:"issuer"`
	Tenant    string            `json:"tenant"`
	IssuedAt  time.Time         `json:"issued_at"`
	ExpiredAt time.Time         `json:"expired_at"`
	Fields    map[string]string `json:"fields"`
}

func (p *tokenPayload) Valid() error {
	if p.Expired() {
		return ErrExpiredToken
	}
	return nil
}
func (p *tokenPayload) Get(key string) string {
	return p.Fields[key]
}
func (p *tokenPayload) GetId() uuid.UUID {
	return p.ID
}
func (p *tokenPayload) GetIssuer() string {
	return p.Issuer
}
func (p *tokenPayload) GetTenant() string {
	return p.Tenant
}
func (p *tokenPayload) Expired() bool {
	return p.ExpiredAt.Before(time.Now())
}

// endregion
// region - token

type jwtImpl struct {
	secret []byte
}

func (j *jwtImpl) Generate(issuer, tenant string, duration time.Duration, fields ...string) (string, error) {
	payload, err := j.newPayload(issuer, tenant, duration)

	if fields != nil && len(fields) > 0 {
		if len(fields)%2 != 0 {
			return "", errors.New("fields must be multiple of 2")
		}
		payload.Fields = make(map[string]string)
		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]
			payload.Fields[key] = value
		}
	}
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(j.secret)
}
func (j *jwtImpl) Validate(token string) (Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return j.secret, nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &tokenPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(Claims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
func (j *jwtImpl) GetClaims(token string) (Claims, error) {
	p := jwt.Parser{}
	cl := tokenPayload{}
	_, _, err := p.ParseUnverified(token, &cl)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}

func (j *jwtImpl) newPayload(issuer, tenant string, duration time.Duration) (*tokenPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &tokenPayload{
		ID:        tokenID,
		Issuer:    issuer,
		Tenant:    tenant,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// endregion
