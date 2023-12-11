package trace

type Funcs interface {
	Funcs() map[string]func(bool)
}

type Descriptions interface {
	Descriptions() []byte
}
