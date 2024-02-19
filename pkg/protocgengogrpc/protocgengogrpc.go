package protocgengogrpc

import (
	"strings"

	"github.com/cat3306/ginctl/pkg/goctl"
	"github.com/cat3306/ginctl/pkg/golang"
	"github.com/cat3306/ginctl/rpc/execx"
	"github.com/cat3306/ginctl/util/env"
)

const (
	Name = "protoc-gen-go-grpc"
	url  = "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
)

func Install(cacheDir string) (string, error) {
	return goctl.Install(cacheDir, Name, func(dest string) (string, error) {
		err := golang.Install(url)
		return dest, err
	})
}

func Exists() bool {
	_, err := env.LookUpProtocGenGoGrpc()
	return err == nil
}

// Version is used to get the version of the protoc-gen-go-grpc plugin.
func Version() (string, error) {
	path, err := env.LookUpProtocGenGoGrpc()
	if err != nil {
		return "", err
	}
	version, err := execx.Run(path+" --version", "")
	if err != nil {
		return "", err
	}
	fields := strings.Fields(version)
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", nil
}
