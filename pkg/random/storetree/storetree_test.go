package storetree

import (
	"bytes"
	"reflect"
	"testing"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func TestStoreTree(t *testing.T) {
	testCases := []struct {
		desc          string
		valueSequence []int
		want          []byte
	}{
		{
			desc:          "Simplest case",
			valueSequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []byte{
				9, 0, 0, 0, 0, 0, 0, 0, //size

				0,                        //depth
				0b0, 0, 0, 0, 0, 0, 0, 0, //path
				4, 0, 0, 0, 0, 0, 0, 0, //key
				4, 0, 0, 0, 0, 0, 0, 0, //value

				1,                         //depth
				0b00, 0, 0, 0, 0, 0, 0, 0, //path
				2, 0, 0, 0, 0, 0, 0, 0, //key
				2, 0, 0, 0, 0, 0, 0, 0, //value

				2,                          //depth
				0b000, 0, 0, 0, 0, 0, 0, 0, //path
				1, 0, 0, 0, 0, 0, 0, 0, //key
				1, 0, 0, 0, 0, 0, 0, 0, //value

				2,                          //depth
				0b100, 0, 0, 0, 0, 0, 0, 0, //path
				3, 0, 0, 0, 0, 0, 0, 0, //key
				3, 0, 0, 0, 0, 0, 0, 0, //value

				1,                         //depth
				0b10, 0, 0, 0, 0, 0, 0, 0, //path
				6, 0, 0, 0, 0, 0, 0, 0, //key
				6, 0, 0, 0, 0, 0, 0, 0, //value

				2,                          //depth
				0b010, 0, 0, 0, 0, 0, 0, 0, //path
				5, 0, 0, 0, 0, 0, 0, 0, //key
				5, 0, 0, 0, 0, 0, 0, 0, //value

				2,                          //depth
				0b110, 0, 0, 0, 0, 0, 0, 0, //path
				8, 0, 0, 0, 0, 0, 0, 0, //key
				8, 0, 0, 0, 0, 0, 0, 0, //value

				3,                           //depth
				0b0110, 0, 0, 0, 0, 0, 0, 0, //path
				7, 0, 0, 0, 0, 0, 0, 0, //key
				7, 0, 0, 0, 0, 0, 0, 0, //value

				3,                           //depth
				0b1110, 0, 0, 0, 0, 0, 0, 0, //path
				9, 0, 0, 0, 0, 0, 0, 0, //key
				9, 0, 0, 0, 0, 0, 0, 0, //value
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := rbt.NewWithIntComparator()
			for _, v := range tC.valueSequence {
				tree.Put(v, v)
			}

			var bw bytes.Buffer

			StoreTree(tree, &bw)

			got := bw.Bytes()
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("got=\n%v\n, want=\n%v\n", got, tC.want)
			}
		})
	}
}
