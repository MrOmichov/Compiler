package main

type WORD [2]byte
type DWORD [4]byte

type DOS_HEADER struct {
	e_magic    WORD     // Magic number
	e_cblp     WORD     // Bytes on last page of file
	e_cp       WORD     // Pages in file
	e_crlc     WORD     // Relocations
	e_cparhdr  WORD     // Size of header in paragraphs
	e_minalloc WORD     // Minimum extra paragraphs needed
	e_maxalloc WORD     // Maximum extra paragraphs needed
	e_ss       WORD     // Initial (relative) SS value
	e_sp       WORD     // Initial SP value
	e_csum     WORD     // Checksum
	e_ip       WORD     // Initial IP value
	e_cs       WORD     // Initial (relative) CS value
	e_lfarlc   WORD     // File address of relocation table
	e_ovno     WORD     // Overlay number
	e_res      [4]WORD  // Reserved words
	e_oemid    WORD     // OEM identifier (for e_oeminfo)
	e_oeminfo  WORD     // OEM information; e_oemid specific
	e_res2     [10]WORD // Reserved words
	e_lfanew   DWORD    // File address of new exe header
}

func getDOS_HEADER() *DOS_HEADER {
	dh := DOS_HEADER{}

	dh.e_magic = [2]byte{'M', 'Z'}

	return &dh
}
