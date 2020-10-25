package suberror

//Error base interface for suberror's error
type Error interface {
	TypeOf(err ErrorType) bool
	Error() string
	GetCode() ErrorCode
	setMessage(message string)
}

//BaseError represent an error message
type BaseError struct {
	message string
	errType ErrorType
}

//TypeOf perform check if this error is part of other ErrorType
func (b *BaseError) TypeOf(err ErrorType) bool {
	return b.errType.TypeOf(err)
}

//Error return representative error message
func (b *BaseError) Error() string {
	return b.message
}

//GetCode return error code of this error
func (b *BaseError) GetCode() ErrorCode {
	return b.errType.GetCode()
}

//String yield string representation
func (b *BaseError) String() string {
	return b.message
}

func (b *BaseError) setMessage(message string) {
	b.message = message
}
