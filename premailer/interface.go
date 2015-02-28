package premailer

import (

)
// Inteface of Premailer
type Premailer interface {
	// Transform process and inlining css
	// It start to collect the rules in the document style tags
	// Calculate specificity and sort the rules based on that
	// It then collects the affected elements
	// And applies the rules on those
	// The leftover rules will put back into a style element
	Transform() (string, error)
}

// Options for controlling behaviour
type Options struct {
	// Remove class attribute from element
	// Default false
	removeClasses bool
	// Copy related CSS properties into HTML attributes (e.g. background-color to bgcolor)
	// Default true
	cssToAttributes bool
}