package util

func TrimByteArray(data []byte, length int) []byte {

	if len(data) <= length {
		return data
	} else {
		result := make([]byte, length)
		copy(result, data)
		return result
	}
}
