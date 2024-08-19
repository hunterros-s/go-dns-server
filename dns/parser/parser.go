package parser

import "github.com/hunterros-s/go-dns-server/dns"

type Parser struct {
	recordfactory   dns.RecordFactory
	questionfactory dns.QuestionFactory
}

func NewParser(rf dns.RecordFactory, qf dns.QuestionFactory) dns.Parser {
	return &Parser{
		recordfactory:   rf,
		questionfactory: qf,
	}
}

func (p *Parser) Parse(packet dns.Packet, buffer dns.Buffer) error {

	if err := p.readHeader(packet, buffer); err != nil {
		return err
	}

	if err := p.parseSections(buffer, packet); err != nil {
		return err
	}

	return nil

}
