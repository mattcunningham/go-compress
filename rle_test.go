package compress

import "testing"

func TestRLE(t *testing.T) {
	if RLE("WWWWeeea") != "4"+NULL+"W"+NULL+"3"+NULL+"e"+NULL+"1"+NULL+"a"+NULL {
		t.Fatal("Test string did not match expected result.")
	}
}

func TestDeRLE(t *testing.T) {
	if DeRLE(RLE("WWWWeeea")) != "WWWWeeea" {
		t.Fatal("Test string did not match expected string.")
	}
}

func benchmarkRLE(testString string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		RLE(testString)
	}
}

func benchmarkDeRLE(testString string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DeRLE(testString)
	}
}

func BenchmarkRLE_Single(b *testing.B) { benchmarkRLE("b", b) }
func BenchmarkRLE_Repeat(b *testing.B) { benchmarkRLE("bbbAAAlllddd", b) }
func BenchmarkRLE_Large(b *testing.B)  { benchmarkRLE("abcdefghjiklmnopqrstuvwxyz0123456789", b) }

func BenchmarkDeRLE_Single(b *testing.B) { benchmarkDeRLE(RLE("b"), b) }
func BenchmarkDeRLE_Repeat(b *testing.B) { benchmarkDeRLE(RLE("bbbAAAlllddd"), b) }
func BenchmarkDeRLE_Large(b *testing.B) {
	benchmarkDeRLE(RLE("abcdefghjiklmnopqrstuvwxyz0123456789"), b)
}
