package compress

import "testing"

const (
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla tincidunt eros est, quis condimentum lorem lobortis ut. Nam convallis tortor pretium quam ornare, quis mollis risus ultricies. Nunc ullamcorper eu mi eu bibendum. Praesent tincidunt, libero eget blandit interdum, arcu ex tempor erat, eu tristique magna odio ut nunc. Ut ultricies lorem ut felis scelerisque ultricies. Quisque ornare turpis ac augue condimentum ultrices. Nullam diam diam, mollis vel ligula at, condimentum sagittis metus. Quisque volutpat pulvinar risus, sit amet pulvinar nisi gravida eget. Maecenas ut ullamcorper tortor. Donec faucibus auctor sapien ut sagittis. Cras ut fermentum tellus, eu aliquam quam. Cras dignissim vitae lorem id tempor. Mauris imperdiet sapien purus, id aliquet sapien porta egestas. Donec dignissim facilisis elit vel facilisis.

Nunc eu ligula libero. Ut bibendum malesuada aliquam. Etiam semper felis neque, a ullamcorper turpis elementum eu. Nulla ut pellentesque dolor. Vivamus tempus a turpis quis volutpat. Maecenas tincidunt congue lacinia. In maximus suscipit elit id fringilla. Etiam dolor purus, egestas nec hendrerit sed, malesuada nec diam. Curabitur efficitur sed felis id bibendum.`
)

func TestBWT(t *testing.T) {
	if BWT("^banana!") != "!a\u0004nnb^aa" {
		t.Error("BWT did not produce the expected output. This failure might be incorrect if the algorithm was adjusted.")
	}
}

func TestBWT_Full(t *testing.T) {
	compressed := BWT(lorem)
	decompressed := DeBWT(compressed)
	if decompressed != lorem {
		t.Fatal("The decompressed lorem ipsum does not equal the original.")
	}
}

func benchmarkBWT(testString string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BWT(testString)
	}
}

func benchmarkDeBWT(testString string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DeBWT(testString)
	}
}

func BenchmarkBWT_Single(b *testing.B)   { benchmarkBWT("a", b) }
func BenchmarkBWT_NoRepeat(b *testing.B) { benchmarkBWT("d*3N)a", b) }
func BenchmarkBWT_Repeat(b *testing.B)   { benchmarkBWT("abbbbc", b) }
func BenchmarkBWT_Large(b *testing.B)    { benchmarkBWT(lorem, b) }

func BenchmarkDeBWT_Single(b *testing.B)   { benchmarkDeBWT(BWT("a"), b) }
func BenchmarkDeBWT_NoRepeat(b *testing.B) { benchmarkDeBWT(BWT("d*3N)a"), b) }
func BenchmarkDeBWT_Repeat(b *testing.B)   { benchmarkDeBWT(BWT("abbbbc"), b) }
func BenchmarkDeBWT_Large(b *testing.B)    { benchmarkDeBWT(BWT(lorem), b) }
