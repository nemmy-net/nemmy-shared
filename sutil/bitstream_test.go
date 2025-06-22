package sutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestBitstreamEncode(t *testing.T) {
	answer := []byte{
		0x99, 0x22, 0x33, 0x44, 0x55, 0x66, 0xCA, 0xDE, 'H', 'e', 'l', 'l', 'o', '!',
	}
	stream := BitStream{
		Buffer: make([]byte, len(answer)),
	}

	err := stream.PutUint(0x992233445566CADE, 8)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = stream.PutStr("Hello!")
	if err != nil {
		t.Fatalf("%v", err)
	}

	isMatch := len(answer) == len(stream.Buffer)
	if isMatch {
		for i := 0; i < len(answer); i++ {
			if answer[i] != stream.Buffer[i] {
				isMatch = false
				break
			}
		}
	}

	if !isMatch {
		answerStr := strings.Builder{}
		bufferStr := strings.Builder{}
		for i := 0; i < len(answer); i++ {
			answerStr.WriteString(fmt.Sprintf("%x ", answer[i]))
			bufferStr.WriteString(fmt.Sprintf("%x ", stream.Buffer[i]))
		}

		t.Fatalf("Incorrect bytes.\nExpected: %v\nGot     : %v", answerStr, bufferStr)
	}

	// Test generic Put method
	stream = BitStream{
		Buffer: make([]byte, len(answer)),
	}

	err = stream.Put(uint64(0x992233445566CADE))
	if err != nil {
		t.Fatal(err)
	}

	u64Answer := uint64(0x992233445566CADE)
	u64, err := stream.GetUint(8)
	if err != nil {
		t.Fatal(err)
	}

	if u64 != u64Answer {
		t.Fatalf("Expected %v, got %v", u64Answer, u64)
	}
}

func TestBitstreamDecode(t *testing.T) {
	answerU64 := uint64(0x992233445566CADE)
	answerStr := "Hello!"
	stream := BitStream{
		Buffer: []byte{
			0x99, 0x22, 0x33, 0x44, 0x55, 0x66, 0xCA, 0xDE, 'H', 'e', 'l', 'l', 'o', '!',
		},
	}

	u64, err := stream.GetUint(8)
	if err != nil {
		t.Fatalf("%v", err)
	} else if u64 != answerU64 {
		t.Fatalf("Expected %v, got %v", answerU64, u64)
	}
	str, err := stream.GetStr(len(answerStr))
	if err != nil {
		t.Fatalf("%v", err)
	} else if str != answerStr {
		t.Fatalf("Expected %v, got %v", answerStr, str)
	}

	// Test the generic Get method
	stream.ReadPos = 0
	var u16Answer uint16 = 0x9922
	var i32Answer int32 = 0x33445566
	var i16Answer int16 = -13602
	var u16 uint16
	var i32 int32
	var i16 int16
	err = stream.Get(&u16, &i32, &i16)

	if err != nil {
		t.Fatal(err)
	}

	if u16 != u16Answer ||
		i32 != i32Answer ||
		i16 != i16Answer {
		t.Fatalf(
			"Expected:"+
				"\n\tu16=%v"+
				"\n\ti32=%v"+
				"\n\ti16=%v"+
				"\nGot:"+
				"\n\tu16=%v"+
				"\n\ti32=%v"+
				"\n\ti16=%v",
			u16Answer, i32Answer, i16Answer,
			u16, i32, i16,
		)
	}
}
