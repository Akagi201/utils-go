package files

import (
	"os"
	"sync"
)

// SynchronizedFile wraps an os.File pointer with a mutex
type SynchronizedFile struct {
	file  *os.File
	mutex sync.Mutex
}

// NewSynchronizedFile synchronizes writing to a writer
func NewSynchronizedFile(f *os.File) *SynchronizedFile {
	sf := &SynchronizedFile{file: f}
	return sf
}

// WriteString writes to the file
func (sf *SynchronizedFile) WriteString(text string) (int, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.file.WriteString(text)
}

// Close closes the file
func (sf *SynchronizedFile) Close() error {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.file.Close()
}
