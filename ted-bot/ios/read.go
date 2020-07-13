package ios

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type File struct {
	name string
	path string
}

func LoaderFile(fileName string, filePath string) []string {
	file := File{name: fileName, path: filePath}
	return file.read()
}

func (f *File) read() []string {

	localFile, err := os.Open(f.fullPath())
	if err != nil {
		log.Fatal(err)
		log.Print("File Not Found")
	}

	reader := bufio.NewReader(localFile)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}
	localFile.Close()
	return lines
}

func (f *File) fullPath() string {
	f.name = strings.ReplaceAll(f.name, "/", "")
	f.path = strings.TrimSuffix(f.path, "/")
	if len(f.path) > 0 && len(f.name) > 0 {
		return f.path + "/" + f.name
	} else if len(f.name) > 0 {
		return f.name
	} else {
		return ""
	}
}
