# Branding

MuxCore household branding package.

Shared product name, marks, and theme tokens for household UI and native clients. Import the Go package for constants and embedded assets, or consume `tokens/theme.json` directly from web builds.

Tracked by https://github.com/Muxcore-Media/umbrella/issues/23

---

## What it provides

| Export | Description |
|--------|-------------|
| `branding.DefaultBrand()` | Product name, short name, tagline |
| `branding.DefaultTheme()` | Color and typography tokens + CSS variable map |
| `branding.DefaultAssets()` | Embedded logo SVGs (`logo.svg`, `logo-mark.svg`) |
| `tokens/theme.json` | Canonical JSON theme for non-Go consumers |

This is a bootstrap package with constants and an assets placeholder contract — not a full design system.

---

## Go usage

```go
import "github.com/Muxcore-Media/branding/pkg/branding"

func main() {
    b := branding.DefaultBrand()
    t := branding.DefaultTheme()

    _ = b.ProductName // "MuxCore"
    _ = t.Colors.Primary // "#3B82F6"
    _ = t.CSSVariables()["--muxcore-color-primary"]

    logo, err := branding.DefaultAssets().Read(branding.AssetLogo)
    if err != nil {
        panic(err)
    }
    _ = logo
}
```

### Umbrella workspace

When checked out as an umbrella submodule:

```bash
cd branding
go test ./...
```

---

## Assets contract

See [assets/README.md](assets/README.md) for file layout and naming rules.

| Asset | Path |
|-------|------|
| Wordmark | `assets/logo/logo.svg` |
| Mark | `assets/logo/logo-mark.svg` |

---

## Theme tokens

`tokens/theme.json` mirrors the Go `Theme` struct. Web clients can import JSON directly; Go clients should use `branding.DefaultTheme()` to avoid drift.

---

## Capabilities

| Capability | Role |
|------------|------|
| `branding.assets` | Shared logo/mark files |
| `branding.theme` | Shared color and typography tokens |

---

## Tests

```bash
go test ./...
```

## License

GPL-3.0
