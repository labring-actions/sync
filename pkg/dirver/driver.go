package dirver

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

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
	cmdC.WaitDelay = 30 * time.Second
	output, err := cmdC.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}

func (d *Driver) RetryDo(args []string) (string, error) {
	var retErr error
	for i := 0; i < 3; i++ {
		out, err := d.Do(args)
		if err != nil {
			glog.Warning("try to ", args, " :", i, "...")
			glog.Warning(out)
			retErr = err
			continue
		}
		return out, nil
	}
	return "", retErr
}

func (d *Driver) Inspect(image string) (string, error) {
	inspectInfo, err := d.RetryDo([]string{"inspect", fmt.Sprintf("docker://%s", image)})
	inspectInfo = strings.Replace(inspectInfo, " ", "", -1)
	inspectInfo = strings.Replace(inspectInfo, "\t", "", -1)
	inspectInfo = strings.Replace(inspectInfo, "\n", "", -1)
	inspectInfo = strings.Replace(inspectInfo, "\\", "", -1)
	return inspectInfo, err
}

func (d *Driver) Pull(image string) (string, error) {
	return d.RetryDo([]string{"pull", image})
}

func (d *Driver) Tag(id string, image string) (string, error) {
	return d.RetryDo([]string{"tag", id, image})
}

func (d *Driver) Push(image string) (string, error) {
	return d.RetryDo([]string{"push", image})
}

func (d *Driver) Login(registry, username, password string) (string, error) {
	return d.RetryDo([]string{"login", "-u", username, "-p", password, registry})
}

func (d *Driver) LoginK(registry, filePath string) (string, error) {
	return d.RetryDo([]string{"login", "-k", filePath, registry})
}
