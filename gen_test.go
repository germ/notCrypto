package notCrypto

import (
	"testing"
)

func TestGenSeed(t *testing.T) {
	args := []string{"02/01/13", "asdf", "waka!", "Stop Smoking"}

	t.Log("genSeed:")
	for _, v := range args {
		t.Log(genSeed(v))
	}
}

func TestGenPad(t *testing.T) {
	args := []string{"a", "b", "b", "Stop Smoking"}

	t.Log("genPad:")
	for _, v := range args {
		c := genPad(genSeed(v))
		n := []int{}
		for j := 0; j < 5; j++ {
			n = append(n, <-c)
		}
		t.Log()
	}
}

func TestGenSlice(t *testing.T) {
	t.Log(genSlice(72, 5))
}
