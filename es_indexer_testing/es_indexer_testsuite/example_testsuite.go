/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package es_indexer_testsuite

import (
	"github.com/kurtosis-tech/kurtosis-go/lib/testsuite"
)

type ExampleTestsuite struct {
	serviceImage string
}

func NewExampleTestsuite(serviceImage string) *ExampleTestsuite {
	return &ExampleTestsuite{serviceImage: serviceImage}
}


func (suite ExampleTestsuite) GetTests() map[string]testsuite.Test {
	return map[string]testsuite.Test{
		"fixedSizeExampleTest": FixedSizeExampleTest{ServiceImage: suite.serviceImage},
	}
}

func (suite ExampleTestsuite) GetNetworkWidthBits() uint32 {
	return 8
}


