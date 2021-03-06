/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 *
 */

package zk

// pingXID represents the XID which is used in ping/keepalive packet headers.
const pingXID = -2

// Below constants represent codes used by Zookeeper to differentiate requests.
// https://zookeeper.apache.org/doc/r3.4.8/api/constant-values.html#org.apache.zookeeper.ZooDefs.OpCode.getData
const (
	opGetData     = 4
	opGetChildren = 8
	opPing        = 11
)
