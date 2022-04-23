package zk

type Zettel struct {
	Address
	Body     string
	Children map[Address]Zettel
}

func (z *Zettel) NewChild(address Address, body string) Zettel {
	child := Zettel{
		Address: address,
		Body:    body,
	}
	if z.Children == nil {
		z.Children = map[Address]Zettel{
			address: child,
		}
	} else {
		z.Children[address] = child
	}
	return child
}
