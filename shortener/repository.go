package shortener

// RedirectRepository handle redirect persistence interface
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
