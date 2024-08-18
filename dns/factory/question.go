package factory

import (
	"github.com/hunterros-s/go-dns-server/dns"
)

type QuestionConstructor func(string, dns.QueryType) dns.Question

type QuestionFactory struct {
	questionconstructor QuestionConstructor
}

func NewQuestionFactory(qc QuestionConstructor) *QuestionFactory {
	return &QuestionFactory{
		questionconstructor: qc,
	}
}

func (qf *QuestionFactory) New(buffer dns.Buffer) (dns.Question, error) {
	name, err := buffer.ReadQName()
	if err != nil {
		return nil, err
	}

	qtype, err := buffer.ReadU16()
	if err != nil {
		return nil, err
	}

	_, err = buffer.ReadU16() // class (we're not using this value)
	if err != nil {
		return nil, err
	}

	qt := dns.QueryType(qtype)

	q := qf.questionconstructor(name, qt)

	return q, nil
}
