/**

Components
1) import section
2) Define the root node
3) Main Function


**/

package main

// import section

import (
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"log"
)

// Define the root node

/*
*

1) Why fs.Inode ?? embeddings
Embedding means RootNode automatically gets ALL the methods that  fs.Inode  has
It's like saying "RootNode IS an Inode, plus anything extra I want to add"
You get methods like  NewInode() ,  Parent() ,  Children()  for FREE without writing them
This is how go-fuse knows your RootNode is a valid filesystem node
*
*/
type RootNode struct {
	fs.Inode
}

// Main Function

/*
*
1) the file system is mounted by calling mount on the root of the tree
2) Pass low Level fuse operations, instead defualts will be passed
3) Waits until umount or ctrl+c

*
*/
func main() {
	root := &RootNode{}

	server, err := fs.Mount("/Volumes/go-fs", root, &fs.Options{
		MountOptions: fuse.MountOptions{
			Name:  "gofs",
			Debug: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(" Mounted at /Volumes/go-fs")
	server.Wait()

}
