package ios

import "testing"

func TestFullPathWithPathAndFileName(t *testing.T) {

	file := File{name: StringFileName, path: StringPath}

	fullPath := file.fullPath()

	if fullPath != StringPath+StringFileName {
		t.Error("Expected full path name "+StringPath+StringFileName+" , but got ", fullPath)
	}
}

func TestFullPath(t *testing.T) {

	file := File{name: StringFileName, path: StringPath}

	fullPath := file.fullPath()

	if fullPath != StringPath+StringFileName {
		t.Error("Expected file name "+StringPath+StringFileName+" , but got ", fullPath)
	}
}

func TestRead(t *testing.T) {

	file := File{name: StringFileName, path: "../" + StringPath}

	lines := file.read()
	if lines[0] != "Oi" {
		t.Error("Expected first line equals Oi, but got ", lines[0])
	}
}
