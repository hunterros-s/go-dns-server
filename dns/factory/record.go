package factory

import "github.com/hunterros-s/go-dns-server/dns"

type RecordConstructor func(dns.RecordInfo, dns.Buffer) (dns.Record, error)

type RecordFactory struct {
	recordconstructor RecordConstructor
}

func NewRecordFactory(rc RecordConstructor) *RecordFactory {
	return &RecordFactory{
		recordconstructor: rc,
	}
}

func (rf *RecordFactory) New(buffer dns.Buffer) (dns.Record, error) {
	ri, err := NewRecordInfo(buffer)
	if err != nil {
		return nil, err
	}

	r, err := rf.recordconstructor(ri, buffer)
	if err != nil {
		return nil, err
	}

	return r, nil
}
