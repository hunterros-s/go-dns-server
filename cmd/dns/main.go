package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hunterros-s/go-dns-server/dns"
	"github.com/hunterros-s/go-dns-server/dns/buffer"
	"github.com/hunterros-s/go-dns-server/dns/factory"
	"github.com/hunterros-s/go-dns-server/dns/packet"
	"github.com/hunterros-s/go-dns-server/dns/parser"
	"github.com/hunterros-s/go-dns-server/dns/registry"
	"github.com/hunterros-s/go-dns-server/dns/udp"
)

func main() {

	r_registry := registry.RecordRegistry{}

	rf := factory.NewRecordFactory(func(ri dns.RecordInfo, b dns.Buffer) (dns.Record, error) {
		factory_func, ok := r_registry.Get(ri.GetQType())
		if !ok {
			fmt.Printf("Unknown packet with qtype: %d\n", ri.GetQType())
			return registry.New_unknown_record(ri, b)
		}
		return factory_func(ri, b)
	})
	qf := factory.NewQuestionFactory(packet.NewQuestion)

	parser := parser.NewParser(rf, qf)

	socket := udp.NewUDPSocket()
	err := socket.Bind(dns.Server{Address: "0.0.0.0", Port: 2053})
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		err = handleQuery(socket, parser)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func lookup(qname string, qtype dns.QueryType, parser dns.Parser) (dns.Packet, error) {
	server := dns.Server{Address: "8.8.8.8", Port: 53}

	socket := udp.NewUDPSocket()
	err := socket.Bind(dns.Server{Address: "0.0.0.0", Port: 43210})
	defer socket.Unbind()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	p := packet.NewPacket()

	p.GetHeader().SetID(12345)
	p.GetHeader().SetRecursionDesired(true)
	p.AppendQuestion(packet.NewQuestion(qname, qtype))

	request_buffer := buffer.NewPacketBuffer([]byte{})
	err = p.Write(request_buffer)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	err = socket.Send_to(request_buffer.Bytes(), server)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	data := make([]byte, 512)

	len, _, err := socket.Recv_from(data)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	rec_buffer := buffer.NewPacketBuffer(data[:len])

	p_back := packet.NewPacket()

	parser.Parse(p_back, rec_buffer)

	return p_back, nil
}

func handleQuery(socket dns.UDPSocket, parser dns.Parser) error {
	data := make([]byte, 512)
	n, src, err := socket.Recv_from(data)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req_buffer := buffer.NewPacketBuffer(data[:n])

	request := packet.NewPacket()

	parser.Parse(request, req_buffer)

	log.Printf("request packet: ")
	displayPacket(request)

	response := packet.NewPacket()
	response.GetHeader().SetID(request.GetHeader().GetID())
	response.GetHeader().SetRecursionAvailable(true)
	response.GetHeader().SetRecursionDesired(true)
	response.GetHeader().SetResponse(true)

	questions := request.GetQuestions()
	if len(questions) > 0 {
		question := questions[len(questions)-1]

		r, err := lookup(question.GetName(), question.GetQType(), parser)
		if err != nil {
			response.GetHeader().SetRescode(dns.SERVAIL)
		} else {
			response.AppendQuestion(question)
			response.GetHeader().SetRescode(r.GetHeader().GetRescode())

			for _, rec := range r.GetAnswers() {
				log.Printf("Answer: ")
				displayAsJSON(rec)
				response.AppendAnswer(rec)
			}
			for _, rec := range r.GetAuthorities() {
				log.Printf("Authority: ")
				displayAsJSON(rec)
				response.AppendAuthority(rec)
			}
			for _, rec := range r.GetResources() {
				log.Printf("Resource: ")
				displayAsJSON(rec)
				response.AppendResource(rec)
			}
		}

	} else {
		response.GetHeader().SetRescode(dns.FORMERR)
	}

	response.GetHeader().SetQuestionsCount(uint16(len(response.GetQuestions())))
	response.GetHeader().SetAnswersCount(uint16(len(response.GetAnswers())))
	response.GetHeader().SetAuthoritativeEntriesCount(uint16(len(response.GetAuthorities())))
	response.GetHeader().SetResourceEntriesCount(uint16(len(response.GetResources())))

	log.Printf("response packet: ")
	displayPacket(response)

	response_buffer := buffer.NewPacketBuffer([]byte{})
	err = response.Write(response_buffer)
	if err != nil {
		return err
	}

	log.Println(hex.EncodeToString(response_buffer.Bytes()))

	socket.Send_to(response_buffer.Bytes(), src)
	return nil
}

func displayPacket(packet dns.Packet) {
	fmt.Println("main:")
	displayAsJSON(packet)

	fmt.Println("questions:")
	for _, v := range packet.GetQuestions() {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("answers:")
	for _, v := range packet.GetAnswers() {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("authorities:")
	for _, v := range packet.GetAuthorities() {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("resources:")
	for _, v := range packet.GetResources() {
		fmt.Printf("%#v\n", v)
	}
}

func displayAsJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}
