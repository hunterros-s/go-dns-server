package domain

type ResponseCode uint8

const (
	NOERROR ResponseCode = iota
	FORMERR
	SERVAIL
	NXDOMAIN
	NOTIMP
	REFUSED
)
