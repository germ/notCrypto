package notCrypto

// Return a seed for the (not)RNG from a string
func genSeed(in string) int {
	var acc int
	for i, v := range in {
		acc += int(v)*(i+1)
	}

	return acc
}

//Create pad that can be read from indefinatly
func genPad(seed int) chan int {
	c := make(chan int)
	m_z, m_w := seed, seed
	go func() {
		for {
			m_z = 36969 * (m_z & 65535) + (m_z >> 16);
			m_w = 18000 * (m_w & 65535) + (m_w >> 16);
			c <- ((m_z << 16) + m_w) % 95
		}
	}()

	return c
}

//Fill a slice with values created from a given seed
func genSlice(seed, length int) []int {
	res := make([]int, length)
	rng := genPad(seed)

	for i := 0; i < length; i++ {
		res[i] = <-rng
	}

	return res
}
