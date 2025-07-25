// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips || mipsle)

package linux

const (
	SYS_CLOSE         = 4006
	SYS_FCNTL         = 4055
	SYS_MPROTECT      = 4125
	SYS_PRCTL         = 4192
	SYS_EPOLL_CTL     = 4249
	SYS_EPOLL_PWAIT   = 4313
	SYS_EPOLL_CREATE1 = 4326
	SYS_EPOLL_PWAIT2  = 4441
	SYS_EVENTFD2      = 4325
	SYS_OPENAT        = 4288
	SYS_PREAD64       = 4200
	SYS_READ          = 4003

	EFD_NONBLOCK = 0x80

	O_LARGEFILE = 0x2000
)

type EpollEvent struct {
	Events    uint32
	pad_cgo_0 [4]byte
	Data      uint64
}
