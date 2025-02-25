//go:build windows
// +build windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

/*
 * Source from https://github.com/WireGuard/wireguard-windows/tree/master/tunnel/winipcfg
 */

package iphlpapi

import (
	"golang.org/x/sys/windows"
	"sync/atomic"
)

//
// Notifications-related functions
//

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-notifyipinterfacechange
//sys	notifyIPInterfaceChange(family uint16, callback uintptr, callerContext uintptr, initialNotification bool, notificationHandle *windows.Handle) (ret error) = iphlpapi.NotifyIpInterfaceChange

// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-cancelmibchangenotify2
//sys	cancelMibChangeNotify2(notificationHandle windows.Handle) (ret error) = iphlpapi.CancelMibChangeNotify2

// https://learn.microsoft.com/en-us/windows/win32/api/netioapi/nf-netioapi-convertinterfaceluidtoguid
//sys   convertInterfaceLUIDToGUID(interfaceLUID *LUID, interfaceGUID *windows.GUID) (ret error) = iphlpapi.ConvertInterfaceLuidToGuid
//sys   convertInterfaceGUIDToLUID(interfaceGUID *windows.GUID, interfaceLUID *LUID) (ret error) = iphlpapi.ConvertInterfaceGuidToLuid
//sys   convertInterfaceIndexToLUID(interfaceIndex uint32, interfaceLUID *LUID) (ret error) = iphlpapi.ConvertInterfaceIndexToLuid

var notifyIpInterfaceChangeSupported atomic.Int32

func IsNotifyIPInterfaceChangeSupported() bool {
	v := notifyIpInterfaceChangeSupported.Load()
	if v == 1 {
		return true
	} else if v == 2 {
		return false
	}
	err := procNotifyIpInterfaceChange.Find()
	if err == nil {
		notifyIpInterfaceChangeSupported.Store(1)
	} else {
		notifyIpInterfaceChangeSupported.Store(2)
	}
	return err == nil
}

func NotifyIpInterfaceChange(family uint16, callback uintptr, callerContext uintptr, initialNotification bool, notificationHandle *windows.Handle) (ret error) {
	if err := procNotifyIpInterfaceChange.Find(); err != nil {
		return err
	}

	return notifyIPInterfaceChange(family, callback, callerContext, initialNotification, notificationHandle)
}
