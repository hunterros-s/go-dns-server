package domain

type QueryType uint16

const (
	UNKNOWN QueryType = 0
	A       QueryType = 1
	NS      QueryType = 2
	CNAME   QueryType = 5
	SOA     QueryType = 6
	PTR     QueryType = 12
	MX      QueryType = 15
	TXT     QueryType = 16
	AAAA    QueryType = 28
	SRV     QueryType = 33
)
