package constants

type URI string

const (
	URL_MAIN URI = "/RupeeGo/TCP"
)

const (
	GetVersion URI = URL_MAIN + "/GetVersion"
)

// String method to get a string representation of the Status
func (s URI) String() string {
	return string(s)
}
