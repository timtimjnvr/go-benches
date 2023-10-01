package json

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func BenchmarkDecodeWithDecoder(b *testing.B) {
	for numChar := 0; numChar < 1e5; numChar += 1e3 {
		b.Run(fmt.Sprintf("%d-characters", numChar), func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				r, _, err := getFileReader(fmt.Sprintf("data/%d.json", numChar))
				if err != nil {
					log.Fatal(err)
				}
				b.StartTimer()

				c, err := decodeWithDecoder(r)
				if err != nil {
					log.Fatal(err)
				}

				_ = c
			}
		})
	}
}

func getFileReader(fileName string) (io.Reader, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, err
	}

	i, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	return f, int(i.Size()), err
}
