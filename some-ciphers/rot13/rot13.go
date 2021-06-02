package rot13

import "algorithm-golang/some-ciphers/caesar"

func rot13(input string) string {
	return caesar.Encrypt(input, 13)
}
