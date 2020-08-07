package helpers

func StringArrayToByteArray(arguments []string) [][]byte {
	var args [][]byte
	for _, str := range arguments {
		arg := []byte(str)
		args = append(args, arg)
	}

	return args
}

func RemoveItemFromSlice(source []string, deleteValue string) []string {
	for i, value := range source {
		if value == deleteValue {
			return append(source[:i], source[i+1:]...)
		}
	}

	return source
}

func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
