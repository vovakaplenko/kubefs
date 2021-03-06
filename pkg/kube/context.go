package kube

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

type Context struct {
	*Settings
	f.BaseDir

	config    *clientcmdapi.Config
	kubectl   dynamic.Interface
	discovery discovery.DiscoveryInterface

	resourceTypes map[string]f.Node
	ContextName   string
}
