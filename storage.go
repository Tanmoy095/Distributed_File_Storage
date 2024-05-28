package main

import (
	"io"
	"log"
	"os"
)

type pathTransformFunc func(string) string

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type StoreOpts struct {
	pathTransformFunc pathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}

}

func (s *Store) writestream(key string, r io.Reader) error {
	//The purpose of this method is to generate a directory path based on the key.
	//Transforming the Path
	pathName := s.pathTransformFunc(key)
	// This line ensures that all directories in the path pathName are created.
	//os.MkdirAll will create the entire directory tree specified by pathName
	//Creating Directories........
	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}
	filename := "somefilename"
	//This line constructs the full path for the file by concatenating the pathName and filename with a / separator.
	pathAndFilename := pathName + "/" + filename
	//Opening the File
	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}
	//Copying Data to the File
	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written (%d) bytes to disk: %s", n, pathAndFilename)
	return nil
}
