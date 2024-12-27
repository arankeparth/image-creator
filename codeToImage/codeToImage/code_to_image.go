package codeToImage

import (
	"fmt"
	"os/exec"
	"strings"
)

func CodeToImage(isUpdateMode bool, code string, webUiUrl string) (string, error) {
	if isUpdateMode {
		cmd := fmt.Sprintf("./codeToImage/codeToImage/dist/app --isUpdate -g %s %s", code, webUiUrl)
		b, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			return "", err
		}
		return strings.Trim(strings.Trim(string(b), "\n"), "\r"), nil
	} else {
		cmd := fmt.Sprintf("./codeToImage/codeToImage/dist/app -g %s %s", code, webUiUrl)
		b, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			return "", err
		}
		return strings.Trim(strings.Trim(string(b), "\n"), "\r"), nil
	}
}
