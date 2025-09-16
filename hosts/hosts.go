package hosts

import "os"

// Manager provides read/write access to the hosts file (wrapper).
type Manager struct {
	hostsFile	*os.File
}

// Init opens /etc/hosts in read/write mode.
func (h *Manager) Init() error {
	var err error
	h.hostsFile, err = os.OpenFile("/etc/hosts", os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	return nil
}