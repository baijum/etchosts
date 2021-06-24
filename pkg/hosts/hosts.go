package hosts

import (
	"os"
	"strings"
)

//go:generate mockgen -package mock -source hosts.go -destination mock/mock_hosts.go
type hostsWriter interface {
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
}

type etcHostsWriter struct {
}

func (w *etcHostsWriter) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// AddEntry add new host entry to /etc/hosrs
func AddEntry(ip string, names ...string) error {
	fd, err := hw.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := fd.Write([]byte(ip + " " + strings.Join(names, " ") + "\n")); err != nil {
		return err
	}
	if err := fd.Close(); err != nil {
		return err
	}
	return nil
}

var hw hostsWriter

func init() {
	hw = &etcHostsWriter{}
}
