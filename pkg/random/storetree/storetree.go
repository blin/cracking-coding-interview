package storetree

import (
	"encoding/binary"
	"io"
	"math/big"
	"strings"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

const (
	Left  = 0
	Right = 1
)

func StoreTree(t *rbt.Tree, w io.Writer) {
	binary.Write(w, binary.LittleEndian, int64(t.Size()))

	var path big.Int
	record(t.Root, &path, 0, w)
}

func readablePath(path *big.Int, depth uint8) string {
	var sPath []string

	for i := uint8(1); i <= depth; i++ {
		direction := path.Bit(int(i))
		if direction == Left {
			sPath = append(sPath, "left")
		} else {
			sPath = append(sPath, "right")
		}
	}

	return strings.Join(sPath, ",")
}

func record(n *rbt.Node, path *big.Int, depth uint8, w io.Writer) {
	if n == nil {
		return
	}

	binary.Write(w, binary.LittleEndian, depth)
	binary.Write(w, binary.LittleEndian, path.Int64())
	binary.Write(w, binary.LittleEndian, int64(n.Key.(int)))
	binary.Write(w, binary.LittleEndian, int64(n.Value.(int)))

	var leftPath big.Int
	leftPath.Set(path)
	leftPath.SetBit(&leftPath, int(depth+1), Left)
	record(n.Left, &leftPath, depth+1, w)

	var rightPath big.Int
	rightPath.Set(path)
	rightPath.SetBit(&rightPath, int(depth+1), Right)
	record(n.Right, &rightPath, depth+1, w)
}
