package utils

func GenerateResetToken(userID uint) (string, error) {
	return GenerateJWT(userID)
}
