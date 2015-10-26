package node

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func NewList(sample *Node) *Node {
	return &Node{TypeList, &List{sample}}
}

type List struct {
	sample *Node
}

func (l *List) Name() string { return l.sample.Name() }

func (l *List) Children() []*Node { return []*Node{l.sample} }

func (l *List) GetBSON() (interface{}, error) {
	return bson.M{
		"type":   TypeList,
		"sample": l.sample,
	}, nil
}

func (l *List) SetBSON(raw bson.Raw) error {
	var m struct {
		Sample *Node `bson:"sample"`
	}
	if err := raw.Unmarshal(&m); err != nil {
		return err
	}
	l.sample = m.Sample
	return nil
}

func (l *List) String() string { return fmt.Sprintf("%s*", l.sample) }

func (l *List) setup(name string) { l.sample.setup(name) }
