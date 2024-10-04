// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

//go:build !windows
// +build !windows

//lint:file-ignore Unconvert (redundant conversions are necessary for cross-platform compatibility)

package sysutil

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

// ProcessIdentity returns a string describing the user and group that this
// process is running as.
func ProcessIdentity() string {
	return fmt.Sprintf("uid %d euid %d gid %d egid %d",
		unix.Getuid(), unix.Geteuid(), unix.Getgid(), unix.Getegid())
}

// IsCrossDeviceLinkErrno checks whether the given error object (as
// extracted from an *os.LinkError) is a cross-device link/rename
// error.
func IsCrossDeviceLinkErrno(errno error) bool {
	return errno == syscall.EXDEV
}

// TerminateSelf sends SIGTERM to the process itself.
func TerminateSelf() error {
	pr, err := os.FindProcess(os.Getpid())
	if err != nil {
		// No-op.
		return nil //nolint:returnerrcheck
	}
	return pr.Signal(unix.SIGTERM)
}
