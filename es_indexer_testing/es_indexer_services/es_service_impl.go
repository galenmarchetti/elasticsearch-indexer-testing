/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package es_indexer_services

const (
	defaultElasticSearchHttpPort = 9200
)
type ElasticSearchImpl struct{
	IPAddr string
}

func (e ElasticSearchImpl) GetHttpSocket() Socket {
	return Socket{
		IPAddr: e.IPAddr,
		Port: defaultElasticSearchHttpPort,
	}
}
