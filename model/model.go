package model

import (

)

type Counter struct {
	UUID string `json:uuid`
	Count int `json:count`
}