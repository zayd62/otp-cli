// Contains functions related to parsing and generating OTP codes

package lib

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/xlzd/gotp"
)

var Apple = "from otp.go"

// parse an otpauth URI
func ParseOTPAuthURI(otpauthURI string) (OTPStruct, error) {
	u, err := url.Parse(otpauthURI)
	if err != nil {
		fmt.Println(err)
		return OTPStruct{}, err
	}

	uriList := strings.Split(strings.TrimLeft(u.Path, "/"), ":")

	var parsedLabel Label

	if len(uriList) == 2 {
		parsedLabel = Label{
			strings.Split(strings.TrimLeft(u.Path, "/"), ":")[0],
			strings.Split(strings.TrimLeft(u.Path, "/"), ":")[1],
		}
	} else {
		parsedLabel = Label{
			strings.Split(strings.TrimLeft(u.Path, "/"), ":")[0],
			"",
		}
	}

	otpStruct := OTPStruct{
		u.Host,
		parsedLabel,
		u.Query().Get("secret"),
		u.Query().Get("issuer"),
	}
	return otpStruct, nil
}

func GenerateOTPCode(otpStruct OTPStruct) string {
	return gotp.NewDefaultTOTP(otpStruct.secret).Now()
}
