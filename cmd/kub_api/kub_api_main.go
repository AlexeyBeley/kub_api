package main

import (
	"flag"
	"path/filepath"

	"github.com/AlexeyBeley/go_common/logger"
	"github.com/AlexeyBeley/kub_api"
	"k8s.io/client-go/util/homedir"
)

var lg = &(logger.Logger{})

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	kub_api.list(kubeconfig)
}
