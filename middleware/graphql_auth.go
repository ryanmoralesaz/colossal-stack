package middleware

import (
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// InjectUserContext extracts JWT and injects user info into request context
func InjectUserContext(c *fiber.Ctx, ctx context.Context) context.Context {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ctx
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return ctx
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return ctx
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return ctx
	}

	// Inject user info into context
	ctx = context.WithValue(ctx, "userID", claims.UserID)
	ctx = context.WithValue(ctx, "email", claims.Email)

	return ctx
}
