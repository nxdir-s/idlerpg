package valobj

const (
	AccessToken  string = "AccessToken"
	RefreshToken string = "RefreshToken"
)

type Token struct {
	Access  string
	Refresh string
}
