package helpers

func StringArrayToByteArray(arguments []string) [][]byte {
	var args [][]byte
	for _, str := range arguments {
		arg := []byte(str)
		args = append(args, arg)
	}

	return args
}
