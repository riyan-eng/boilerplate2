package util

import (
	"boilerplate/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
)

type CustomClaims struct {
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	UID       string `json:"uid"`
	jwt.StandardClaims
}

type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}

func GenerateJWT(issuer string, userID string, companyID string, autoLogoffMinutes int) (string, string, int64, error) {
	accessToken, accessUID, expiredAt, _ := createToken(issuer, 30, userID, companyID, config.GetEnv("JWT_SECRET_ACCESS"))
	refreshToken, refreshUID, _, _ := createToken(issuer, 60, userID, companyID, config.GetEnv("JWT_SECRET_REFRESH"))
	cacheJSON, err := json.Marshal(CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	})

	ctx := context.Background()
	config.Redis.Set(ctx, fmt.Sprintf("token-%s", userID), string(cacheJSON), time.Minute*time.Duration(autoLogoffMinutes))
	return accessToken, refreshToken, expiredAt, err
}

func createToken(issuer string, expireMinutes int, userID string, companyID string, secret string) (t string, uid string, exp int64, err error) {
	exp = time.Now().Add(time.Minute * 5).Unix()
	uid = uuid.New().String()
	claims := CustomClaims{
		userID,
		companyID,
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix(),
			Issuer:    issuer,
		},
	}
	mySigningKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString(mySigningKey)
	return
}

func ParseToken(tokenString string, secret string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	claims = token.Claims.(*CustomClaims)
	return
}

func ValidateToken(claims *CustomClaims, isRefresh bool, ctx *fasthttp.RequestCtx) (err error) {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, _ := config.Redis.Get(ctx, fmt.Sprintf("token-%s", claims.UserID)).Result()
		cachedTokens := new(CachedTokens)
		err := json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string
		if isRefresh {
			tokenUID = cachedTokens.RefreshUID
		} else {
			tokenUID = cachedTokens.AccessUID
		}
		if err != nil || tokenUID != claims.UID {
			return errors.New("token not found")
		}
		return nil
	})

	err = g.Wait()
	return
}
