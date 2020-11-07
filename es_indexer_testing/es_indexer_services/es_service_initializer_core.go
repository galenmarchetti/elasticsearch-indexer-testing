/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package es_indexer_services

import (
	"github.com/kurtosis-tech/kurtosis-go/lib/services"
	"os"
)

const (
	elasticSearchServiceTestVolumeMountpoint = "/shared"
)

type ElasticSearchServiceInitializerCore struct{}

func (e ElasticSearchServiceInitializerCore) GetUsedPorts() map[int]bool {
	return map[int]bool{
		defaultElasticSearchHttpPort: true,
	}
}

func (e ElasticSearchServiceInitializerCore) GetServiceFromIp(ipAddr string) services.Service {
	return ElasticSearchImpl{IPAddr: ipAddr}
}

func (e ElasticSearchServiceInitializerCore) GetFilesToMount() map[string]bool {
	// TODO give an example of mounting files
	return map[string]bool{}
}

func (e ElasticSearchServiceInitializerCore) InitializeMountedFiles(mountedFiles map[string]*os.File, dependencies []services.Service) error {
	// TODO give example of mounting files
	return nil
}

func (e ElasticSearchServiceInitializerCore) GetTestVolumeMountpoint() string {
	return elasticSearchServiceTestVolumeMountpoint
}

func (e ElasticSearchServiceInitializerCore) GetStartCommand(mountedFileFilepaths map[string]string, ipPlaceholder string, dependencies []services.Service) ([]string, error) {
	// If there was a specific start command that we wanted Docker to run, we'd return the string array here. By
	//	returning nil, we tell Kurtosis to run the image with whatever CMD or ENTRYPOINT is specified in the
	//	Dockerfile will be used instead. This prevents the Kurtosis code from needing to know specifics about
	//	the Docker image (e.g. what filepath the binary to run is located at)
	return nil, nil
}