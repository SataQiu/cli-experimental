// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wiretest

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/cli-experimental/internal/pkg/apply"
	"sigs.k8s.io/cli-experimental/internal/pkg/clik8s"
	delete2 "sigs.k8s.io/cli-experimental/internal/pkg/delete"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/dispatch"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/list"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/output"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/parse"
	"sigs.k8s.io/cli-experimental/internal/pkg/prune"
	"sigs.k8s.io/cli-experimental/internal/pkg/resourceconfig"
	"sigs.k8s.io/cli-experimental/internal/pkg/status"
	"sigs.k8s.io/cli-experimental/internal/pkg/wirecli/wireconfig"
	"sigs.k8s.io/cli-experimental/internal/pkg/wirecli/wirek8s"
)

// Injectors from wire.go:

func InitializeStatus(resourceConfigs clik8s.ResourceConfigs, commit *object.Commit, writer io.Writer) (*status.Status, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	clientset, err := wirek8s.NewKubernetesClientSet(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	statusStatus := &status.Status{
		Resources: resourceConfigs,
		Out:       writer,
		Clientset: clientset,
		Commit:    commit,
	}
	return statusStatus, func() {
		cleanup()
	}, nil
}

func InitializeApply(resourceConfigs clik8s.ResourceConfigs, commit *object.Commit, writer io.Writer) (*apply.Apply, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	restMapper, err := wirek8s.NewRestMapper(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, err := wirek8s.NewClient(dynamicInterface, restMapper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	applyApply := &apply.Apply{
		DynamicClient: client,
		Out:           writer,
		Resources:     resourceConfigs,
		Commit:        commit,
	}
	return applyApply, func() {
		cleanup()
	}, nil
}

func InitializeCommandBuilder(writer io.Writer) (*dy.CommandBuilder, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	clientset, err := wirek8s.NewKubernetesClientSet(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	clientsetClientset, err := wirek8s.NewExtensionsClientSet(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	commandLister := &list.CommandLister{
		Client:        clientsetClientset,
		DynamicClient: dynamicInterface,
	}
	commandParser := &parse.CommandParser{}
	commandOutputWriter := &output.CommandOutputWriter{
		Output: writer,
	}
	dispatcher := &dispatch.Dispatcher{
		KubernetesClient: clientset,
		DynamicClient:    dynamicInterface,
		Writer:           commandOutputWriter,
	}
	commandBuilder := &dy.CommandBuilder{
		KubernetesClient: clientset,
		Lister:           commandLister,
		Parser:           commandParser,
		Dispatcher:       dispatcher,
		Writer:           commandOutputWriter,
	}
	return commandBuilder, func() {
		cleanup()
	}, nil
}

func InitializeDispatcher(writer io.Writer) (*dispatch.Dispatcher, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	clientset, err := wirek8s.NewKubernetesClientSet(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	commandOutputWriter := &output.CommandOutputWriter{
		Output: writer,
	}
	dispatcher := &dispatch.Dispatcher{
		KubernetesClient: clientset,
		DynamicClient:    dynamicInterface,
		Writer:           commandOutputWriter,
	}
	return dispatcher, func() {
		cleanup()
	}, nil
}

func InitializeLister(writer io.Writer) (*list.CommandLister, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	clientset, err := wirek8s.NewExtensionsClientSet(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	commandLister := &list.CommandLister{
		Client:        clientset,
		DynamicClient: dynamicInterface,
	}
	return commandLister, func() {
		cleanup()
	}, nil
}

func InitializeDelete(resourceConfigs clik8s.ResourceConfigs, commit *object.Commit, writer io.Writer) (*delete2.Delete, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	restMapper, err := wirek8s.NewRestMapper(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, err := wirek8s.NewClient(dynamicInterface, restMapper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	deleteDelete := &delete2.Delete{
		DynamicClient: client,
		Out:           writer,
		Resources:     resourceConfigs,
		Commit:        commit,
	}
	return deleteDelete, func() {
		cleanup()
	}, nil
}

func InitializePrune(resourcePruneConfigs clik8s.ResourcePruneConfigs, commit *object.Commit, writer io.Writer) (*prune.Prune, func(), error) {
	config, cleanup, err := NewRestConfig()
	if err != nil {
		return nil, nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	restMapper, err := wirek8s.NewRestMapper(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, err := wirek8s.NewClient(dynamicInterface, restMapper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	prunePrune := &prune.Prune{
		DynamicClient: client,
		Out:           writer,
		Resources:     resourcePruneConfigs,
		Commit:        commit,
	}
	return prunePrune, func() {
		cleanup()
	}, nil
}

func InitializConfigProvider() resourceconfig.ConfigProvider {
	pluginConfig := wireconfig.NewPluginConfig()
	factory := wireconfig.NewResMapFactory(pluginConfig)
	fileSystem := wireconfig.NewFileSystem()
	transformerFactory := wireconfig.NewTransformerFactory()
	kustomizeProvider := wireconfig.NewKustomizeProvider(factory, fileSystem, transformerFactory, pluginConfig)
	return kustomizeProvider
}

func InitializeRawConfigProvider() resourceconfig.ConfigProvider {
	rawConfigFileProvider := &resourceconfig.RawConfigFileProvider{}
	return rawConfigFileProvider
}

// wire.go:

func InitializeKustomization() ([]string, func(), error) {
	f1, err := ioutil.TempDir("/tmp", "TestApply")
	if err != nil {
		return nil, nil, err
	}
	err = ioutil.WriteFile(filepath.Join(f1, "kustomization.yaml"), []byte(`apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: test-map

inventory:
  type: ConfigMap
  configMap:
    name: inventory
    namespace: default

namespace: default
`), 0644)
	if err != nil {
		return nil, nil, err
	}

	f2, err := ioutil.TempDir("/tmp", "TestApply")
	if err != nil {
		return nil, nil, err
	}
	err = ioutil.WriteFile(filepath.Join(f2, "kustomization.yaml"), []byte(`apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
configMapGenerator:
- name: test-map
  literals:
  - foo=bar

inventory:
  type: ConfigMap
  configMap:
    name: inventory
    namespace: default

namespace: default
`), 0644)
	if err != nil {
		return nil, nil, err
	}

	return []string{f1, f2}, func() {
		os.RemoveAll(f1)
		os.RemoveAll(f2)
	}, nil
}
