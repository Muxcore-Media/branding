package branding

import (
	"strings"
	"testing"
)

func TestDefaultBrand(t *testing.T) {
	b := DefaultBrand()
	if b.ProductName != "MuxCore" {
		t.Errorf("ProductName = %q, want MuxCore", b.ProductName)
	}
	if b.ShortName != "MuxCore" {
		t.Errorf("ShortName = %q, want MuxCore", b.ShortName)
	}
	if b.Tagline == "" {
		t.Error("Tagline must not be empty")
	}
}

func TestDefaultThemeColors(t *testing.T) {
	theme := DefaultTheme()
	colors := []string{
		theme.Colors.Primary,
		theme.Colors.Background,
		theme.Colors.Text,
	}
	for _, c := range colors {
		if !strings.HasPrefix(c, "#") {
			t.Errorf("color %q should be a hex value", c)
		}
	}
}

func TestThemeCSSVariables(t *testing.T) {
	vars := DefaultTheme().CSSVariables()
	if vars["--muxcore-color-primary"] != "#3B82F6" {
		t.Errorf("primary CSS var = %q", vars["--muxcore-color-primary"])
	}
	if len(vars) < 10 {
		t.Errorf("expected at least 10 CSS variables, got %d", len(vars))
	}
}

func TestEmbeddedAssets(t *testing.T) {
	assets := DefaultAssets()
	for _, name := range []string{AssetLogo, AssetLogoMark} {
		data, err := assets.Read(name)
		if err != nil {
			t.Fatalf("Read(%q): %v", name, err)
		}
		if !strings.Contains(string(data), "<svg") {
			t.Errorf("asset %q does not look like SVG", name)
		}
	}
}
