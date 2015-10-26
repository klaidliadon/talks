package node

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func NewObject(name string, children ...*Node) *Node {
	var cmap = make(map[string]*Node)
	for _, c := range children {
		cmap[c.Name()] = c
	}
	return &Node{TypeObject, &Object{name, cmap}}
}

type Object struct {
	name     string
	children map[string]*Node
}

func (o *Object) Name() string { return o.name }

func (o *Object) Children() []*Node {
	var children = make([]*Node, 0, len(o.children))
	for _, c := range o.children {
		children = append(children, c)
	}
	return children
}

func (o *Object) GetBSON() (interface{}, error) {
	return bson.M{
		"children": o.children,
		"type":     TypeObject,
	}, nil
}

func (o *Object) SetBSON(raw bson.Raw) error {
	var u struct {
		Name     string              `bson:"name"`
		Children map[string]bson.Raw `bson:"children`
	}
	if err := raw.Unmarshal(&u); err != nil {
		return err
	}
	o.name = u.Name
	o.children = make(map[string]*Node)
	for i, c := range u.Children {
		var n Node
		if err := c.Unmarshal(&n); err != nil {
			return err
		}

		o.children[i] = &n
	}
	return nil
}

func (o *Object) String() string {
	return fmt.Sprintf("[%s]{%s}", o.name, o.Children())
}

func (o *Object) setup(name string) {
	o.name = name
	for i, c := range o.children {
		c.setup(i)
	}
}
