package messages

const (
	SuccessfulCreateMessage   = "Successfully created %s"
	SuccessfulListMessage     = "Successfully listed %ss"
	SuccessfulRetrieveMessage = "Successfully retrieved %s"
	SuccessfulEditMessage     = "Successfully edited %s"
	SuccessfulDeleteMessage   = "Successfully deleted %s"
	SuccessfulLoginMessage    = "Successfully logged in"

	ErrCreateFailure = "Failed to create %s in %s"
	ErrRetrieveData  = "Failed to retrieve data in %s"
	ErrEditFailure   = "Failed to edit %s in %s"
	ErrDeleteFailure = "Failed to delete %s in %s"
	ErrLoginFailure  = "Failed to login in %s"

	ErrParseForm        = "Failed to parse form in %s"
	ErrParseURLParams   = "Failed to parse URL parameters in %s"
	ErrRetrieveDatabase = "Failed to retrieve database in %s"
	ErrEncodeView       = "Failed to encode data into JSON format in %s"
)
