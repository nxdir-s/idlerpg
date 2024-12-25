package components

type Card struct {
	Heading string
	Text    string
	Img     string
}

type StatsCard struct {
	Stat          string
	Value         string
	PercentChange string
	Img           string
}

type BlogCard struct {
	Heading string
	Text    string
	Img     string
	Link    string
}

type LogInCard struct {
	OAuth      []OAuthProvider
	SignInLink string
}

type RegistrationCard struct {
	OAuth      []OAuthProvider
	SignInLink string
}

type OAuthProvider struct {
	Name string
	Icon string
	Link string
}
