package apperrors

import "errors"

var (
	ErrRecordNotFoundInDB = errors.New("record_not_found_in_database")
	ErrDB                 = errors.New("db_error")

	ErrInvalidNumberSatellites = errors.New("invalid_number_of_satellites")
)
