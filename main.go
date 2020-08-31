package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
func dirTree(f io.Writer, path string, printFiles bool) error {

	return buildDirTree(f, path, printFiles, 0, []bool{})
}

func buildDirTree(f io.Writer, path string, printFiles bool, level int, history []bool) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if !printFiles {
		files = onlyDirs(files)
	}

	for index, file := range files {

		if level != 0 {
			for i := 0; i < level; i++ {
				p := history[i]
				if p {
					_, _ = f.Write([]byte("│"))
				}
				_, _ = f.Write([]byte("\t"))

			}
		}

		if len(files) == index+1 {
			_, _ = f.Write([]byte("└───"))
		} else {
			_, _ = f.Write([]byte("├───"))
		}

		_, _ = f.Write([]byte(file.Name()))
		if !file.IsDir() {
			size := strconv.Itoa(int(file.Size())) + "b"
			if file.Size() == 0 {
				size = "empty"
			}
			f.Write([]byte(" (" + size + ")"))
		}

		_, _ = f.Write([]byte("\n"))

		if file.IsDir() {
			_ = buildDirTree(f, path+string(os.PathSeparator)+file.Name(), printFiles, level+1, append(history, len(files) != index+1))
		}
	}

	return nil
}

func onlyDirs(files []os.FileInfo) (result []os.FileInfo) {
	for _, file := range files {
		if file.IsDir() {
			result = append(result, file)
		}
	}
	return
}
