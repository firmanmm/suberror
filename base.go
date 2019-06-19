package suberror

//RuntimeError a global error that is categorized as runtime error
var RuntimeError = newErrorType()

//IOError represent any IO related error
var IOError = RuntimeError.Derive()

//NetworkError represent any network related error
var NetworkError = IOError.Derive()
