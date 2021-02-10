package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	filename := ".env"
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		current_mode_re := `^APP_MODE=(\w+)$`
		new_mode_re := `^#\sAPP_MODE=(\w+)$`

		current_matched, err := regexp.MatchString(current_mode_re, line)
		new_matched, err2 := regexp.MatchString(new_mode_re, line)

		if err != nil || err2 != nil {
			log.Fatalln(err)
		}

		if current_matched {
			current_mode := regexp.MustCompile(current_mode_re)
			lines[i] = current_mode.ReplaceAllString(line, "# APP_MODE=${1}")
		}

		if new_matched {
			new_mode := regexp.MustCompile(new_mode_re)
			lines[i] = new_mode.ReplaceAllString(line, "APP_MODE=${1}")
			fmt.Println(new_mode.ReplaceAllString(line, "Switched to ${1} mode."))
		}

	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(filename, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}
}
