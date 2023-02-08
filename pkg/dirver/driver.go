package dirver

import (
	"os/exec"

	"github.com/gogf/gf/os/glog"
)

type Driver struct {
	SealosPath string
}

func NewDriver(sealosPath string) Driver {
	return Driver{SealosPath: sealosPath}
}

func (d *Driver) Do(args []string) (string, error) {
	glog.Info(args)
	cmdC := exec.Command(d.SealosPath, args...)
	output, err := cmdC.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}

func (d *Driver) Inspect(image string) (string, error) {
	return d.Do([]string{"inspect", image})
}

func (d *Driver) Pull(image string) (string, error) {
	return d.Do([]string{"pull", image})

}

func (d *Driver) Tag(id string, image string) (string, error) {
	return d.Do([]string{"tag", id, image})
}

func (d *Driver) Push(image string) (string, error) {
	return d.Do([]string{"push", image})
}

func (d *Driver) Login(registry, username, password string) (string, error) {
	return d.Do([]string{"login", "-u", username, "-p", password, registry})
}

func (d *Driver) LoginK(registry, filePath string) (string, error) {
	return d.Do([]string{"login", "-k", filePath, registry})
}
