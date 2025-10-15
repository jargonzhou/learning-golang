package example_lg

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Example of real world. - 'Learning Go': 12. Concurrency in Go, P.309
// pipeline: 3 services(A,B,C), call 2 of them, then take their result to the third, result its result.
// timeout: 50ms

type Input struct {
	A string
	B string
}

type COut struct {
	// frequenceCount map[rune]int
}

// A, B processor
type aOut struct {
}

type bOut struct {
}

type abProcessor struct {
	outA chan aOut
	outB chan bOut
	err  chan error
}

func newABProcessor() *abProcessor {
	return &abProcessor{
		outA: make(chan aOut, 1),
		outB: make(chan bOut, 1),
		err:  make(chan error, 2), // buffer 2 errors at most
	}
}

func (p *abProcessor) start(ctx context.Context, data Input) {
	// invoke A
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.err <- err
			return
		}
		p.outA <- aOut
	}()

	// invoke B
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.err <- err
			return
		}
		p.outB <- bOut
	}()
}

func (p *abProcessor) wait(ctx context.Context) (cIn, error) {
	var cData cIn
	for count := 0; count < 2; count++ { // 2 times
		select {
		case a := <-p.outA:
			cData.a = a
		case b := <-p.outB:
			cData.b = b
		case err := <-p.err:
			return cIn{}, err
		case <-ctx.Done():
			return cIn{}, ctx.Err()
		}
	}
	return cData, nil
}

// mock result: A, B
func getResultA(ctx context.Context, in string) (aOut, error) {
	return aOut{}, nil
}

func getResultB(ctx context.Context, in string) (bOut, error) {
	return bOut{}, nil
}

// C processor

type cIn struct {
	a aOut
	b bOut
}

type cProcessor struct {
	outC chan COut
	err  chan error
}

func newCProcessor() *cProcessor {
	return &cProcessor{
		outC: make(chan COut, 1),
		err:  make(chan error, 1),
	}
}

func (p *cProcessor) start(ctx context.Context, input cIn) {
	go func() {
		cOut, err := getResultC(ctx, input)
		if err != nil {
			p.err <- err
		}
		p.outC <- cOut
	}()
}

func (p *cProcessor) wait(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.err:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

// mock result: C
func getResultC(ctx context.Context, c cIn) (COut, error) {
	return COut{}, nil
}

// invoker
func GatherAndProcessServices(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	// A, B
	ab := newABProcessor()
	ab.start(ctx, data)
	inputC, err := ab.wait(ctx)
	if err != nil {
		return COut{}, err
	}

	// C
	c := newCProcessor()
	c.start(ctx, inputC)
	out, err := c.wait(ctx)
	return out, err
}

func TestGatherAndProcessServices(t *testing.T) {
	input := Input{
		A: "hello",
		B: "golang",
	}
	out, err := GatherAndProcessServices(context.Background(), input)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", out)
}
