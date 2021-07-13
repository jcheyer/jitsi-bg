package router

type Option func(*Router)

func WithListener(l string) Option {
	return func(r *Router) {
		r.listener = l
	}
}
