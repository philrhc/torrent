//go:build !windows

package torrentfs

import (
	"context"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"

	"github.com/anacrolix/torrent"
)

type fileNode struct {
	node
	f *torrent.File
}

var _ fusefs.NodeOpener = fileNode{}

func (fn fileNode) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Size = uint64(fn.f.Length())
	attr.Mode = defaultMode
	return nil
}

func (fn fileNode) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fusefs.Handle, error) {
	return fileHandle{fn, fn.f}, nil
}
