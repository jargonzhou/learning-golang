// The or-channel. 'Concurrency in Go' P.94
package example_cig

// OrDoneChannels combine done channels into a single done channel that closes if any of the channels close.
func OrDoneChannels(channels ...<-chan any) <-chan any {
	// termination cases: 0, 1 channel
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	// create channel and goroutine
	orDone := make(chan any)
	go func() {
		// close when this goroutine exit
		defer close(orDone)

		switch len(channels) {
		case 2: // 2 channels
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 3+ channels
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-OrDoneChannels(append(channels[3:], orDone)...): // tail call: append last
			}
		}
	}()
	return orDone
}
