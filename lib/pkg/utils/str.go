package utils

// GetStringIndexInSlice returns index of string in slice of strings if not present returns -1
func GetStringIndexInSlice(strToFind string, sliceOfStrings []string) int {
	for strInd, str := range sliceOfStrings {
		if str == strToFind {
			return strInd
		}
	}
	return -1
}

// IsStringInSlice returns if a string is present in the array of strings
func IsStringInSlice(strToFind string, sliceOfStrings []string) bool {
	return GetStringIndexInSlice(strToFind, sliceOfStrings) != -1
}
