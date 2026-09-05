# MuxCore branding assets

Canonical logo and mark files for household UI and native clients.

## Layout

| Path | Purpose |
|------|---------|
| `logo/logo.svg` | Full wordmark (horizontal) |
| `logo/logo-mark.svg` | Compact mark / app icon source |

Future slots (not shipped in v0.1):

| Path | Purpose |
|------|---------|
| `logo/favicon.ico` | Browser favicon |
| `marks/app-icon-*.png` | Platform app icons (iOS, Android, tvOS) |

## Naming rules

- Use lowercase kebab-case file names.
- Prefer SVG for vector marks; raster exports live under `marks/` when added.
- Do not rename shipped paths without a semver bump in `muxcore.json`.

## Go consumers

```go
import "github.com/Muxcore-Media/branding/pkg/branding"

data, err := branding.DefaultAssets().Read(branding.AssetLogo)
```

Assets are embedded via the `assets` package (`assets/embed.go`).

## Non-Go consumers

Copy from this directory or fetch tagged releases. Theme tokens live in `tokens/theme.json`.
