package clipboard

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
	"runtime"
)

func GetLocal() (data []byte, err error) {
	out := bytes.NewBuffer(nil)
	if runtime.GOOS == "linux" {
		cmd := exec.Command("xclip", "-o", "-sel", "clipboard")
		cmd.Stdout = out
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		if err := cmd.Wait(); err != nil {
			return nil, err
		}
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("pbpaste")
		cmd.Stdout = out
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		if err := cmd.Wait(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.Errorf("unknown os: " + runtime.GOOS)
	}
	data = out.Bytes()
	return data, nil
}

func SetLocal(data []byte) error {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("xclip", "-i", "-sel", "clipboard")
		cmd.Stdin = bytes.NewReader(data)
		if err := cmd.Start(); err != nil {
			return err
		}
		return cmd.Wait()
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("pbcopy")
		cmd.Stdin = bytes.NewReader(data)
		if err := cmd.Start(); err != nil {
			return err
		}
		return cmd.Wait()
	}
	return errors.Errorf("unknown os: " + runtime.GOOS)
}
