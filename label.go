package print

import (
	"errors"
	"io"
	"os/exec"
	"strings"
)

func CenterLabelText(text string) string {
	if text == "" {
		return ""
	}

	const (
		cols = 17
		rows = 5
	)

	lines := strings.Split(text, "\n")
	if len(lines) > rows {
		lines = lines[:rows]
	}
	centered := strings.Repeat("\n", (rows-len(lines)+1)/2)
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		runes := []rune(lines[i])
		if len(runes) > cols {
			runes = runes[:cols]
		}
		lines[i] = string(runes)
		lines[i] = strings.Repeat(" ", (cols-len(runes)+1)/2) + lines[i]
		centered += lines[i] + "\n"
	}
	centered = strings.TrimSuffix(centered, "\n")
	return centered
}

func LabelCentered(text string) error {
	text = CenterLabelText(text)
	lp := exec.Command("lp")
	pipe, _ := lp.StdinPipe()
	io.Copy(pipe, strings.NewReader(text))
	pipe.Close()
	output, err := lp.CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + ": " + string(output))
	}
	return nil
}
