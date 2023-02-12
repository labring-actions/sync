package process

import (
	"fmt"
	"sort"

	"github.com/gogf/gf/os/glog"
	"github.com/labring-actions/sync/pkg/config"
	"github.com/labring-actions/sync/pkg/dirver"
	"github.com/labring-actions/sync/pkg/util"
)

type Processer struct {
	Mapper util.Mapper
	Config config.Config
	Driver dirver.Driver
}

func NewProcesser(jsonPath, configPath, sealosPath string) Processer {
	return Processer{
		Mapper: util.NewMapper(jsonPath),
		Config: config.NewConfig(configPath),
		Driver: dirver.NewDriver(sealosPath),
	}
}

func (p *Processer) Process() error {
	if err := p.Login(); err != nil {
		return err
	}
	if err := p.Mapper.FromJsonFile(); err != nil {
		return err
	}

	repos := make([]string, 0, len(p.Config.Images))
	for k := range p.Mapper.Data {
		repos = append(repos, k)
	}
	sort.Strings(repos)

	for _, repo := range repos {
		for _, tag := range p.Config.Images[repo] {
			if err := p.ProcessOneImage(fmt.Sprintf("%s:%s", repo, tag)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Processer) Login() error {
	glog.Info("try to login")
	info, err := p.Driver.Login(p.Config.SrcRegistry, p.Config.SrcRegisrtyUsername, p.Config.SrcRegisrtyPassword)
	if err != nil {
		return err
	}
	glog.Info("login to src registry ", p.Config.SrcRegistry)
	glog.Info(info)

	info, err = p.Driver.LoginK(p.Config.DestRegistry, p.Config.DestRegistryKubeconfig)
	if err != nil {
		return err
	}
	glog.Info("login to dest registry", p.Config.DestRegistry)
	glog.Info(info)
	return nil
}

func (p *Processer) LoadMapper() error {
	err := p.Mapper.FromJsonFile()
	if err != nil {
		return err
	}
	return nil
}

func (p *Processer) ProcessOneImage(image string) error {
	srcImage := fmt.Sprintf("%s/%s", p.Config.SrcRegistry, image)
	destImage := fmt.Sprintf("%s/%s", p.Config.DestRegistry, image)
	needSync, err := p.Check(srcImage)
	if err != nil {
		return err
	}
	if !needSync {
		glog.Info("skip sync image ", image)
		return nil
	}
	if err := p.SyncImage(srcImage, destImage); err != nil {
		return err
	}
	if err := p.MapImage(srcImage); err != nil {
		return err
	}
	return nil
}

func (p *Processer) Check(image string) (bool, error) {
	glog.Info("check image", image)
	inspectInfo, err := p.Driver.Inspect(image)
	if err != nil {
		return false, err
	}
	imageInfo, ok := p.Mapper.Data[image]
	if !ok || imageInfo != inspectInfo {
		return true, nil
	}
	return false, nil
}

func (p *Processer) SyncImage(srcImage, destImage string) error {
	glog.Info("SyncImage ", srcImage, " to ", destImage)
	if _, err := p.Driver.Pull(srcImage); err != nil {
		return err
	}
	if _, err := p.Driver.Tag(srcImage, destImage); err != nil {
		return err
	}
	if _, err := p.Driver.Push(destImage); err != nil {
		return err
	}
	return nil
}

func (p *Processer) MapImage(image string) error {
	inspectInfo, err := p.Driver.Inspect(image)
	if err != nil {
		return err
	}
	p.Mapper.Data[image] = inspectInfo
	return nil
}

func (p *Processer) SaveDegist() error {
	if err := p.Mapper.ToJsonFile(); err != nil {
		return err
	}
	return nil
}

func (p *Processer) Exit() error {
	err := p.SaveDegist()
	if err != nil {
		return err
	}
	return nil
}
