package hosts

import (
	"os"
	"bufio"
	"io"
	"slices"
	"strings"
) 

type FocusStatus string
type Manager struct {
	hostsFile 	*os.File
	Status 		FocusStatus
	Domains 	[]string
}

const (
	ipAddress 				   = "127.0.0.1"
	FocusStatusOn FocusStatus  = "on"
	FocusStatusOff FocusStatus = "off"
	CommentStart			   = "#focusgopher:start"
	CommentEnd			  	   = "#focusgopher:end"
	CommentStatusOn			   = "#focusgopher:on"
	CommentStatusOff		   = "#focusgopher:off"
)

func (h *Manager) Init() error {
	var err error
	h.hostsFile, err = os.OpenFile("/etc/hosts", os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (h *Manager) Close() error {
	if h.hostsFile != nil {
		return h.hostsFile.Close()
	}

	return nil
}