package branding

// Theme holds CSS-friendly color and typography tokens shared across household UI.
type Theme struct {
	Colors     ColorTokens
	Typography TypographyTokens
}

// ColorTokens are hex color values suitable for CSS custom properties or native theming.
type ColorTokens struct {
	Primary       string
	PrimaryHover  string
	Background    string
	Surface       string
	Text          string
	TextMuted     string
	Accent        string
	Border        string
	Success       string
	Warning       string
	Error         string
}

// TypographyTokens hold font family and scale hints for clients.
type TypographyTokens struct {
	FontFamilySans string
	FontFamilyMono string
	FontSizeBase   string
	FontSizeSm     string
	FontSizeLg     string
}

// DefaultTheme returns the canonical MuxCore household theme tokens.
func DefaultTheme() Theme {
	return Theme{
		Colors: ColorTokens{
			Primary:      "#3B82F6",
			PrimaryHover: "#2563EB",
			Background:   "#0F172A",
			Surface:      "#1E293B",
			Text:         "#F8FAFC",
			TextMuted:    "#94A3B8",
			Accent:       "#38BDF8",
			Border:       "#334155",
			Success:      "#22C55E",
			Warning:      "#EAB308",
			Error:        "#EF4444",
		},
		Typography: TypographyTokens{
			FontFamilySans: "system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif",
			FontFamilyMono: "ui-monospace, 'Cascadia Code', 'Source Code Pro', monospace",
			FontSizeBase:   "16px",
			FontSizeSm:     "14px",
			FontSizeLg:     "18px",
		},
	}
}

// CSSVariables returns a map of CSS custom property names to token values.
func (t Theme) CSSVariables() map[string]string {
	return map[string]string{
		"--muxcore-color-primary":        t.Colors.Primary,
		"--muxcore-color-primary-hover":  t.Colors.PrimaryHover,
		"--muxcore-color-background":     t.Colors.Background,
		"--muxcore-color-surface":        t.Colors.Surface,
		"--muxcore-color-text":           t.Colors.Text,
		"--muxcore-color-text-muted":     t.Colors.TextMuted,
		"--muxcore-color-accent":         t.Colors.Accent,
		"--muxcore-color-border":         t.Colors.Border,
		"--muxcore-color-success":        t.Colors.Success,
		"--muxcore-color-warning":        t.Colors.Warning,
		"--muxcore-color-error":          t.Colors.Error,
		"--muxcore-font-family-sans":     t.Typography.FontFamilySans,
		"--muxcore-font-family-mono":     t.Typography.FontFamilyMono,
		"--muxcore-font-size-base":       t.Typography.FontSizeBase,
		"--muxcore-font-size-sm":         t.Typography.FontSizeSm,
		"--muxcore-font-size-lg":         t.Typography.FontSizeLg,
	}
}
