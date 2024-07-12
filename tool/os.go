package tool

import (
	"os"
	"os/exec"
)

func CreateAndWriteFile(path string, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		defer func() {
			err := file.Close()
			if err != nil {
				return
			}
		}()
		if err != nil {
			return err
		}
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func FmtGoCode(path string) error {
	c := exec.Command("go", "fmt", path)
	return c.Run()
}
