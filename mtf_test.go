package compress

import "testing"

const (
	hamlet = `To be or not to be-that is the question:
Whether 'tis nobler in the mind to suffer
The slings and arrows of outrageous fortune,
Or to take arms against a sea of troubles,
And, by opposing, end them. To die, to sleep-
No more-and by a sleep to say we end
The heartache and the thousand natural shocks
That flesh is heir to-'tis a consummation
Devoutly to be wished. To die, to sleep-
To sleep, perchance to dream. Aye, there's the rub,
For in that sleep of death what dreams may come,
When we have shuffled off this mortal coil,
Must give us pause. There's the respect
That makes calamity of so long life.`
)

func TestMtF(t *testing.T) {
	compressed := MtF(hamlet)
	decompressed := DeMtF(compressed)
	if decompressed != hamlet {
		t.Fatal("The decompressed version of hamlet does not equal the original.")
	}
}

func benchmarkMtF(benchmarkString string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		MtF(benchmarkString)
	}
}

func benchmarkDeMtF(benchmarkSlice []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DeMtF(benchmarkSlice)
	}
}

func BenchmarkMtF_Single(b *testing.B) { benchmarkMtF("z", b) }      // single character
func BenchmarkMtF_Random(b *testing.B) { benchmarkMtF("l^cu8&", b) } // these characters are far apart and non-repeating
func BenchmarkMtF_Repeat(b *testing.B) { benchmarkMtF("AAAAAA", b) } // these characters repeat, after finding first char
func BenchmarkMtF_Large(b *testing.B)  { benchmarkMtF(hamlet, b) }   // larger string

func BenchmarkDeMtF_Single(b *testing.B) { benchmarkDeMtF(MtF("z"), b) }
func BenchmarkDeMtF_Random(b *testing.B) { benchmarkDeMtF(MtF("l^cu8&"), b) }
func BenchmarkDeMtF_Repeat(b *testing.B) { benchmarkDeMtF(MtF("AAAAAA"), b) }
func BenchmarkDeMtF_Large(b *testing.B)  { benchmarkDeMtF(MtF(hamlet), b) }
