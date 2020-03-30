package dynamicpgxpool

type CredentialsProvider interface {
	Credentials() (username string, password string, err error)
}
