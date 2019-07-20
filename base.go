package suberror

//RuntimeError a global error that is categorized as runtime error
var RuntimeError = newBaseErrorType()

//ClientError represent any error due to client input
var ClientError = RuntimeError.Derive()

//InternalError represent any error that happen because of server problem
var InternalError = RuntimeError.Derive()

//IOError represent any IO related error
var IOError = InternalError.Derive()

//NetworkError represent any network related error
var NetworkError = IOError.Derive()
