// Copyright 2013 Tumblr, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "circuit/load"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"vena"
	"vena/server"
)

var (
	flagConfig = flag.String("config", "", "vena database configuration file name")
)

func main() {
	flag.Parse()

	if *flagConfig == "" {
		flag.Usage()
		os.Exit(1)
	}

	raw, err := ioutil.ReadFile(*flagConfig)
	if err != nil {
		panic(err)
	}
	config := &vena.Config{}
	if err = json.Unmarshal(raw, config); err != nil {
		panic(err)
	}

	println("Starting VENA shard servers")
	server.Spawn(config)
	println("VENA servers started successfully.")
}