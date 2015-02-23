package premailer

import (

)

type Premailer interface {
	Transform() (string, error)
}