package notCrypto

import "io/ioutil"

// Encrypt a given string using a raw key
func Encode(seed, data string) string {
	pad := genSlice(genSeed(seed), len(data))
	var res string
	
	for i, v := range data {
		//Hardcoded ASCII. 
		c := ((int(v) - asciiStart) + pad[i]) % asciiPrint
		res += string(asciiStart + c)
	}

	return res
}

// Decode a given string using a raw key
func Decode(seed, data string) string {
	pad := genSlice(genSeed(seed), len(data))
	var res string
	
	for i, v := range data {
		c := (int(v) - asciiStart) - pad[i] 
		if c < 0 {
			res += string(asciiEnd + (c + 1))
		} else {
			res += string(c + asciiStart)
		}
	}
	return res
}

// Wrapper for handling files
func EncodeFile(name, key string) (string, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	f = f[:len(f)-1]
	return Encode(key, string(f)), nil
}

// Yet another wrapper. There has to be a way to merge
// this with EncodeFile
func DecodeFile(name, key string) (string, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	f = f[:len(f)-1]
	return Decode(key, string(f)), nil
}
