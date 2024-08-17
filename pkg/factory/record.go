package factory

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type RecordConstructor func(domain.RecordInfo, domain.Buffer) (domain.Record, error)

type RecordFactory struct {
	recordconstructor RecordConstructor
}

func NewRecordFactory(rc RecordConstructor) *RecordFactory {
	return &RecordFactory{
		recordconstructor: rc,
	}
}

func (rf *RecordFactory) New(buffer domain.Buffer) (domain.Record, error) {
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
