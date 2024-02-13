package iterator

import (
	"bufio"
	"os"
	"sync"
)

var mtx = sync.Mutex{}

type LineByLineIterator struct {
	path     string
	fojb     *os.File
	scanner  *bufio.Scanner
	hasMore  bool
	isClosed bool
}

func NewLineByLineIterator(path string) (*LineByLineIterator, error) {
	fobj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fobj)

	return &LineByLineIterator{
		path:    path,
		scanner: scanner,
	}, nil
}

func (d *LineByLineIterator) HasNext() bool {
	if d.isClosed {
		return false
	}
	if d.hasMore {
		return true
	}
	mtx.Lock()
	d.hasMore = d.scanner.Scan()
	mtx.Unlock()
	return d.hasMore
}

func (d *LineByLineIterator) GetNext() interface{} {
	if d.HasNext() {
		str := d.scanner.Bytes()
		d.hasMore = false
		return string(str)
	}
	return nil
}

func (d *LineByLineIterator) Close() error {
	return d.fojb.Close()
}
