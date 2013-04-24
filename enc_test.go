package notCrypto

import (
	"testing"
	"io/ioutil"
)

const (
	inputFile  = "tests/input"
	outputFile = "tests/output"
)

func TestEncode(t *testing.T) {
	a := Encode("love", "you")
	b := Decode("love", a)
	t.Log(a)
	t.Log(b)

	if b != "you" {
		t.Fail()
	}
}

func TestEncodeFile(t *testing.T) {
	a, err := EncodeFile(inputFile, "love")
	t.Log(a, err)
	if err != nil {
		t.FailNow()
	}

	b, err := DecodeFile(outputFile, "love")
	t.Log(b, err)
	if err != nil {
		t.FailNow()
	}

	cmp, err := ioutil.ReadFile(inputFile)
	if string(cmp[:len(cmp)-1]) != b {
		t.Fail()
	}
}
