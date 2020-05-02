package main

import (
	"github.com/owenkao/go_job_queue/pkg/owen"
	"github.com/owenkao/go_job_queue/pkg/singleton"
)

func main() {
	owen.Init()
	singleton.GoInstance()
}
