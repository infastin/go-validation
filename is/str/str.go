package isstr

import (
	"os"
	"regexp"

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
	ErrUUIDv7         = validation.NewRuleError("is_uuid_v5", "must be a valid UUID v7")
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
	ErrISBN10         = validation.NewRuleError("is_isbn_10", "must be a valid ISBN-10")
	ErrISBN13         = validation.NewRuleError("is_isbn_13", "must be a valid ISBN-13")
	ErrISBN           = validation.NewRuleError("is_isbn", "must be a valid ISBN")
	ErrMongoID        = validation.NewRuleError("is_mongo_id", "must be a valid hex-encoded MongoDB ObjectId")
	ErrCurrencyCode   = validation.NewRuleError("is_currency_code", "must be a valid currency code")
	ErrCountryCode2L  = validation.NewRuleError("is_two_letter_country_code", "must be a valid two-letter country code")
	ErrCountryCode3L  = validation.NewRuleError("is_three_letter_country_code", "must be a valid three-letter country code")
	ErrLanguageCode2L = validation.NewRuleError("is_two_letter_language_code", "must be a valid two-letter language code")
	ErrLanguageCode3L = validation.NewRuleError("is_three_letter_language_code", "must be a valid three-letter language code")
	ErrPath           = validation.NewRuleError("is_path", "must be a valid path")
	ErrFile           = validation.NewRuleError("is_file", "must be a valid path to a file")
	ErrDirectory      = validation.NewRuleError("is_directory", "must be a valid path to a directory")
)

var (
	rxUUID7 = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
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

func ExistingEmail[T ~string](v T) error {
	if !govalidator.IsExistingEmail(string(v)) {
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

func UUIDv7[T ~string](v T) error {
	if !rxUUID7.MatchString(string(v)) {
		return ErrUUIDv7
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

func ISBN10[T ~string](v T) error {
	if !govalidator.IsISBN10(string(v)) {
		return ErrISBN10
	}
	return nil
}

func ISBN13[T ~string](v T) error {
	if !govalidator.IsISBN13(string(v)) {
		return ErrISBN13
	}
	return nil
}

func ISBN[T ~string](v T) error {
	if !govalidator.IsISBN10(string(v)) && !govalidator.IsISBN13(string(v)) {
		return ErrISBN
	}
	return nil
}

func MongoID[T ~string](v T) error {
	if !govalidator.IsMongoID(string(v)) {
		return ErrMongoID
	}
	return nil
}

func CurrencyCode[T ~string](v T) error {
	if !govalidator.IsISO4217(string(v)) {
		return ErrCurrencyCode
	}
	return nil
}

func CountryCode2L[T ~string](v T) error {
	if !govalidator.IsISO3166Alpha2(string(v)) {
		return ErrCountryCode2L
	}
	return nil
}

func CountryCode3L[T ~string](v T) error {
	if !govalidator.IsISO3166Alpha3(string(v)) {
		return ErrCountryCode3L
	}
	return nil
}

func LanguageCode2L[T ~string](v T) error {
	if !govalidator.IsISO693Alpha2(string(v)) {
		return ErrLanguageCode2L
	}
	return nil
}

func LanguageCode3L[T ~string](v T) error {
	if !govalidator.IsISO693Alpha3b(string(v)) {
		return ErrLanguageCode3L
	}
	return nil
}

func Path[T ~string](v T) error {
	if _, err := os.Stat(string(v)); err != nil {
		return ErrPath
	}
	return nil
}

func File[T ~string](v T) error {
	if stat, err := os.Stat(string(v)); err != nil || !stat.Mode().IsRegular() {
		return ErrFile
	}
	return nil
}

func Directory[T ~string](v T) error {
	if stat, err := os.Stat(string(v)); err != nil || !stat.Mode().IsDir() {
		return ErrDirectory
	}
	return nil
}
