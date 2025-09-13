package kvs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
	"spike.com/key-value-store/config"
)

type FileTransactionLogger struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File

	*config.AppConfig
}

func newFileTransactionLogger(appConfig *config.AppConfig) (TransactionLogger, error) {
	var filename = appConfig.Transact.File.Name
	appConfig.Log.Debug("Open file", zap.String("filename", filename))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	return &FileTransactionLogger{file: file, AppConfig: appConfig}, nil
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	l.Log.Debug("WritePut", zap.String("key", key), zap.String("value", value))
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.Log.Debug("WriteDelete", zap.String("key", key))
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}

func (l *FileTransactionLogger) Run() {
	l.Log.Debug("FileTransactionLogger start to run")

	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.Log.Debug("Receive event", zap.Any("event", e))
			l.lastSequence++

			e.Sequence = l.lastSequence // temporary for json marshal
			data, err := json.Marshal(e)
			if err != nil {
				errors <- err
				return
			}
			_, err = fmt.Fprintln(l.file, string(data))
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

func (l *FileTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	l.Log.Debug("Read events")

	scanner := bufio.NewScanner(l.file)
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		var e Event

		defer close(outEvent)
		defer close(outError)

		for scanner.Scan() {
			err := json.Unmarshal(scanner.Bytes(), &e)
			if err != nil {
				l.Log.Error("input parse error", zap.Error(err))
				outError <- fmt.Errorf("input parse error: %w", err)
				return
			}

			if l.lastSequence >= e.Sequence {
				l.Log.Error("transaction numbers out of sequence",
					zap.Uint64("lastSequence", l.lastSequence),
					zap.Uint64("sequence", e.Sequence))
				outError <- fmt.Errorf("transaction numbers out of sequence")
				return
			}

			l.lastSequence = e.Sequence
			outEvent <- e
		}

		if err := scanner.Err(); err != nil {
			l.Log.Error("transaction log read failure", zap.Error(err))
			outError <- fmt.Errorf("transaction log read failure: %w", err)
			return
		}

	}()

	return outEvent, outError
}
