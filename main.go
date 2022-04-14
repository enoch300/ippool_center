package main

import (
	"ippool_center/cmd"
	"ippool_center/utils/log"
)

func init() {
	log.NewLogger(3)
}

func main() {
	cmd.Execute()
}
