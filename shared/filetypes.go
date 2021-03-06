package shared

import (
	"os/exec"
	"strings"
)

const (
	// File types
	fileTypeUNDEFINED = iota
	fileTypeELFEXECUTABLE
	fileTypeELFOBJECT
	fileTypeELFSHARED
	fileTypeMACHEXECUTABLE
	fileTypeMACHOBJECT
	fileTypeMACHSHARED
	fileTypeARCHIVE
)

func getFileType(realPath string) (fileType int) {
	// We need the file command to guess the file type
	cmd := exec.Command("file", realPath)
	out, err := cmd.Output()
	if err != nil {
		LogFatal("There was an error getting the type of %s. Make sure that the 'file' command is installed.", realPath)
	}

	// Test the output
	if fo := string(out); strings.Contains(fo, "ELF") && strings.Contains(fo, "executable") {
		fileType = fileTypeELFEXECUTABLE
	} else if strings.Contains(fo, "Mach-O") && strings.Contains(fo, "executable") {
		fileType = fileTypeMACHEXECUTABLE
	} else if strings.Contains(fo, "ELF") && strings.Contains(fo, "shared") {
		fileType = fileTypeELFSHARED
	} else if strings.Contains(fo, "Mach-O") && strings.Contains(fo, "dynamically linked shared") {
		fileType = fileTypeMACHSHARED
	} else if strings.Contains(fo, "current ar archive") {
		fileType = fileTypeARCHIVE
	} else if strings.Contains(fo, "ELF") && strings.Contains(fo, "relocatable") {
		fileType = fileTypeELFOBJECT
	} else if strings.Contains(fo, "Mach-O") && strings.Contains(fo, "object") {
		fileType = fileTypeMACHOBJECT
	} else {
		fileType = fileTypeUNDEFINED
	}

	return
}
