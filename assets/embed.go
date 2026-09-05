// Package assets embeds canonical MuxCore brand marks for offline consumers.
package assets

import "embed"

//go:embed logo/*
var FS embed.FS

// Logo is the full wordmark SVG path relative to the embedded FS root.
const Logo = "logo/logo.svg"

// LogoMark is the compact mark SVG path relative to the embedded FS root.
const LogoMark = "logo/logo-mark.svg"
