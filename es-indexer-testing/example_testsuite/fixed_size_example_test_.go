/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package example_testsuite

import (
	"fmt"
	"github.com/galenmarchetti/elasticsearch-indexer-testing/es-indexer-testing/example_networks/fixed_size_example_network"
	"github.com/kurtosis-tech/kurtosis-go/lib/networks"
	"github.com/kurtosis-tech/kurtosis-go/lib/testsuite"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	numNodes = 5
)

type FixedSizeExampleTest struct {
	ServiceImage string
}

func (test FixedSizeExampleTest) Run(network networks.Network, context testsuite.TestContext) {
	// NOTE: We have to do this as the first line of every test because Go doesn't have generics
	castedNetwork := network.(fixed_size_example_network.FixedSizeExampleNetwork)

	for i := 0; i < castedNetwork.GetNumNodes(); i++ {
		logrus.Infof("Making query against node #%v...", i)
		service, err := castedNetwork.GetService(i)
		if err != nil {
			context.Fatal(stacktrace.Propagate(err, "An error occurred when getting the service interface for node #%v", i))
		}
		socket := service.GetHelloWorldSocket()
		serviceUrl := fmt.Sprintf("http://%v:%v", socket.IPAddr, socket.Port)
		if _, err := http.Get(serviceUrl); err != nil {
			context.Fatal(stacktrace.Propagate(err, "Received an error when calling the example service endpoint for node #%v", i))
		}
		logrus.Infof("Successfully queried node #%v", i)
	}
}

func (test FixedSizeExampleTest) GetNetworkLoader() (networks.NetworkLoader, error) {
	return fixed_size_example_network.NewFixedSizeExampleNetworkLoader(numNodes, test.ServiceImage), nil
}

func (test FixedSizeExampleTest) GetExecutionTimeout() time.Duration {
	return 30 * time.Second
}

func (test FixedSizeExampleTest) GetSetupBuffer() time.Duration {
	return 30 * time.Second
}

