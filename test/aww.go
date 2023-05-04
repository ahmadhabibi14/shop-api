type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

claims := &JWTClaim{
	Username: username,
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	},
}