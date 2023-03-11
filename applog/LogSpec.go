package Monitor

import (
    "io"
    "log"
    "net"
    "os"
)


type Monitor struct {
    *log.Logger
}

// Write implements the io.Writer interface.
func (m *Monitor) 1Write(p []byte) (int, error) {
    return len(p), m.Output(2, string(p))
}