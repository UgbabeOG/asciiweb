package renderer

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) ([]string, error) {
	paths := fmt.Sprintf("banners/%s.txt", path)
	data, err := os.ReadFile(paths)
	if err != nil {
		return nil, err
	}
	content := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return content, nil
}

func cleanInput(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	cleaned := s
	cleaned = strings.ReplaceAll(cleaned, "\\n", "\n")
	return cleaned

}

func GenerateASCII(text, banner string) (string, error) {
	if banner == "" {
		banner = "standard"
	}
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		return "", fmt.Errorf("invalid banner style selected. :%v. ", banner)
	}

	bannerStyle, err := readFile(banner)
	if err != nil {

		return "", err
	}
	input := strings.Split(cleanInput(text), "\n")
	var res strings.Builder
	for _, word := range input {
		if word == "" {
			res.WriteString("\n")
		}
		for i := range 8 {
			for _, v := range word {
				asciiVal := int(v)
				if asciiVal < 32 || asciiVal > 126 {

					return "", fmt.Errorf("%v is not a printable ASCII character", v)
				}
				offset := (asciiVal-32)*9 + 1
				if offset+8 < len(bannerStyle) {
					res.WriteString(bannerStyle[offset+i])
				}

			}
			res.WriteString("\n")
		}
	}
	return res.String(), nil
}
