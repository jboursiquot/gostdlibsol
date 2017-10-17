package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// createFile()
	// deleteFile()
	// checkExistence()
	// renameFile()
	// copyFile()
	// writeToFile()
	// writeToFileWithIOUtil()
	// writeToFileWithBufferedWriter()
	// readFile()
	// readFileAgain()
	// readWithBufferedReader()
	// readWithScanner()
	// createDir()
	// createDirs()
	// deleteDir()
	dirTraversal()
}

func createFile() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	log.Printf("Created %s\n", f.Name())
}

func deleteFile() {
	os.Create("file.txt") // don't do this, handle your errors
	err := os.Remove("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Deleted file.txt")
}

func checkExistence() {
	fi, err := os.Stat("file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Does not exist.")
		}
	}
	log.Printf("Exists, last modified %v\n", fi.ModTime())
}

func renameFile() {
	f, _ := os.Create("file.txt") // don't do this, handle your errors
	err := os.Rename(f.Name(), "renamed.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func copyFile() {
	of, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer of.Close()

	nf, err := os.Create("copy.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	bw, err := io.Copy(nf, of)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Bytes written: %d\n", bw)

	if err := nf.Sync(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done")
}

func writeToFile() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte("Errors are values.\n")); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done.")
}

func writeToFileWithIOUtil() {
	b := []byte("Clear is better than clever.\n")
	if err := ioutil.WriteFile("myfile.txt", b, 0666); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done.")
}

func writeToFileWithBufferedWriter() {
	f, err := os.Create("panic.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	if _, err := bw.WriteString("Don't panic.\n"); err != nil {
		log.Println(err)
	}
	log.Printf("Buffered: %d\n", bw.Buffered())
	log.Printf("Available: %d\n", bw.Available())
	if err := bw.Flush(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done.")
}

func readFile() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func readFileAgain() {
	bs, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func readWithBufferedReader() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	bs, err := br.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(bs))

	bs, err = br.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(bs))
}

func readWithScanner() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	ln := 0
	for s.Scan() {
		ln++
		log.Printf("%d - %s", ln, s.Text())
	}
	if s.Err() != nil {
		log.Fatalln(err)
	}

	log.Println("Done.")
}

func createDir() {
	if err := os.Mkdir("mydir", 0744); err != nil {
		log.Fatalln(err)
	}
	fi, err := os.Stat("mydir")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("is 'dir' a directory?: %t", fi.IsDir())
}

func createDirs() {
	dirs := []string{"dir", "sub", "sub", "sub"}
	path := filepath.Join(dirs...)
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatalln(err)
	}
}

func deleteDir() {
	if err := os.Remove("mydir"); err != nil {
		log.Fatalln(err)
	}
	// if err := os.RemoveAll("dir"); err != nil {
	// 	log.Fatalln(err)
	// }
}

func dirTraversal() {
	file, err := os.Create("combined.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	bw := bufio.NewWriter(file)

	f := func(p string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		log.Println(p)

		bs, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}

		if _, err := bw.Write(bs); err != nil {
			return err
		}
		bw.Flush()

		return nil
	}

	if err := filepath.Walk("proverbs", f); err != nil {
		log.Fatalln(err)
	}
	log.Println("Done.")
}
