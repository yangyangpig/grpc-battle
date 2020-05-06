package dbdriver

type Driver interface {
	Insert()
	Update()
	Query()
	Delete()
}
