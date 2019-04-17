/*
   Copyright 2019 txn2
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"flag"
	"os"

	"github.com/txn2/micro"
	"github.com/txn2/provision"
)

var (
	elasticServerEnv = getEnv("ELASTIC_SERVER", "http://elasticsearch:9200")
	systemPrefixEnv  = getEnv("SYSTEM_PREFIX", "system_")
)

func main() {

	esServer := flag.String("esServer", elasticServerEnv, "Elasticsearch Server")
	systemPrefix := flag.String("systemPrefix", systemPrefixEnv, "Prefix for system indices.")

	serverCfg, _ := micro.NewServerCfg("Auth")
	server := micro.NewServer(serverCfg)

	// Provision API
	provApi, err := provision.NewApi(&provision.Config{
		Logger:        server.Logger,
		HttpClient:    server.Client,
		ElasticServer: *esServer,
		IdxPrefix:     *systemPrefix,
		Token:         server.Token,
	})
	if err != nil {
		server.Logger.Fatal("failure to instantiate the auth API: " + err.Error())
		os.Exit(1)
	}

	// Auth a user
	server.Router.POST("/user", provApi.AuthUserHandler)

	// run provisioning server
	server.Run()
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
