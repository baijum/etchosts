package hosts

import (
	"os"
	"path"
	"testing"

	"github.com/baijum/etchosts/pkg/hosts/mock"
	"github.com/golang/mock/gomock"
)

func TestAddEntry(t *testing.T) {
	dir := t.TempDir()
	hosts := path.Join(dir, "hosts")
	os.Create(hosts)
	ctrl := gomock.NewController(t)
	m := mock.NewMockhostsWriter(ctrl)
	hw = m
	m.EXPECT().OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, os.FileMode(0644)).Return(os.OpenFile(hosts, os.O_APPEND|os.O_WRONLY, 0644))
	err := AddEntry("127.0.0.1", "first.local", "second.local")
	if err != nil {
		t.Error("error should not occur:", err)
	}
	content, _ := os.ReadFile(hosts)
	expected := "127.0.0.1 first.local second.local\n"
	if string(content) != expected {
		t.Error("content not matching:", "expected:", expected, "real:", string(content))
	}
}
