package store

// Conn interfaces represent a very high-level connection to some
// underlying data structure store.  The Do method is inteded to
// subsume both command issuance and low-level connection management,
// such as closing a connection if necessary after a command.
type Conn interface {
	Do(cmd string, args ...interface{}) (reply interface{}, err error)
}

// conn allows users to promote a function to a Conn interface.
type conn struct {
	do func(string, ...interface{}) (interface{}, error)
}

// Do invokes the doer's do method.
func (d *conn) Do(cmd string, args ...interface{}) (interface{}, error) {
	return d.do(cmd, args...)
}

// NewConn promotes a function to a Conn interface.
func NewConn(doFunc func(string, ...interface{}) (interface{}, error)) Conn {
	return &conn{do: doFunc}
}
