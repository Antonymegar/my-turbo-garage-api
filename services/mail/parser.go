package mail

import (
	"fmt"
	"strings"

	"myturbogarage/helpers"
)

// ParseTemplate ...
func ParseTemplate(template string, data map[string]any) (string, error) {
	base, err := helpers.ReadFile("templates/base.html")
	if err != nil {
		return "", err
	}

	content, err := helpers.ReadFile(template)
	if err != nil {
		return "", err
	}

	body := strings.ReplaceAll(string(base), "{{content}}", string(content))

	for k, v := range data {
		body = strings.ReplaceAll(body, fmt.Sprintf("{{%v}}", k), fmt.Sprintf("%v", v))
	}

	return body, nil
}
