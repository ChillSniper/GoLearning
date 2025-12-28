package constlearn

const (
	TypeBooks = iota
	TypePage  = "page"
)

type ReadTester interface {
	ReadTest(p []byte) (n int, err error)
}
