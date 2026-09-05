package branding

import (
	"io/fs"

	"github.com/Muxcore-Media/branding/assets"
)

// Asset path constants (see assets/README.md for the full naming contract).
const (
	AssetLogo     = assets.Logo
	AssetLogoMark = assets.LogoMark
)

// Assets exposes embedded branding files for offline consumers.
type Assets struct {
	fs fs.FS
}

// DefaultAssets returns an Assets view over the embedded asset tree.
func DefaultAssets() *Assets {
	return &Assets{fs: assets.FS}
}

// Read returns the bytes for a named asset (e.g. AssetLogo).
func (a *Assets) Read(name string) ([]byte, error) {
	return fs.ReadFile(a.fs, name)
}

// Open opens a named asset for streaming reads.
func (a *Assets) Open(name string) (fs.File, error) {
	return a.fs.Open(name)
}
