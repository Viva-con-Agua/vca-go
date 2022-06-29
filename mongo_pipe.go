package vcago

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoPipe struct {
	Pipe []bson.D
}

func NewMongoPipe() *MongoPipe {
	print("vcago.MongoPipe will be deleted soon. Use vmdb.Pipeline")
	return &MongoPipe{
		Pipe: []bson.D{},
	}
}

func (i *MongoPipe) LookupUnwind(from string, root string, child string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$" + as}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}

func (i *MongoPipe) Lookup(from string, root string, child string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "as", Value: as},
		}}}
	i.Pipe = append(i.Pipe, lookup)
}

func (i *MongoPipe) Match(m *MongoMatch) {
	if *m != nil {
		match := bson.D{{
			Key:   "$match",
			Value: *m,
		}}
		i.Pipe = append(i.Pipe, match)
	}
}
