package models

import "errors"

var (
	ERR_RECORD_NOT_FOUND = errors.New("Record not found")
	ERR_RECORD_DB        = errors.New("Internal server error")
	ERR_DATE_PARSING     = errors.New("Error while parsing date")
)
