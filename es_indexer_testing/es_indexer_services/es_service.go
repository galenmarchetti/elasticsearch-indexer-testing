/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package es_indexer_services

import (
	"github.com/kurtosis-tech/kurtosis-go/lib/services"
)

type ElasticSearchService interface {
	services.Service

	GetHttpSocket() Socket
}
