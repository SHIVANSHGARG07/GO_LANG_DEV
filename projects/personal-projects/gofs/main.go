/**

Components
1) import section
2) Define the root node
3) Main Function


**/

package main

// import section

import (
	"context"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"log"
	"syscall"
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

type FileNode struct {
	fs.Inode
	content string
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
			Debug: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(" Mounted at /Volumes/go-fs")
	server.Wait()

}

func (r *RootNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {

	entries := []fuse.DirEntry{
		{
			Name: "hello.txt",
			Mode: fuse.S_IFREG,
		},
	}

	return fs.NewListDirStream(entries), 0

}

func (r *RootNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	log.Printf("🔍 Lookup called for: %s", name)

	// check if file exists
	if name != "hello.txt" {
		return nil, syscall.ENOENT
	}

	//create a node for file
	stable := fs.StableAttr{
		Mode: fuse.S_IFREG,
	}

	// create a filenode with content
	fileNode := &FileNode{content: "Hello from GOFS!!!!"}

	// alternate to newInode
	node := r.NewPersistentInode(ctx, fileNode, stable)

	return node, 0

}

// Open is called when opening the file for reading
func (f *FileNode) Open(ctx context.Context, flags uint32) (fs.FileHandle, uint32, syscall.Errno) {
	log.Println("📂 OPEN called!")
	return nil, 0, 0
}

// after open getattr should be called, else it considers filesie as 0
// Getattr returns file attributes (size, permissions, etc.)
func (f *FileNode) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	log.Println("📊 GETATTR called!")

	// Set file size to the length of our content
	out.Size = uint64(len(f.content))

	// Set file mode (permissions)
	out.Mode = fuse.S_IFREG | 0644 // Regular file, rw-r--r--

	log.Printf("   File size: %d bytes", out.Size)
	return 0
}

// Read is called when someone reads the file (cat)
// before this we need Open method
func (f *FileNode) Read(ctx context.Context, fh fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {

	log.Println("📖 READ CALLED!") // ← Add this

	// convert string to bytes
	data := []byte(f.content)

	// calculate and position
	end := off + int64(len(dest))
	if end > int64(len(data)) {
		end = int64(len(data))
	}

	return fuse.ReadResultData(data[off:end]), 0

}
