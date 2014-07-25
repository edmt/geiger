package main

import (
	l4g "code.google.com/p/log4go"
	"os"
	"regexp"
)

func CountFile(path string) {
	stats, _ := os.Stat(path)
	l4g.Debug("Nombre: %s", stats.Name())
	l4g.Debug("Tamano: %d", stats.Size())

	expr := regexp.MustCompile("(.{36}).{13}$")
	matches := expr.FindStringSubmatch(stats.Name())
	
	var uuid string
	if len(matches) == 2 {
		uuid = matches[1]
	}
	l4g.Debug("UUID: %s", uuid)
}
