package lib

type input struct {
	vcshost     string
	clientid    string
	secretid    string
	redirecturl string
}

func NewInput(vcshost, clientid, secretid, redirecturl string) *input {
	return &input{vcshost, clientid, secretid, redirecturl}
}
