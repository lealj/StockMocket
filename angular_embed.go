//go:build !debug

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

// CURRENTLY NOT WORKING/TESTED; purpose is for this to (by itself or by being used by main) to create single .exe that runs
// both the FE and BE

var assets embed.FS

func GetFrontend() http.FileSystem {
	if subAsset, err := fs.Sub(assets, "dist"); err == nil {
		return http.FS(subAsset)
	}

	panic("Failed to load assets")
}

var _angularHandler = http.FileServer(GetFrontend())

// AngularHandler loads angular assets
// This version loads angular from files/assets
// Adds indefinite cache to all files
var AngularHandler = http.HandlerFunc(
	func(writer http.ResponseWriter, rout *http.Request) {
		writer.Header().Add("Cache-Control", "public, max-age=604800")
		_angularHandler.ServeHTTP(writer, rout)
	},
)
