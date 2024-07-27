package routes

import (
	"github.com/VikashChauhan51/collections"
	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
)

func ScanRoutes(db interfaces.Database) []Route {
	allRoutes := collections.NewArrayList[Route]()
	booksRoutes := GetBookRoutes(db)
	allRoutes.AddRange(booksRoutes)
	return allRoutes.Items()
}
