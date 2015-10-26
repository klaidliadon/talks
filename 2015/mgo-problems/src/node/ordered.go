package node

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func NewOrdObject(name string, children ...*Node) *Node {
	return &Node{TypeOrdObject, &OrdObject{name, children}}
}

type OrdObject struct {
	name     string
	children []*Node
}

func (o *OrdObject) Name() string { return o.name }

func (o *OrdObject) Children() []*Node { return o.children }

func (o *OrdObject) GetBSON() (interface{}, error) {
	var children bson.D
	for _, c := range o.children {
		children = append(children, bson.DocElem{c.Name(), c})
	}
	return bson.M{
		"children": children,
		"type":     TypeOrdObject,
	}, nil
}

func (o *OrdObject) SetBSON(raw bson.Raw) error {
	var u struct {
		Name     string    `bson:"name"`
		Children bson.RawD `bson:"children`
	}
	if err := raw.Unmarshal(&u); err != nil {
		return err
	}
	o.name = u.Name
	o.children = make([]*Node, 0, len(u.Children))
	for _, c := range u.Children {
		var n Node
		if err := c.Value.Unmarshal(&n); err != nil {
			return err
		}
		n.setup(c.Name)
		o.children = append(o.children, &n)
	}
	return nil
}

func (o *OrdObject) String() string {
	return fmt.Sprintf("[%s]{%s}", o.name, o.children)
}

func (o *OrdObject) setup(name string) { o.name = name }
