package branding

// Product identity constants for MuxCore household UI and clients.
const (
	ProductName = "MuxCore"
	ShortName   = "MuxCore"
	Tagline     = "Your household media hub."
)

// Brand holds canonical product identity strings.
type Brand struct {
	ProductName string
	ShortName   string
	Tagline     string
}

// DefaultBrand returns the canonical MuxCore brand identity.
func DefaultBrand() Brand {
	return Brand{
		ProductName: ProductName,
		ShortName:   ShortName,
		Tagline:     Tagline,
	}
}
