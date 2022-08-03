package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Warn("test")
	log.Warn("test", "abc")

	abc := "11111"

	log.Warn("def", abc, "lll")
	log.Warn("def", abc+"lll")

	ccc := 3
	log.Warn("def", ccc, "lll")
}
