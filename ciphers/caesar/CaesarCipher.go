package caesar

func Encrypt(input string, key int) string {
	key = (key%26 + 26) % 26

	var outputBuffer []byte

	for _, r := range input {
		newByte := byte(r)

		if 'A' <= newByte && newByte <= 'Z' {
			outputBuffer = append(outputBuffer, byte(int('A')+int(int(newByte-'A')+key)%26))
		} else if 'a' <= newByte && newByte <= 'z' {
			outputBuffer = append(outputBuffer, byte(int('a')+int(int(newByte-'a')+key)%26))
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}

	}
	return string(outputBuffer)
}

func Decrypt(input string, key int) string {
	return Encrypt(input, 26-key)
}