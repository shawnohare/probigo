package mockstore

type Store struct {
	top     map[string]interface{}
	structs map[string]map[string]interface{}
}

const (
	ErrInsufficientArgs = "Insufficient arguments."
	ErrBadCmd           = "Command %s not supported."
	ErrBadKey           = "Keys and fields need to be strings."
)

func (s Store) get(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	return s.top[args[0]], nil
}

func (s Store) fget(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	key := args[0].(string)
	field := args[1].(string)
	return s.structs[key][field]
}

func (s Store) set(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	switch key := args[0].(type) {
	case string:
		s.top[key] = args[1]
		return nil, nil
	default:
		return nil, errors.New(ErrBadKey)
	}
}

func (s Store) fset(key string, args ...interface{}) (interface{}, error) {
	if len(args) < 3 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	var key string
	var field string
	switch t := args[0].(type) {
	case string:
		key = t
	default:
		return nil, errors.New(ErrBadKey)
	}
	switch t := args[1].(type) {
	case string:
		field = t
	default:
		return nil, errors.New(ErrBadKey)
	}
	s.structs[key][field] = args[2]
}

func (s Store) Do(cmd string, args ...interface{}) (interface{}, error) {
	// Get commands.
	switch {
	case cmd == "GET":
		return s.get(args)
	case cmd == "HGET" || "SISMEMBER" || "ZSCORE":
		return s.fget(args)
	case cmd == "SET":
		return s.get(args)
	case cmd == "HSET" || "SADD" || "ZADD":
		return s.fget(args)
	default:
		return nil, fmt.Errorf(ErrBadCmd, cmd)
	}
}
