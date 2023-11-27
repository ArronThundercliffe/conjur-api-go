package authn

type iamAuthenticator struct {
	Login              string
	AwsStsSessionToken string
	AwsIamRole         string
	AwsAccount         string
	AuthnIamServiceId  string
	AwsRegion          string
	Authenticate       func(login, awsstssessiontoken, awsiamrole, awsaccount, authniamserviceid, awsregion string) ([]byte, error)
}

func (a *iamAuthenticator) RefreshToken() ([]byte, error) {
	return a.Authenticate(a.Login, a.AwsStsSessionToken, a.AwsIamRole, a.AwsAccount, a.AuthnIamServiceId, a.AwsRegion)
}

func (a *iamAuthenticator) NeedsTokenRefresh() bool {
	return false
}
