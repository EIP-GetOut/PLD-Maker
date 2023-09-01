package pdf

type Pdf interface {
	SetName(name string)
	GetName() string
	PrintName()
}
