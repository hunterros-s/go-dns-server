package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/domain"
	"github.com/hunterros-s/go-dns-server/pkg/factory"
	"github.com/hunterros-s/go-dns-server/pkg/packet"
	"github.com/hunterros-s/go-dns-server/pkg/parser"
	"github.com/hunterros-s/go-dns-server/pkg/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <filename>")
		return
	}
	fileName := os.Args[1]
	packet, err := processFile(fileName)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

	displayResults(packet)
}

func processFile(fileName string) (domain.Packet, error) {
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

	rf := factory.NewRecordFactory(func(ri domain.RecordInfo, b domain.Buffer) (domain.Record, error) {
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

func displayResults(packet domain.Packet) {
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
