package main

import (
	"fmt"
	"go-proto-buffer/pb"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

func GenerateEmployee() *pb.Employee {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Meow",
		Email:       "meow@example.com",
		Occupation:  pb.Occupation_DESIGNER,
		PhoneNumber: []string{"000-0000-0000"},
		Project:     map[string]*pb.Company_Project{"MeowProject": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "I'm meow.",
		},
		Birthday: &pb.Date{
			Year:  2022,
			Month: 12,
			Day:   31,
		},
	}

	return employee
}

func Serialize(employee *pb.Employee) []byte {
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Serialize failed.", err)
	}
	return binData
}

func Deserialize(data []byte, message protoreflect.ProtoMessage) protoreflect.ProtoMessage {
	err := proto.Unmarshal(data, message)
	if err != nil {
		log.Fatalln("Deserialize failed.", err)
	}
	return message
}

func WriteFile(filename string, data []byte) {
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		log.Fatalln("Write file failed.", err)
	}
}

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Read file failed.", err)
	}
	return data
}

func ToJson(message protoiface.MessageV1) string {
	m := jsonpb.Marshaler{}
	res, err := m.MarshalToString(message)
	if err != nil {
		log.Fatalln("Marshal to json failed.", err)
	}
	return res

}

func fromJson(json string, message protoiface.MessageV1) protoiface.MessageV1 {

	if err := jsonpb.UnmarshalString(json, message); err != nil {
		log.Fatalln(("Unmarshal failed."))
	}

	return message
}

func main() {
	employee := GenerateEmployee()

	binData := Serialize(employee)
	WriteFile("test.bin", binData)

	data := ReadFile("test.bin")
	emp := Deserialize(data, &pb.Employee{})
	fmt.Println()

	json := ToJson(employee)
	res := fromJson(json, &pb.Employee{})

	fmt.Println(res)
}
