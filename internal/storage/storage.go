package storage

type Storage interface {
	Save(p *Page) error
	Get() (Page, error)
	Remove(p *Page)
	// IsExists()
}

type Page struct {
	Name     string
	Meta     string
	Contacts string
}
