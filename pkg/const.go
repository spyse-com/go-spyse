package spyse

// Options for search operators
const (
	OperatorEqual              = "eq"
	OperatorNotEqual           = "not_eq"
	OperatorContains           = "contains"
	OperatorNotContains        = "not_contains"
	OperatorStartsWith         = "starts"
	OperatorEndsWith           = "ends"
	OperatorExists             = "exists"
	OperatorNotExists          = "not_exists"
	OperatorGreaterThanOrEqual = "gte"
	OperatorLessThanOrEqual    = "lte"
)

// Options for DNS history
const (
	DNSTYPEA     = "A"
	DNSTYPEMX    = "MX"
	DNSTYPENS    = "NS"
	DNSTYPETXT   = "TXT"
	DNSTYPEAAAA  = "AAAA"
	DNSTYPECNAME = "CNAME"
)
