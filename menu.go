// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

// Constants for MENUITEMINFO.fMask
const (
	MIIM_STATE      = 1
	MIIM_ID         = 2
	MIIM_SUBMENU    = 4
	MIIM_CHECKMARKS = 8
	MIIM_TYPE       = 16
	MIIM_DATA       = 32
	MIIM_STRING     = 64
	MIIM_BITMAP     = 128
	MIIM_FTYPE      = 256
)

// Constants for MENUITEMINFO.fType
const (
	MFT_STRING       = MF_STRING
	MFT_BITMAP       = MF_BITMAP
	MFT_MENUBARBREAK = MF_MENUBARBREAK
	MFT_MENUBREAK    = MF_MENUBREAK
	MFT_OWNERDRAW    = MF_OWNERDRAW
	MFT_RADIOCHECK   = 0x00000200
	MFT_SEPARATOR    = MF_SEPARATOR
	MFT_RIGHTORDER   = 0x00002000
	MFT_RIGHTJUSTIFY = MF_RIGHTJUSTIFY
)

// Constants for MENUITEMINFO.fState
const (
	MFS_GRAYED    = 0x00000003
	MFS_DISABLED  = MFS_GRAYED
	MFS_CHECKED   = MF_CHECKED
	MFS_HILITE    = MF_HILITE
	MFS_ENABLED   = MF_ENABLED
	MFS_UNCHECKED = MF_UNCHECKED
	MFS_UNHILITE  = MF_UNHILITE
	MFS_DEFAULT   = MF_DEFAULT
)

// Constants for MENUITEMINFO.hbmp*
const (
	HBMMENU_CALLBACK        = -1
	HBMMENU_SYSTEM          = 1
	HBMMENU_MBAR_RESTORE    = 2
	HBMMENU_MBAR_MINIMIZE   = 3
	HBMMENU_MBAR_CLOSE      = 5
	HBMMENU_MBAR_CLOSE_D    = 6
	HBMMENU_MBAR_MINIMIZE_D = 7
	HBMMENU_POPUP_CLOSE     = 8
	HBMMENU_POPUP_RESTORE   = 9
	HBMMENU_POPUP_MAXIMIZE  = 10
	HBMMENU_POPUP_MINIMIZE  = 11
)

// MENUINFO mask constants
const (
	MIM_APPLYTOSUBMENUS = 0x80000000
	MIM_BACKGROUND      = 0x00000002
	MIM_HELPID          = 0x00000004
	MIM_MAXHEIGHT       = 0x00000001
	MIM_MENUDATA        = 0x00000008
	MIM_STYLE           = 0x00000010
)

// MENUINFO style constants
const (
	MNS_AUTODISMISS = 0x10000000
	MNS_CHECKORBMP  = 0x04000000
	MNS_DRAGDROP    = 0x20000000
	MNS_MODELESS    = 0x40000000
	MNS_NOCHECK     = 0x80000000
	MNS_NOTIFYBYPOS = 0x08000000
)

const (
	// Menu flags for Add/Check/EnableMenuItem()
	MF_INSERT = 0x00000000
	MF_CHANGE = 0x00000080
	MF_APPEND = 0x00000100
	MF_DELETE = 0x00000200
	MF_REMOVE = 0x00001000

	MF_BYCOMMAND  = 0x00000000
	MF_BYPOSITION = 0x00000400

	MF_SEPARATOR = 0x00000800

	MF_ENABLED  = 0x00000000
	MF_GRAYED   = 0x00000001
	MF_DISABLED = 0x00000002

	MF_UNCHECKED       = 0x00000000
	MF_CHECKED         = 0x00000008
	MF_USECHECKBITMAPS = 0x00000200

	MF_STRING    = 0x00000000
	MF_BITMAP    = 0x00000004
	MF_OWNERDRAW = 0x00000100

	MF_POPUP        = 0x00000010
	MF_MENUBARBREAK = 0x00000020
	MF_MENUBREAK    = 0x00000040

	MF_UNHILITE = 0x00000000
	MF_HILITE   = 0x00000080

	MF_DEFAULT      = 0x00001000
	MF_SYSMENU      = 0x00002000
	MF_HELP         = 0x00004000
	MF_RIGHTJUSTIFY = 0x00004000

	MF_MOUSESELECT = 0x00008000
)

type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      HMENU
	HbmpChecked   HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16
	Cch           uint32
	HbmpItem      HBITMAP
}

type MENUINFO struct {
	CbSize          uint32
	FMask           uint32
	DwStyle         uint32
	CyMax           uint32
	HbrBack         HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr
}
