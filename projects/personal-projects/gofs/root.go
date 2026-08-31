package main

import (
	"context"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"log"
	"syscall"
)

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
