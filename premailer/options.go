package premailer

import "github.com/inbucket/html2text"

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

type PremailerOption func(*Options)

func WithRemoveClasses(remove bool) PremailerOption {
	return func(o *Options) {
		o.RemoveClasses = remove
	}
}

func WithCssToAttributes(cssToAttr bool) PremailerOption {
	return func(o *Options) {
		o.CssToAttributes = cssToAttr
	}
}

func WithKeepBangImportant(keep bool) PremailerOption {
	return func(o *Options) {
		o.KeepBangImportant = keep
	}
}

// NewOptions return an Options instance with default value
func NewOptions(o ...PremailerOption) *Options {
	options := &Options{
		CssToAttributes: true,
	}

	for _, opt := range o {
		opt(options)
	}
	return options
}

type Html2TextOption func(*Html2TextOptions)

type Html2TextOptions struct {
	PrettyTables        bool                           // Turns on pretty ASCII rendering for table elements.
	OmitLinks           bool                           // Turns on omitting links
	TextOnly            bool                           // Returns only plain text
	PrettyTablesOptions *html2text.PrettyTablesOptions // Configures pretty ASCII rendering for table elements.
}

func (o *Html2TextOptions) toHtml2TextOptions() html2text.Options {
	return html2text.Options{PrettyTables: o.PrettyTables, OmitLinks: o.OmitLinks, TextOnly: o.TextOnly, PrettyTablesOptions: o.PrettyTablesOptions}
}

func WithPrettyTables(pretty bool) Html2TextOption {
	return func(o *Html2TextOptions) {
		o.PrettyTables = pretty
	}
}

func WithOmitLinks(omit bool) Html2TextOption {
	return func(o *Html2TextOptions) {
		o.OmitLinks = omit
	}
}

func WithTextOnly(textOnly bool) Html2TextOption {
	return func(o *Html2TextOptions) {
		o.TextOnly = textOnly
	}
}

func WithPrettyTablesOptions(opts *html2text.PrettyTablesOptions) Html2TextOption {
	return func(o *Html2TextOptions) {
		o.PrettyTablesOptions = opts
	}
}

// newHtml2TextOptions return an Html2TextOptions instance with PrettyTables value true
func newHtml2TextOptions(opts ...Html2TextOption) *Html2TextOptions {
	options := Html2TextOptions{
		PrettyTables: true,
	}
	for _, o := range opts {
		o(&options)
	}
	return &options
}
