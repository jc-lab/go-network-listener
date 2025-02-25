//go:build windows
// +build windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package iphlpapi

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output ziphlpapi_windows.go iphlpapi.go
