// Copyright (C) 2016 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netutils

import (
	"net"
	"net/netip"
	"strings"
	"syscall"
)

func extractFamilyFromAddress(address string) (int, error) {
	ip, err := netip.ParseAddr(address)
	if err != nil {
		return 0, err
	}
	if ip.Is6() {
		return syscall.AF_INET6, nil
	}
	return syscall.AF_INET, nil
}

func extractFamilyFromConn(conn net.Conn) int {
	family := syscall.AF_INET
	if strings.Contains(conn.RemoteAddr().String(), "[") {
		family = syscall.AF_INET6
	}
	return family
}

func extractProtoFromAddress(address string) (string, error) {
	family, err := extractFamilyFromAddress(address)
	if err != nil {
		return "", err
	}
	if family == syscall.AF_INET6 {
		return "tcp6", nil
	}
	return "tcp4", nil
}
