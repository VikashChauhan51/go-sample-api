package routes

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
)

func ScanRoutes(db interfaces.Database) []Route {
	var allRoutes []Route
	booksRoutes := GetBookRoutes(db)
	allRoutes = append(allRoutes, booksRoutes...)
	return allRoutes
}
