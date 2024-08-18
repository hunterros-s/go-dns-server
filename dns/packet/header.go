package packet

import "github.com/hunterros-s/go-dns-server/dns"

type Header struct {
	ID uint16 // 16 bits

	Response             bool             // 1 bit
	Opcode               uint8            // 4 bits
	Authoritative_answer bool             // 1 bit
	Truncated_message    bool             // 1 bit
	Recursion_desired    bool             // 1 bit
	Recursion_available  bool             // 1 bit
	Z                    uint8            // 3 bit
	Rescode              dns.ResponseCode // 4 bits

	Questions             uint16 // 16 bits
	Answers               uint16 // 16 bits
	Authoritative_entries uint16 // 16 bits
	Resource_entries      uint16 // 16 bits
}

func newHeader() dns.Header {
	return &Header{
		ID: 0,

		Response:             false,
		Opcode:               0,
		Authoritative_answer: false,
		Truncated_message:    false,
		Recursion_desired:    false,
		Recursion_available:  false,
		Z:                    0,
		Rescode:              dns.NOERROR,

		Questions:             0,
		Answers:               0,
		Authoritative_entries: 0,
		Resource_entries:      0,
	}
}

// write
func (h *Header) Write(buffer dns.Buffer) error {
	if err := buffer.WriteU16(h.ID); err != nil {
		return err
	}

	flags1 := uint8(0)
	if h.Recursion_desired {
		flags1 |= 1 << 0
	}
	if h.Truncated_message {
		flags1 |= 1 << 1
	}
	if h.Authoritative_answer {
		flags1 |= 1 << 2
	}
	flags1 |= h.Opcode << 3
	if h.Response {
		flags1 |= 1 << 7
	}

	if err := buffer.WriteU8(flags1); err != nil {
		return err
	}

	flags2 := uint8(0)
	flags2 |= uint8(h.Rescode)
	// if h.CheckingDisabled {
	// 	flags2 |= 1 << 4
	// }
	// if h.AuthedData {
	// 	flags2 |= 1 << 5
	// }
	flags2 |= h.Z << 4 // would be 6 if it's 1 bit
	if h.Recursion_available {
		flags2 |= 1 << 7
	}

	if err := buffer.WriteU8(flags2); err != nil {
		return err
	}

	if err := buffer.WriteU16(h.Questions); err != nil {
		return err
	}

	if err := buffer.WriteU16(h.Answers); err != nil {
		return err
	}

	if err := buffer.WriteU16(h.Authoritative_entries); err != nil {
		return err
	}

	if err := buffer.WriteU16(h.Resource_entries); err != nil {
		return err
	}

	return nil
}

// Getter methods
func (h *Header) GetID() uint16 {
	return h.ID
}

func (h *Header) GetResponse() bool {
	return h.Response
}

func (h *Header) GetOpcode() uint8 {
	return h.Opcode
}

func (h *Header) GetAuthoritativeAnswer() bool {
	return h.Authoritative_answer
}

func (h *Header) GetTruncatedMessage() bool {
	return h.Truncated_message
}

func (h *Header) GetRecursionDesired() bool {
	return h.Recursion_desired
}

func (h *Header) GetRecursionAvailable() bool {
	return h.Recursion_available
}

func (h *Header) GetZ() uint8 {
	return h.Z
}

func (h *Header) GetRescode() dns.ResponseCode {
	return h.Rescode
}

func (h *Header) GetQuestionsCount() uint16 {
	return h.Questions
}

func (h *Header) GetAnswersCount() uint16 {
	return h.Answers
}

func (h *Header) GetAuthoritativeEntriesCount() uint16 {
	return h.Authoritative_entries
}

func (h *Header) GetResourceEntriesCount() uint16 {
	return h.Resource_entries
}

// Setter methods
func (h *Header) SetID(id uint16) {
	h.ID = id
}

func (h *Header) SetResponse(response bool) {
	h.Response = response
}

func (h *Header) SetOpcode(opcode uint8) {
	h.Opcode = opcode
}

func (h *Header) SetAuthoritativeAnswer(authoritativeAnswer bool) {
	h.Authoritative_answer = authoritativeAnswer
}

func (h *Header) SetTruncatedMessage(truncatedMessage bool) {
	h.Truncated_message = truncatedMessage
}

func (h *Header) SetRecursionDesired(recursionDesired bool) {
	h.Recursion_desired = recursionDesired
}

func (h *Header) SetRecursionAvailable(recursionAvailable bool) {
	h.Recursion_available = recursionAvailable
}

func (h *Header) SetZ(z uint8) {
	h.Z = z
}

func (h *Header) SetRescode(rescode dns.ResponseCode) {
	h.Rescode = rescode
}

func (h *Header) SetQuestionsCount(questions uint16) {
	h.Questions = questions
}

func (h *Header) SetAnswersCount(answers uint16) {
	h.Answers = answers
}

func (h *Header) SetAuthoritativeEntriesCount(authoritativeEntries uint16) {
	h.Authoritative_entries = authoritativeEntries
}

func (h *Header) SetResourceEntriesCount(resourceEntries uint16) {
	h.Resource_entries = resourceEntries
}
