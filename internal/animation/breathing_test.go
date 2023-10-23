package animation

import "testing"

func TestBreathingFunction(t *testing.T) {
	tests := []struct {
		name       string
		function   BreathingFunction
		wantValues []uint16
	}{
		{
			"linear",
			BreatheLinear,
			[]uint16{255, 128, 0, 127, 255},
		},
		{
			"sine",
			BreatheSine,
			[]uint16{255, 127, 0, 127, 255},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			durations := calculateDurations(len(test.wantValues), 1000)
			for i := range test.wantValues {
				actual := test.function(durations[i], 1000) >> 8
				expected := test.wantValues[i]
				if actual != expected {
					t.Errorf("got != want: %3d: %3d != %3d", durations[i], actual, expected)
				} else {
					t.Logf("got == want: %3d: %3d == %3d", durations[i], actual, expected)
				}
			}
		})

		t.Run(test.name+"_zero", func(t *testing.T) {
			var min, max uint8 = 0xFF, 0
			var minT, maxT Milliseconds = 0, 0
			for i := 0; i <= 1000; i++ {
				actual := uint8(test.function(Milliseconds(i), 1000) >> 8)
				if actual < min {
					min = actual
					minT = Milliseconds(i)
				}
				if actual > max {
					max = actual
					maxT = Milliseconds(i)
				}
			}

			if min > 0x00+1 {
				t.Errorf("min > 0x00+1, lowest is %d at %d", min, minT)
			} else {
				t.Log("min is", min, "at", minT)
			}
			if max < 0xFF-1 {
				t.Errorf("max < 0xFF-1, highest is %d at %d", max, maxT)
			} else {
				t.Log("max is", max, "at", maxT)
			}
		})
	}
}

func calculateDurations(length, max int) []Milliseconds {
	durations := make([]Milliseconds, length)
	for i := 0; i < length; i++ {
		durations[i] = Milliseconds(max * i / (length - 1))
	}
	return durations
}

func BenchmarkBreatheSine(b *testing.B) {
	b.Run("fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			breatheSineFast(500, 1000)
		}
	})
	b.Run("accurate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			breatheSineAccurate(500, 1000)
		}
	})
}
