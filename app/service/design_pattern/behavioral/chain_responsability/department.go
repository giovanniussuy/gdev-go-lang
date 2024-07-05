package chainresponsability

type Department interface {
	execute(*Patient)
	setNext(Department)
}
