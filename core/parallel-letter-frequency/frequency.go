package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of
// each rune in a given text using concurrent routines
// and returns this data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	m := make(FreqMap)
	c := make(chan FreqMap, len(s))

	// 1 routine per string receives the string as arg and returns the map on the channel
	for i := range s {
		go func(s string, c chan<- FreqMap) {
			c <- Frequency(s)
		}(s[i], c)
	}

	// Later we read from the channel and accumulate the results
	for range s {
		f := <-c // Just for clarity receiving the map first and then iterating over it.
		for k, v := range f {
			m[k] += v
		}
	}

	return m
}
