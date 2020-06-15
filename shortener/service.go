package shortener

// RedirectService is an inteface to handle a redirect
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
