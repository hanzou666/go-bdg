package bdg

import (
	"bufio"
	"encoding/json"
	"os"
)

func (a *Alignment) WriteJson(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	alnJson, err := json.Marshal(a)
	if err != nil {
		return err
	}
	if _, err := writer.WriteString(string(alnJson)); err != nil {
		return err
	}
	if _, err := writer.WriteString("\n"); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
