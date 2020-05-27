package iterator

import "bufio"

type ScannerIterator struct {
	scanner *bufio.Scanner
}

func NewScannerIterator(scanner *bufio.Scanner) *ScannerIterator {
	return &ScannerIterator{scanner: scanner}
}

func (s *ScannerIterator) Value() interface{} {
	return s.scanner.Text()
}

func (s *ScannerIterator) Scan() bool {
	return s.scanner.Scan()
}
