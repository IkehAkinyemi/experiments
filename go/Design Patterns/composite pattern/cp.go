package main

import "fmt"

// Component interface
type INode interface {
	print(string)
}

// File struct
type File struct {
	name string
}

// Directory struct
type Directory struct {
	nodes []INode
	name  string
}

// NewFile function
func NewFile(name string) *File {
	return &File{name: name}
}

// NewDirectory function
func NewDirectory(name string) *Directory {
	return &Directory{name: name, nodes: []INode{}}
}

// print method for File
func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

// print method for Directory
func (d *Directory) print(indentation string) {
	fmt.Println(indentation + d.name)
	for _, i := range d.nodes {
		i.print(indentation + indentation)
	}
}

// add method for Directory
func (d *Directory) add(i INode) {
	d.nodes = append(d.nodes, i)
}

func main() {
	root := NewDirectory("root")
	file1 := NewFile("File1")
	file2 := NewFile("File2")
	file3 := NewFile("File3")

	root.add(file1)

	dir1 := NewDirectory("dir1")
	dir1.add(file2)
	dir1.add(file3)
	root.add(dir1)

	root.print(" ")
}
