package premailer

import "github.com/jaytaylor/html2text"

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

	// If false then text nodes will be escaped for those characters &'<>\"\r
	// If true, then no escaping will be done for text nodes
	// This could be open for XSS attacks if the content is not sanitized
	UnescapedTextNode bool
	// Options for html2text conversion
	// Default is &html2text.Options{PrettyTables: true}
	Html2TextOptions *html2text.Options
}

// NewOptions return an Options instance with default value
func NewOptions() *Options {
	options := &Options{}
	options.CssToAttributes = true
	options.KeepBangImportant = false
	options.UnescapedTextNode = false
	options.Html2TextOptions = &html2text.Options{PrettyTables: true}
	return options
}
