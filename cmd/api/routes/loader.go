package routes

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
)

// ScanRoutes scans all route files in the current directory and collects routes.
func ScanRoutes(db interfaces.Database) ([]Route, error) {
	var allRoutes []Route

	// Find all *_routes.go files in the current directory
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), "_routes.go") {
			routes, err := loadRoutesFromPlugin(path, db)
			if err != nil {
				return err
			}
			allRoutes = append(allRoutes, routes...)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan route files: %v", err)
	}

	return allRoutes, nil
}

// loadRoutesFromPlugin loads routes from a plugin file.
func loadRoutesFromPlugin(path string, db interfaces.Database) ([]Route, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open plugin %s: %v", path, err)
	}

	sym, err := p.Lookup("Routes")
	if err != nil {
		return nil, fmt.Errorf("failed to find GetRoutes in plugin %s: %v", path, err)
	}

	getRoutesFunc, ok := sym.(func(interfaces.Database) []Route)
	if !ok {
		return nil, fmt.Errorf("GetRoutes function has incorrect signature in plugin %s", path)
	}

	return getRoutesFunc(db), nil
}
