package factory

import (
	"github.com/hunterros-s/go-dns-server/pkg/domain"
)

type QuestionConstructor func(string, domain.QueryType) domain.Question

type QuestionFactory struct {
	questionconstructor QuestionConstructor
}

func NewQuestionFactory(qc QuestionConstructor) *QuestionFactory {
	return &QuestionFactory{
		questionconstructor: qc,
	}
}

func (qf *QuestionFactory) New(buffer domain.Buffer) (domain.Question, error) {
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

	qt := domain.QueryType(qtype)

	q := qf.questionconstructor(name, qt)

	return q, nil
}
