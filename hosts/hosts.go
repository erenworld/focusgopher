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
	hostsFile 		*os.File
	Status 			FocusStatus
	Domains 		[]string
}

const (
	ipAddress 				   = "127.0.0.1"
	FocusStatusOn  FocusStatus = "on"
	FocusStatusOff FocusStatus = "off"
	CommentStart			   = "#focusgopher:start"
	CommentEnd			  	   = "#focusgopher:end"
	CommentStatusOn			   = "#focusgopher:on"
	CommentStatusOff		   = "#focusgopher:off"
)

func (h *Manager) Init() error {
	// var err error
	// h.hostsFile, err = os.OpenFile("/etc/hosts", os.O_RDWR, 0600)
	h.Status = FocusStatusOff
	f, err := os.OpenFile(hostsPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var extractErr error
	h.Domains, h.Status, extractErr = h.ExtractDomains(string(data))
	if extractErr != nil {
		return extractErr
	}

	return nil
}

func (h *Manager) Close() error {
	if h.hostsFile != nil {
		return h.hostsFile.Close()
	}

	return nil
}

func (h *Manager) ExtractDomains(data string) ([]string, FocusStatus, error) {
	domains := []string{}
	inComment := false
	status := FocusStatusOff

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		trimLine := strings.TrimSpace(line)
		if trimLine == CommentStart {
			inComment = true
			continue
		} else if trimLine == CommentEnd {
			inComment = false
			break
		}

		if inComment {
			if trimLine == CommentStatusOn {
				status = FocusStatusOn
				continue
			}
			if trimLine == CommentStatusOff {
				status = FocusStatusOff
				continue
			}

			uncommentedLine := strings.Replace(trimLine, "#", "", 1)
			fields := strings.Fields(uncommentedLine) 
			if len(fields) > 1 {
				if !slices.Contains(domains, fields[1]) {
					domains = append(domains, fields[1])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return domains, status, err
	}

	return domains, status, nil
}