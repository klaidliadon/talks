package node

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func NewLeaf(name string) *Node {
	return &Node{TypeLeaf, &Leaf{name}}
}

type Leaf struct {
	name string
}

func (l *Leaf) Name() string { return l.name }

func (l *Leaf) Children() []*Node { return nil }

func (l *Leaf) GetBSON() (interface{}, error) {
	return bson.M{
		"type": TypeLeaf,
	}, nil
}

func (l *Leaf) SetBSON(raw bson.Raw) error {
	return nil
}

func (l *Leaf) String() string { return fmt.Sprintf("(%s)", l.name) }

func (l *Leaf) setup(name string) { l.name = name }
