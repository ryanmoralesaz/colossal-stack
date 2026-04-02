package graph

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

// GetUserIDFromContext extracts user ID from context
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value("userID").(uint)
	if !ok {
		return 0, fmt.Errorf("unauthorized: authentication required for this operation")
	}
	return userID, nil
}
