package packet

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type Header struct {
	ID uint16 // 16 bits

	Response             bool                // 1 bit
	Opcode               uint8               // 4 bits
	Authoritative_answer bool                // 1 bit
	Truncated_message    bool                // 1 bit
	Recursion_desired    bool                // 1 bit
	Recursion_available  bool                // 1 bit
	Z                    uint8               // 3 bit
	Rescode              domain.ResponseCode // 4 bits

	Questions             uint16 // 16 bits
	Answers               uint16 // 16 bits
	Authoritative_entries uint16 // 16 bits
	Resource_entries      uint16 // 16 bits
}

func newHeader() Header {
	return Header{
		ID: 0,

		Response:             false,
		Opcode:               0,
		Authoritative_answer: false,
		Truncated_message:    false,
		Recursion_desired:    false,
		Recursion_available:  false,
		Z:                    0,
		Rescode:              domain.NOERROR,

		Questions:             0,
		Answers:               0,
		Authoritative_entries: 0,
		Resource_entries:      0,
	}
}

func (p *Packet) GetID() uint16 {
	return p.Header.ID
}

func (p *Packet) GetResponse() bool {
	return p.Header.Response
}

func (p *Packet) GetOpcode() uint8 {
	return p.Header.Opcode
}

func (p *Packet) GetAuthoritativeAnswer() bool {
	return p.Header.Authoritative_answer
}

func (p *Packet) GetTruncatedMessage() bool {
	return p.Header.Truncated_message
}

func (p *Packet) GetRecursionDesired() bool {
	return p.Header.Recursion_desired
}

func (p *Packet) GetRecursionAvailable() bool {
	return p.Header.Recursion_available
}

func (p *Packet) GetZ() uint8 {
	return p.Header.Z
}

func (p *Packet) GetRescode() domain.ResponseCode {
	return p.Header.Rescode
}

func (p *Packet) GetQuestionsCount() uint16 {
	return p.Header.Questions
}

func (p *Packet) GetAnswersCount() uint16 {
	return p.Header.Answers
}

func (p *Packet) GetAuthoritativeEntriesCount() uint16 {
	return p.Header.Authoritative_entries
}

func (p *Packet) GetResourceEntriesCount() uint16 {
	return p.Header.Resource_entries
}

func (p *Packet) SetID(id uint16) {
	p.Header.ID = id
}

func (p *Packet) SetResponse(response bool) {
	p.Header.Response = response
}

func (p *Packet) SetOpcode(opcode uint8) {
	p.Header.Opcode = opcode
}

func (p *Packet) SetAuthoritativeAnswer(authoritativeAnswer bool) {
	p.Header.Authoritative_answer = authoritativeAnswer
}

func (p *Packet) SetTruncatedMessage(truncatedMessage bool) {
	p.Header.Truncated_message = truncatedMessage
}

func (p *Packet) SetRecursionDesired(recursionDesired bool) {
	p.Header.Recursion_desired = recursionDesired
}

func (p *Packet) SetRecursionAvailable(recursionAvailable bool) {
	p.Header.Recursion_available = recursionAvailable
}

func (p *Packet) SetZ(z uint8) {
	p.Header.Z = z
}

func (p *Packet) SetRescode(rescode domain.ResponseCode) {
	p.Header.Rescode = rescode
}

func (p *Packet) SetQuestionsCount(questions uint16) {
	p.Header.Questions = questions
}

func (p *Packet) SetAnswersCount(answers uint16) {
	p.Header.Answers = answers
}

func (p *Packet) SetAuthoritativeEntriesCount(authoritativeEntries uint16) {
	p.Header.Authoritative_entries = authoritativeEntries
}

func (p *Packet) SetResourceEntriesCount(resourceEntries uint16) {
	p.Header.Resource_entries = resourceEntries
}
