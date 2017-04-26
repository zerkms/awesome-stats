package lib

type formatter interface {
	format(s Stats) (string, error)
}
