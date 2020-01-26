package dynamicpgxpool

type CredentialsProvider interface{
	GetCredentials() (username string, password string, err error)
}