package constlearn

const (
	TypeBooks = iota
	TypePage  = "page"
)

type ReadTester interface {
	ReadTest(p []byte) (n int, err error)
}

func test() {
	const str = "abc"
	const x, y = "1", "2"
	const a, b string = "abc", "def"

}

const (
	UnKnown = 0
	Success = 1
	Fail    = 2
)
