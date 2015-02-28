package premailer

import ()

// NewOptions return an Options instance with default value
func NewOptions() *Options {
	options := &Options{}
	options.CssToAttributes = true
	return options
}
