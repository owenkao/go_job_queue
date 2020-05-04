package main

import (
	"github.com/owenkao/go_job_queue/internal/mymap"
	"github.com/owenkao/go_job_queue/pkg/owen"
	"github.com/owenkao/go_job_queue/pkg/pattern/singleton"
)

func main() {
	owen.Init()
	singleton.GetInstance()
	//jobqueue.Init()
	mymap.Init()
}
