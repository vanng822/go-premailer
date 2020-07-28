package premailer

// Options for controlling behaviour
type Options struct {
	// Remove class attribute from element
	// Default false
	RemoveClasses bool
	// Copy related CSS properties into HTML attributes (e.g. background-color to bgcolor)
	// Default true
	CssToAttributes bool

	// If true, then style declarations that have "!important" will keep the "!important" in the final
	// style attribute
	// Example:
	//		<style>p { width: 100% !important }</style><p>Text</p>
	// gives
	//		<p style="width: 100% !important">Text</p>
	KeepBangImportant bool
}

// NewOptions return an Options instance with default value
func NewOptions() *Options {
	options := &Options{}
	options.CssToAttributes = true
	options.KeepBangImportant = false
	return options
}
