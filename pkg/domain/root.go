package domain

type Buffer interface {
	ReadU16() (uint16, error)
	ReadU32() (uint32, error)
	ReadQName() (string, error)
}

// probably don't need all of these
type Packet interface {
	// Getters
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

	GetQuestions() []Question
	GetAnswers() []Record
	GetAuthorities() []Record
	GetResources() []Record

	// Setters
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

	AppendQuestion(question Question)
	AppendAnswer(answer Record)
	AppendAuthority(authority Record)
	AppendResource(resource Record)
}

type Question interface {
	// Getters
	GetName() string
	GetQType() QueryType
	// Setters
	SetName(string)
	SetQType(QueryType)
}

type Record interface {
}

type RecordInfo interface {
	GetQName() string
	GetQType() QueryType
	GetQClass() uint16
	GetTTL() uint32
	GetRDataLength() uint16
}

type Parser interface {
	Parse(*Packet, *Buffer) error
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
