package hypervctl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/containers/podman/v4/pkg/strongunits"
)

// ResizeDisk takes a diskPath and strongly typed new size and uses powershell
// to change its size.  There is no error protection for trying to size a disk
// smaller than the current size.
func ResizeDisk(diskPath string, newSize strongunits.GiB) error {
	resize := exec.Command("powershell", []string{"-command", fmt.Sprintf("Resize-VHD %s %d", diskPath, newSize.ToBytes())}...)
	resize.Stdout = os.Stdout
	resize.Stderr = os.Stderr
	if err := resize.Run(); err != nil {
		return fmt.Errorf("unable to resize disk image: %q", err)
	}
	return nil
}
