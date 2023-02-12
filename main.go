package main

import (
	"flag"
	"time"

	"github.com/gogf/gf/os/glog"
	"github.com/labring-actions/sync/process"
)

func main() {
	var jsonPath, configPath, sealosPath string
	flag.StringVar(&jsonPath, "json-path", "digest.json", "digest mapper json file path")
	flag.StringVar(&configPath, "config", "config.yaml", "yaml file of images to sync")
	flag.StringVar(&sealosPath, "sealos-path", "sealos", "sealos path")
	p := process.NewProcesser(jsonPath, configPath, sealosPath)

	//c := make(chan os.Signal)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//go func() {
	//	for s := range c {
	//		switch s {
	//		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
	//			fmt.Println("Program Exit...", s)
	//			err := p.Exit()
	//			if err != nil {
	//				return
	//			}
	//		default:
	//			fmt.Println("other signal", s)
	//		}
	//	}
	//}()

	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := p.SaveDegist()
				if err != nil {
					return
				}
			}
		}
	}()

	err := p.Process()
	if err != nil {
		glog.Error(err)
		return
	}
}
