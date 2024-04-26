package isstr

import (
	"github.com/asaskevich/govalidator"
	"github.com/infastin/go-validation"
)

var (
	ErrHexColor       = validation.NewRuleError("is_hex_color", "must be a valid hexadecimal color code")
	ErrRGBColor       = validation.NewRuleError("is_rgb_color", "must be a valid RGB color code")
	ErrLowerCase      = validation.NewRuleError("is_lower_case", "must be in lower case")
	ErrUpperCase      = validation.NewRuleError("is_upper_case", "must be in upper case")
	ErrAlpha          = validation.NewRuleError("is_alpha", "must contain English letters only")
	ErrNumeric        = validation.NewRuleError("is_numeric", "must contain digits only")
	ErrAlphanumeric   = validation.NewRuleError("is_alphanumeric", "must contain English letters and digits only")
	ErrASCII          = validation.NewRuleError("is_ascii", "must contain ASCII characters only")
	ErrPrintableASCII = validation.NewRuleError("is_printable_ascii", "must contain printable ASCII characters only")
	ErrEmail          = validation.NewRuleError("is_email", "must be a valid email address")
	ErrURL            = validation.NewRuleError("is_url", "must be a valid URL")
	ErrUUIDv3         = validation.NewRuleError("is_uuid_v3", "must be a valid UUID v3")
	ErrUUIDv4         = validation.NewRuleError("is_uuid_v4", "must be a valid UUID v4")
	ErrUUIDv5         = validation.NewRuleError("is_uuid_v5", "must be a valid UUID v5")
	ErrUUID           = validation.NewRuleError("is_uuid", "must be a valid UUID")
	ErrULID           = validation.NewRuleError("is_uuid", "must be a valid ULID")
	ErrJSON           = validation.NewRuleError("is_json", "must be in valid JSON format")
	ErrIP             = validation.NewRuleError("is_ip", "must be a valid IP address")
	ErrIPv4           = validation.NewRuleError("is_ipv4", "must be a valid IPv4 address")
	ErrIPv6           = validation.NewRuleError("is_ipv6", "must be a valid IPv6 address")
	ErrSubdomain      = validation.NewRuleError("is_sub_domain", "must be a valid subdomain")
	ErrDomain         = validation.NewRuleError("is_domain", "must be a valid domain")
	ErrDNSName        = validation.NewRuleError("is_dns_name", "must be a valid DNS name")
	ErrHost           = validation.NewRuleError("is_host", "must be a valid IP address or DNS name")
	ErrPort           = validation.NewRuleError("is_port", "must be a valid port number")
	ErrLatitude       = validation.NewRuleError("is_latitude", "must be a valid latitude")
	ErrLongitude      = validation.NewRuleError("is_longitude", "must be a valid longitude")
	ErrSSN            = validation.NewRuleError("is_ssn", "must be a valid social security number")
	ErrSemver         = validation.NewRuleError("is_semver", "must be a valid semantic version")
	ErrBase64         = validation.NewRuleError("is_base64", "must be encoded in Base64")
	ErrDataURI        = validation.NewRuleError("is_data_uri", "must be a Base64-encoded data URI")
	ErrDialString     = validation.NewRuleError("is_dial_string", "must be a valid dial string")
	ErrRequestURL     = validation.NewRuleError("is_request_url", "must be a valid request URL")
	ErrRequestURI     = validation.NewRuleError("request_is_request_uri", "must be a valid request URI")
	ErrCreditCard     = validation.NewRuleError("is_credit_card", "must be a valid credit card number")
)

func HexColor[T ~string](v T) error {
	if !govalidator.IsHexcolor(string(v)) {
		return ErrHexColor
	}
	return nil
}

func RGBColor[T ~string](v T) error {
	if !govalidator.IsRGBcolor(string(v)) {
		return ErrRGBColor
	}
	return nil
}

func LowerCase[T ~string](v T) error {
	if !govalidator.IsLowerCase(string(v)) {
		return ErrLowerCase
	}
	return nil
}

func Alpha[T ~string](v T) error {
	if !govalidator.IsAlpha(string(v)) {
		return ErrAlpha
	}
	return nil
}

func Numeric[T ~string](v T) error {
	if !govalidator.IsNumeric(string(v)) {
		return ErrNumeric
	}
	return nil
}

func Alphanumeric[T ~string](v T) error {
	if !govalidator.IsAlphanumeric(string(v)) {
		return ErrAlphanumeric
	}
	return nil
}

func ASCII[T ~string](v T) error {
	if !govalidator.IsASCII(string(v)) {
		return ErrASCII
	}
	return nil
}

func PrintableASCII[T ~string](v T) error {
	if !govalidator.IsPrintableASCII(string(v)) {
		return ErrPrintableASCII
	}
	return nil
}

func UpperCase[T ~string](v T) error {
	if !govalidator.IsUpperCase(string(v)) {
		return ErrUpperCase
	}
	return nil
}

func Email[T ~string](v T) error {
	if !govalidator.IsEmail(string(v)) {
		return ErrEmail
	}
	return nil
}

func URL[T ~string](v T) error {
	if !govalidator.IsURL(string(v)) {
		return ErrURL
	}
	return nil
}

func UUID[T ~string](v T) error {
	if !govalidator.IsUUID(string(v)) {
		return ErrUUID
	}
	return nil
}

func UUIDv3[T ~string](v T) error {
	if !govalidator.IsUUIDv3(string(v)) {
		return ErrUUIDv3
	}
	return nil
}

func UUIDv4[T ~string](v T) error {
	if !govalidator.IsUUIDv4(string(v)) {
		return ErrUUIDv4
	}
	return nil
}

func UUIDv5[T ~string](v T) error {
	if !govalidator.IsUUIDv5(string(v)) {
		return ErrUUIDv5
	}
	return nil
}

func ULID[T ~string](v T) error {
	if !govalidator.IsULID(string(v)) {
		return ErrULID
	}
	return nil
}

func JSON[T ~string](v T) error {
	if !govalidator.IsJSON(string(v)) {
		return ErrJSON
	}
	return nil
}

func IP[T ~string](v T) error {
	if !govalidator.IsIP(string(v)) {
		return ErrIP
	}
	return nil
}

func IPv4[T ~string](v T) error {
	if !govalidator.IsIPv4(string(v)) {
		return ErrIPv4
	}
	return nil
}

func IPv6[T ~string](v T) error {
	if !govalidator.IsIPv6(string(v)) {
		return ErrIPv6
	}
	return nil
}

func DNSName[T ~string](v T) error {
	if !govalidator.IsDNSName(string(v)) {
		return ErrDNSName
	}
	return nil
}

func Host[T ~string](v T) error {
	if !govalidator.IsHost(string(v)) {
		return ErrHost
	}
	return nil
}

func Port[T ~string](v T) error {
	if !govalidator.IsPort(string(v)) {
		return ErrPort
	}
	return nil
}

func Latitude[T ~string](v T) error {
	if !govalidator.IsLatitude(string(v)) {
		return ErrLatitude
	}
	return nil
}

func Longitude[T ~string](v T) error {
	if !govalidator.IsLongitude(string(v)) {
		return ErrLongitude
	}
	return nil
}

func SSN[T ~string](v T) error {
	if !govalidator.IsSSN(string(v)) {
		return ErrSSN
	}
	return nil
}

func Semver[T ~string](v T) error {
	if !govalidator.IsSemver(string(v)) {
		return ErrSemver
	}
	return nil
}

func Base64[T ~string](v T) error {
	if !govalidator.IsBase64(string(v)) {
		return ErrBase64
	}
	return nil
}

func DataURI[T ~string](v T) error {
	if !govalidator.IsDataURI(string(v)) {
		return ErrDataURI
	}
	return nil
}

func DialString[T ~string](v T) error {
	if !govalidator.IsDialString(string(v)) {
		return ErrDialString
	}
	return nil
}

func RequestURL[T ~string](v T) error {
	if !govalidator.IsRequestURL(string(v)) {
		return ErrRequestURL
	}
	return nil
}

func RequestURI[T ~string](v T) error {
	if !govalidator.IsRequestURI(string(v)) {
		return ErrRequestURI
	}
	return nil
}

func CreditCard[T ~string](v T) error {
	if !govalidator.IsCreditCard(string(v)) {
		return ErrCreditCard
	}
	return nil
}
