package clue

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"runtime"
)

//UseValue is a struct that includes VarMap as a map[string]string to make specific values persistent using a cache file
type UseValue struct {
	VarMap map[string]string
}

//GetValue is a struct that includes references to string values as map[string]*string
type GetValue struct {
	VarMap map[string]*string
}

//DeleteGobFile deletes a specific cache file stored as suffix
func DeleteGobFile(suffix string) (err error) {
	fileLocation := fmt.Sprintf("%v/%v-%v.gob", os.TempDir(), suffix, os.Getppid())
	err = os.Remove(fileLocation)
	if err != nil {
		return fmt.Errorf("Problem removing file: %v", err)
	}
	return nil
}

//EncodeGobFile encodes a Go-Binary file that is made of a UseValue type with a map.
func EncodeGobFile(suffix string, useValue UseValue) (err error) {
	fileLocation := fmt.Sprintf("%v/%v-%v.gob", os.TempDir(), suffix, os.Getppid())
	if os.Getenv("CLUE_DEBUG") == "true" {
		fmt.Println("fileLocaton: ", fileLocation)
		fmt.Printf("useValue: %+v\n", useValue)
	}
	fmt.Println(fileLocation)
	file, err := os.Create(fileLocation)
	if err != nil {
		return fmt.Errorf("Problem creating file: %v", err)
	}

	if runtime.GOOS != "windows" {
		if err = file.Chmod(0600); err != nil {
			return fmt.Errorf("Problem setting persmission onfile: %v", err)
		}
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Problem closing file:", err)
		}
	}()

	fileWriter := bufio.NewWriter(file)

	encoder := gob.NewEncoder(fileWriter)
	err = encoder.Encode(useValue)
	//fmt.Println(useValue)
	if err != nil {
		return fmt.Errorf("Problem encoding gob: %v", err)
	}
	fileWriter.Flush()
	return
}

//DecodeGobFile adds imports Go-Binary contents that was set previously to the GetValue type with a map and references to strings
func DecodeGobFile(suffix string, getValue *GetValue) (err error) {
	fileLocation := fmt.Sprintf("%v/%v-%v.gob", os.TempDir(), suffix, os.Getppid())
	if os.Getenv("CLUE_DEBUG") == "true" {
		fmt.Println("fileLocation: " + fileLocation)
	}
	file, err := os.Open(fileLocation)
	if err != nil {
		if os.IsExist(err) {
			log.Fatal("Problem opening file:", err)
		} else {
			return nil
		}
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Problem closing file:", err)
		}
	}()

	fileReader := bufio.NewReader(file)

	decoder := gob.NewDecoder(fileReader)
	err = decoder.Decode(&getValue)
	if err != nil {
		return fmt.Errorf("Problem decoding file: %v", err)
	}
	if os.Getenv("CLUE_DEBUG") == "true" {
		fmt.Printf("getValue: %+v\n", getValue)
	}
	return
}
