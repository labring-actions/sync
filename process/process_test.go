package process

import (
	"fmt"
	"os"
	"testing"
)

func TestSyncImage(t *testing.T) {
	os.Create("config.yaml")
	p := NewProcesser("", "config.yaml", "sealos")
	err := p.SyncImage("docker.io/labring/nginx:1.21.6", "hub.sealos.cn/labring/nginx:1.21.6")
	if err != nil {
		fmt.Println(err)
		return
	}
}
