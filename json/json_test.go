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
				r, err := getFileReader(fmt.Sprintf("data/%d.json", numChar))
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

func getFileReader(fileName string) (io.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return f, nil
}
