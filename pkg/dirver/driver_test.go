package dirver

import (
	"fmt"
	"testing"
)

func TestInspect(t *testing.T) {
	d := NewDriver("sealos")
	info, err := d.Inspect("docker.io/labring/nginx:1.21.6")
	fmt.Println(info, err)
}

func TestPull(t *testing.T) {
	d := NewDriver("sealos")
	info, err := d.Pull("docker.io/labring/nginx:1.21.6")
	fmt.Println(info, err)
}

func TestLogin(t *testing.T) {
	d := NewDriver("sealos")
	info, err := d.Login("docker.io", "lingdie", "")
	fmt.Println(info, err)
}
