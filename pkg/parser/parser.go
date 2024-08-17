package parser

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type Parser struct {
	recordfactory   domain.RecordFactory
	questionfactory domain.QuestionFactory
}

func NewParser(rf domain.RecordFactory, qf domain.QuestionFactory) *Parser {
	return &Parser{
		recordfactory:   rf,
		questionfactory: qf,
	}
}

func (p *Parser) Parse(packet domain.Packet, buffer domain.Buffer) error {

	if err := p.readHeader(packet, buffer); err != nil {
		return err
	}

	if err := p.parseSections(buffer, packet); err != nil {
		return err
	}

	return nil

}
