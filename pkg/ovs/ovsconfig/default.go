// Copyright 2020 Antrea Authors
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

//go:build !windows
// +build !windows

package ovsconfig

import (
	"path"
	"time"
)

const (
	DefaultOVSRunDir = "/var/run/openvswitch"
	//We use winpipe here because it is hardcoded as the socket value that is parsed in our openvswitch db library, github.com/TomCodeLV/OVSDB-golang-lib/pkg/ovsdb .
	//defaultConnNetwork = "unix"
	defaultConnNetwork = "unix"
	// Wait up to 5 seconds when getting port.
	defaultGetPortTimeout    = 5 * time.Second
	defaultOvsVersionMessage = "OVS version not found in ovsdb. Please configure your OVS (ovsdb) to provide version information."
)

func GetConnAddress(ovsRunDir string) string {
	return path.Join(ovsRunDir, defaultOVSDBFile)
}
