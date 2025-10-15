package error_example

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcRemainderAndMod(t *testing.T) {
	numerator := 20
	denominator := 3
	remainer, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainer, mod) // 6 2

	denominator = 0
	assert := assert.New(t)
	_, _, err = calcRemainderAndMod(20, 0)
	assert.NotNil(err)
	fmt.Println(err) // denominator is 0
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		// return 0, 0, errors.New("denominator is 0") // use string for simple errors
		return 0, 0, fmt.Errorf("denominator is %d", 0)
	}
	return numerator / denominator, numerator % denominator, nil
}

//
// sentinel errors
//

func TestSentinelErrors(t *testing.T) {
	data := []byte("This a not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("told you so")
	}
}

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrFoo = Sentinel("foo error")
	ErrBar = Sentinel("bad error")
)
