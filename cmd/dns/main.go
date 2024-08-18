package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hunterros-s/go-dns-server/dns"
	"github.com/hunterros-s/go-dns-server/dns/buffer"
	"github.com/hunterros-s/go-dns-server/dns/factory"
	"github.com/hunterros-s/go-dns-server/dns/packet"
	"github.com/hunterros-s/go-dns-server/dns/parser"
	"github.com/hunterros-s/go-dns-server/dns/registry"
	"github.com/hunterros-s/go-dns-server/dns/udp"
)

func main() {

	log.Println("starting")

	qname := "www.yahoo.com"
	qtype := dns.MX

	server := dns.Server{Address: "8.8.8.8", Port: 53}

	log.Println("creating socket")
	socket := udp.NewUDPSocket()
	err := socket.Bind(dns.Server{Address: "0.0.0.0", Port: 43210})
	if err != nil {
		log.Println(err.Error())
		return
	}

	p := packet.NewPacket()

	p.GetHeader().SetID(12345)
	p.GetHeader().SetRecursionDesired(true)
	p.AppendQuestion(packet.NewQuestion(qname, qtype))

	log.Println("outgoing packet:")
	displayResults(p)

	request_buffer := buffer.NewPacketBuffer([]byte{})
	err = p.Write(request_buffer)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("sending to socket")
	err = socket.Send_to(request_buffer.Bytes(), server)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("recv. from socket")

	data := make([]byte, 512)

	len, server, err := socket.Recv_from(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("recieved %d bytes\n", len)
	log.Println(server.Address)
	log.Println(server.Port)

	log.Println(hex.EncodeToString(data[:len]))

	rec_buffer := buffer.NewPacketBuffer(data[:len])

	log.Println("creating packet from response")
	p_back := packet.NewPacket()
	r_registry := registry.RecordRegistry{}

	rf := factory.NewRecordFactory(func(ri dns.RecordInfo, b dns.Buffer) (dns.Record, error) {
		factory_func, ok := r_registry.Get(ri.GetQType())
		if !ok {
			log.Println("unknown packet seen")
			return registry.New_unknown_record(ri, b)
		}
		return factory_func(ri, b)
	})
	qf := factory.NewQuestionFactory(packet.NewQuestion)

	parser := parser.NewParser(rf, qf)

	parser.Parse(p_back, rec_buffer)

	displayResults(p_back)
}

func processFile(fileName string) (dns.Packet, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	b := make([]byte, 512)
	_, err = file.Read(b)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	pb := buffer.NewPacketBuffer(b)

	p := packet.NewPacket()
	r_registry := registry.RecordRegistry{}

	rf := factory.NewRecordFactory(func(ri dns.RecordInfo, b dns.Buffer) (dns.Record, error) {
		factory_func, ok := r_registry.Get(ri.GetQType())
		if !ok {
			log.Println("unknown packet seen")
			return registry.New_unknown_record(ri, b)
		}
		return factory_func(ri, b)
	})
	qf := factory.NewQuestionFactory(packet.NewQuestion)

	parser := parser.NewParser(rf, qf)

	parser.Parse(p, pb)

	return p, nil
}

func displayResults(packet dns.Packet) {
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
