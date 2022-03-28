package errors_code

const defaultErr = 99

var listErrors = map[int]string{
	99:  "unknown error",
	100: "empty request",
	101: "error in last index lookup; pos = -1",
	102: "command line error",
	103: "error in parsing request body - payload",
}

func GetErr(num int, str ...string) (int, string) {
	err, ok := listErrors[num]
	if !ok {
		return defaultErr, listErrors[defaultErr]
	}
	if len(str) > 0 {
		err = err + " : " + str[0]
	}
	return num, err
}
