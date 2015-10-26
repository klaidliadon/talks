package main

import (
	"fmt"
	"os"

	. "node"

	"gopkg.in/klaidliadon/console.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = console.New(console.Cfg{
	Color: true,
	File:  console.FileShow,
	Date:  console.DateHour,
}, os.Stdout)

func main() {
	compare(Schema{bson.NewObjectId(), "root", NewOrdObject("root",
		NewLeaf("a"),
		NewOrdObject("b",
			NewLeaf("b1"),
			NewLeaf("b2"),
		),
		NewList(NewLeaf("c")),
	)})
	compare(Schema{bson.NewObjectId(), "root", NewObject("root",
		NewLeaf("a"),
		NewObject("b",
			NewLeaf("b1"),
			NewLeaf("b2"),
		),
		NewList(NewLeaf("c")),
	)})
}

func compare(a Schema) {
	log := log.Clone(fmt.Sprintf("%-13s - ", a.Root.T.String()))
	d, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer d.Close()
	c := d.DB("test").C("nodes")
	if err := c.Insert(a); err != nil {
		panic(err)
	}
	log.Info("Before load %s", a.Root)
	var b Schema
	if err := c.FindId(a.Id).One(&b); err != nil {
		panic(err)
	}
	log.Info("After load  %s", b.Root)
}
