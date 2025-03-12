package frequencymap_test

import (
	"testing"

	frequencymap "github.com/KylerWilson01/huffman-compression/internal/frequency-map"
)

func TestFindFrequencyOfChars(t *testing.T) {
	tests := []struct {
		name, input string
		want        map[rune]int
	}{
		{name: "pass", input: "this is a test", want: map[rune]int{
			't': 3,
			'h': 1,
			'i': 2,
			's': 3,
			'a': 1,
			'e': 1,
		}},
		{name: "pass with symbols", input: "!!this, \"is; a, test...", want: map[rune]int{
			't': 3,
			'h': 1,
			'i': 2,
			's': 3,
			'a': 1,
			'e': 1,
			'!': 2,
			',': 2,
			'"': 1,
			';': 1,
			'.': 3,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fm := frequencymap.NewFrequencyMap(test.input)
			fm.FindFrequencyOfChars()

			out := fm.GetFrequency()

			for char, freq := range test.want {
				if outFreq, ok := out[char]; !ok || freq != outFreq {
					t.Fatalf("%s not in out map", string(char))
				}
			}
		})
	}
}
