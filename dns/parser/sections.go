package parser

import "github.com/hunterros-s/go-dns-server/dns"

func (p *Parser) parseSections(buffer dns.Buffer, packet dns.Packet) error {
	if err := p.parseQuestions(buffer, packet); err != nil {
		return err
	}
	if err := p.parseAnswers(buffer, packet); err != nil {
		return err
	}
	if err := p.parseAuthorities(buffer, packet); err != nil {
		return err
	}
	if err := p.parseResources(buffer, packet); err != nil {
		return err
	}
	return nil
}

func (p *Parser) parseQuestions(buffer dns.Buffer, packet dns.Packet) error {
	for i := 0; i < int(packet.GetHeader().GetQuestionsCount()); i++ {
		question, err := p.questionfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendQuestion(question)
	}
	return nil
}

func (p *Parser) parseAnswers(buffer dns.Buffer, packet dns.Packet) error {
	for i := 0; i < int(packet.GetHeader().GetAnswersCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendAnswer(record)
	}
	return nil
}

func (p *Parser) parseAuthorities(buffer dns.Buffer, packet dns.Packet) error {
	for i := 0; i < int(packet.GetHeader().GetAuthoritativeEntriesCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendAuthority(record)
	}
	return nil
}

func (p *Parser) parseResources(buffer dns.Buffer, packet dns.Packet) error {
	for i := 0; i < int(packet.GetHeader().GetResourceEntriesCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendResource(record)
	}
	return nil
}
