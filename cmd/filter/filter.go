package filter

type Filter struct {
	Labels []string `help:"Include all devices with these labels"`
	URLS   []string `help:"Include all devices with these URLs"`
}
