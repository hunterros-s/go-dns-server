package dns

type Buffer interface {
	Bytes() []byte

	ReadU16() (uint16, error)
	ReadU32() (uint32, error)
	ReadQName() (string, error)

	WriteByte(val byte) error
	WriteU8(val uint8) error
	WriteU16(val uint16) error
	WriteU32(val uint32) error
	WriteQName(qname string) error
	Set(pos uint16, val uint8) error
	SetU16(pos uint16, val uint16) error

	Step(steps uint16)
	Pos() uint16
}

type Header interface {
	GetID() uint16
	GetResponse() bool
	GetOpcode() uint8
	GetAuthoritativeAnswer() bool
	GetTruncatedMessage() bool
	GetRecursionDesired() bool
	GetRecursionAvailable() bool
	GetZ() uint8
	GetRescode() ResponseCode
	GetQuestionsCount() uint16
	GetAnswersCount() uint16
	GetAuthoritativeEntriesCount() uint16
	GetResourceEntriesCount() uint16

	SetID(id uint16)
	SetResponse(response bool)
	SetOpcode(opcode uint8)
	SetAuthoritativeAnswer(authoritativeAnswer bool)
	SetTruncatedMessage(truncatedMessage bool)
	SetRecursionDesired(recursionDesired bool)
	SetRecursionAvailable(recursionAvailable bool)
	SetZ(z uint8)
	SetRescode(rescode ResponseCode)
	SetQuestionsCount(questions uint16)
	SetAnswersCount(answers uint16)
	SetAuthoritativeEntriesCount(authoritativeEntries uint16)
	SetResourceEntriesCount(resourceEntries uint16)

	Write(Buffer) error
}

type Packet interface {
	GetHeader() Header
	GetQuestions() []Question
	GetAnswers() []Record
	GetAuthorities() []Record
	GetResources() []Record

	SetHeader(Header)
	AppendQuestion(question Question)
	AppendAnswer(answer Record)
	AppendAuthority(authority Record)
	AppendResource(resource Record)

	Write(Buffer) error
}

type Question interface {
	// Getters
	GetName() string
	GetQType() QueryType
	// Setters
	SetName(string)
	SetQType(QueryType)

	Write(Buffer) error
}

type Record interface {
	Write(Buffer) error
}

type RecordInfo interface {
	GetQName() string
	GetQType() QueryType
	GetQClass() uint16
	GetTTL() uint32
	GetRDataLength() uint16
}

type Parser interface {
	Parse(Packet, Buffer) error
}

type RecordFactory interface {
	New(Buffer) (Record, error)
}

type QuestionFactory interface {
	New(Buffer) (Question, error)
}

type RecordFactoryRegistry interface {
	Get(QueryType) (func(RecordInfo, Buffer) (Record, error), bool)
}

type Server struct {
	Address string
	Port    int
}

type UDPSocket interface {
	Bind(Server) error
	Unbind() error
	Send_to([]byte, Server) error
	Recv_from([]byte) (int, Server, error)
}
