package convert

type Stringer interface {
	ToString() string
}

type Byter interface {
	ToByte() []byte
}

type Converter interface {
	Stringer
	Byter
}
