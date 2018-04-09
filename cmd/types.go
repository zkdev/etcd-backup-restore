// Copyright © 2018 The Gardener Authors.
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

package cmd

import (
	"github.com/sirupsen/logrus"
)

const (
	backupFormatVersion             = "v1"
	defaultServerPort               = 8080
	defaultName                     = "default"
	defaultInitialAdvertisePeerURLs = "http://localhost:2380"
)

var (
	logger = logrus.New()

	//snapshotter flags
	schedule              string
	etcdEndpoints         string
	maxBackups            int
	etcdConnectionTimeout int

	//server flags
	port int
	//restore flags
	restoreCluster      string
	restoreClusterToken string
	restoreDataDir      string
	restorePeerURLs     []string
	restoreName         string
	skipHashCheck       bool

	//snapstore flags
	storageProvider  string
	storageContainer string
	storagePrefix    string
)