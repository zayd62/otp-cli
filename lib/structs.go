package lib

type Label struct {
	// the issuer
	issuer string
	// the account
	account string
}

// a struct representing an OTP based on https://github.com/google/google-authenticator/wiki/Key-Uri-Format
// - a dashed list
type OTPStruct struct {
	// the otp type
	otpType string
	// the label
	label Label
	// the secret
	secret string
	// the issuer
	issuer string
}
