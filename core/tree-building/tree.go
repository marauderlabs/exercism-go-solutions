package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record stores the id of the node and its parent ID
type Record struct {
	ID     int
	Parent int
}

// Node represents a node in a tree
type Node struct {
	ID       int
	Children []*Node
}

const rootID = 0

// Build buils the tree given the slice of records
func Build(r []Record) (*Node, error) {

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})

	nodes := make(map[int]*Node)

	for i, rec := range r {
		id, p := rec.ID, rec.Parent
		if i != id || p > id || (p == id && id != 0) {
			errmsg := fmt.Sprintf("Unexpected ID in record: %v", rec)
			return nil, errors.New(errmsg)
		}

		nodes[id] = &Node{ID: id}
		if i != rootID {
			myPar := nodes[p]
			myPar.Children = append(myPar.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
