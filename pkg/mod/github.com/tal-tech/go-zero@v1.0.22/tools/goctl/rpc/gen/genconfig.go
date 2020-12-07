package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/util"
)

const configTemplate = `package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
}
`

func (g *defaultRpcGenerator) genConfig() error {
	configPath := g.dirM[dirConfig]
	fileName := filepath.Join(configPath, fileConfig)
	if util.FileExists(fileName) {
		return nil
	}

	text, err := util.LoadTemplate(category, configTemplateFileFile, configTemplate)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
}
