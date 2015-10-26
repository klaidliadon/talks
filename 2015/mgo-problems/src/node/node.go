//go:generate stringer -type=Type

package node

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Schema struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Root *Node         `bson:"root"`
}

func (s *Schema) SetBSON(raw bson.Raw) error {
	var a struct {
		Id   bson.ObjectId `bson:"_id"`
		Name string        `bson:"name"`
		Root *Node         `bson:"root"`
	}
	if err := raw.Unmarshal(&a); err != nil {
		return err
	}
	*s = Schema(a)
	s.Root.setup(a.Name)
	return nil
}

type Type int

const (
	_ Type = iota
	TypeLeaf
	TypeObject
	TypeList
	TypeOrdObject
)

type Node struct {
	T Type `bson:"type"`
	INode
}

type INode interface {
	bson.Getter
	bson.Setter
	Name() string
	Children() []*Node
	setup(string)
	String() string
}

func (n *Node) GetBSON() (interface{}, error) {
	switch n.T {
	case TypeLeaf:
		return n.INode.(*Leaf), nil
	case TypeObject:
		return n.INode.(*Object), nil
	case TypeList:
		return n.INode.(*List), nil
	case TypeOrdObject:
		return n.INode.(*OrdObject), nil
	default:
		return nil, fmt.Errorf("Unknown Type %d", n.T)
	}
}

func (n *Node) SetBSON(raw bson.Raw) error {
	var s struct {
		T Type `bson:"type"`
	}
	if err := raw.Unmarshal(&s); err != nil {
		return err
	}
	n.T = s.T
	switch n.T {
	case TypeLeaf:
		var a Leaf
		if err := a.SetBSON(raw); err != nil {
			return err
		}
		n.INode = &a
	case TypeObject:
		var a Object
		if err := a.SetBSON(raw); err != nil {
			return err
		}
		n.INode = &a
	case TypeList:
		var a List
		if err := a.SetBSON(raw); err != nil {
			return err
		}
		n.INode = &a
	case TypeOrdObject:
		var a OrdObject
		if err := a.SetBSON(raw); err != nil {
			return err
		}
		n.INode = &a
	default:
		return fmt.Errorf("Unknown Type: %s", s.T)
	}
	return nil
}
