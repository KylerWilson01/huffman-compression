// Package frequencymap holds everything related to a frequency map
package frequencymap

// FrequencyMap is the struct that hold information on the map
type FrequencyMap struct {
	input     string
	frequency map[rune]int
}

// NewFrequencyMap creates a pointer to a FrequencyMap
func NewFrequencyMap(input string) *FrequencyMap {
	return &FrequencyMap{
		input:     input,
		frequency: make(map[rune]int),
	}
}

// FindFrequencyOfChars fills out the map insidet the struct
func (fm *FrequencyMap) FindFrequencyOfChars() {
	fm.frequency = make(map[rune]int)

	for _, r := range fm.input {
		if _, ok := fm.frequency[r]; ok {
			fm.frequency[r]++
		} else {
			fm.frequency[r] = 1
		}
	}
}

// GetFrequency returns the frequency map
func (fm *FrequencyMap) GetFrequency() map[rune]int {
	return fm.frequency
}
