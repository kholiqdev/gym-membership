package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gym/config"
	"gym/internal/protocol/http/response"
	"time"
)

func JwtVerifyAccess(role string) echo.MiddlewareFunc {
	if role == "" {
		role = "member"
	}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Get().JwtSecretKey),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				claims := t.Claims.(jwt.MapClaims)
				rawExp := claims["exp"].(float64)
				exp := int64(rawExp)
				if t.Method.Alg() != "HS256" {
					log.Err(errors.New("unexpected signing method")).Msgf("unexpected signing method: %v", t.Method.Alg())
					return nil, errors.New("unexpected jwt signing method")
				} else if claims["token_type"] != "access_token" {
					log.Err(errors.New("unexpected token type")).Msgf("unexpected token type: %v", claims["token_type"])
					response.Err(c, 401, fmt.Errorf("unexpected token type %v", claims["token_type"]))
					return nil, fmt.Errorf("unexpected token type %v", claims["token_type"])
				} else if exp < time.Now().Unix() {
					log.Err(errors.New("token expired")).Msgf("token expired: %v", exp)
					response.Err(c, 401, errors.New("token expired"))
					return nil, errors.New("token expired")
				}

				return []byte(config.Get().JwtSecretKey), nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				log.Err(errors.New("invalid token")).Msgf("invalid token: %v", token.Raw)
				return nil, errors.New("invalid token")
			}
			if token.Claims.(jwt.MapClaims)["role"] != role {
				log.Err(errors.New("invalid token role")).Msgf("invalid token role: %v", token.Raw)
				return nil, errors.New("invalid token role")
			}
			//Send User Auth To Context
			c.Set("user_id", token.Claims.(jwt.MapClaims)["id"])
			return token, nil
		},
	})
}

func JwtVerifyRefresh(role string) echo.MiddlewareFunc {
	if role == "" {
		role = "member"
	}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Get().JwtSecretKey),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				claims := t.Claims.(jwt.MapClaims)
				rawExp := claims["exp"].(float64)
				exp := int64(rawExp)

				if t.Method.Alg() != "HS256" {
					log.Err(errors.New("unexpected signing method")).Msgf("unexpected signing method: %v", t.Method.Alg())
					return nil, errors.New("unexpected jwt signing method")
				} else if claims["token_type"] != "refresh_token" {
					log.Err(errors.New("unexpected token type")).Msgf("unexpected token type: %v", claims["token_type"])
					response.Err(c, 401, fmt.Errorf("unexpected token type %v", claims["token_type"]))
					return nil, fmt.Errorf("unexpected token type %v", claims["token_type"])
				} else if exp < time.Now().Unix() {
					log.Err(errors.New("token expired")).Msgf("token expired: %v", exp)
					response.Err(c, 401, errors.New("token expired"))
					return nil, errors.New("token expired")
				}

				return []byte(config.Get().JwtSecretKey), nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				log.Err(errors.New("invalid token")).Msgf("invalid token: %v", token.Raw)
				return nil, errors.New("invalid token")
			}
			if token.Claims.(jwt.MapClaims)["role"] != role {
				log.Err(errors.New("invalid token role")).Msgf("invalid token role: %v", token.Raw)
				return nil, errors.New("invalid token role")
			}
			//Send User ID To Context
			c.Set("user_id", token.Claims.(jwt.MapClaims)["id"])

			return token, nil
		},
	})
}
