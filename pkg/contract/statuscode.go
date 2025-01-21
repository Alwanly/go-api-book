package contract

const (
	// Common error messages
	ErrorValidatePayload       = "Failed to validate payload"
	ErrorMutatePayload         = "Failed to mutate payload"
	ErrorInsufficientPrivilege = "User does not have privilege to perform this action"

	// Common error message database
	ErrorFailedToFindRecord   = "Failed to find record"
	ErrorFailedToReadCursor   = "Failed to read cursor"
	ErrorFailedToCountRecord  = "Failed to count record"
	ErrorFailedToDeleteRecord = "Failed to delete record"
	ErrorFailedToInsertRecord = "Failed to insert record"
	ErrorFailedToUpdateRecord = "Failed to update record"
)

type StatusCode string

const (
	StatusCodeSuccess               = StatusCode("000000")
	StatusCodeBindingFailed         = StatusCode("000001")
	StatusCodeValidationFailed      = StatusCode("000002")
	StatusCodeUnauthorized          = StatusCode("000011")
	StatusCodeUserOrPasswordInvalid = StatusCode("000012")
	StatusCodeInternalServerError   = StatusCode("000013")
	StatusCodeSequenceError         = StatusCode("000014")
)

func CreateStatusCode(code string) StatusCode {
	return StatusCode(code)
}
