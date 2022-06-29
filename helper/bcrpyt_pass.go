package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(actualHashedPassword string, inputPlainPassword string) bool {
	byteActualHashedPassword := []byte(actualHashedPassword)
	byteInputPlainPassword := []byte(inputPlainPassword)

	err := bcrypt.CompareHashAndPassword(byteActualHashedPassword, byteInputPlainPassword)
	return err == nil //nil means it is match
}

//untuk cek error jangan langsung err.Error() tanpa cek if err!=nil, karena akan muncul error
//runtime error: invalid memory address or nil pointer dereference
