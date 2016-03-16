// Package mockstore provides a mock Store interface for use with testing
// the probigo package.
package mockstore

import (
	"errors"
	"fmt"
)

type Store struct {
	top     map[string]interface{}
	structs map[string]map[string]interface{}
}

const (
	ErrInsufficientArgs = "Insufficient arguments."
	ErrBadCmd           = "Command %s not supported."
	ErrBadKey           = "Keys and fields need to be strings."
	ErrArgNotString     = "Argument is not a string."
)

func asString(arg interface{}) (string, error) {
	switch t := arg.(type) {
	case string:
		return t, nil
	default:
		return "", errors.New(ErrArgNotString)
	}
}

func (s Store) get(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	key, err := asString(args[0])
	if err != nil {
		return nil, err
	}
	return s.top[key], nil
}

func (s Store) fget(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	key, err := asString(args[0])
	if err != nil {
		return nil, err
	}
	field, err := asString(args[1])
	if err != nil {
		return nil, err
	}
	return s.structs[key][field], nil
}

func (s Store) set(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	key, err := asString(args[0])
	if err != nil {
		return nil, err
	}
	s.top[key] = args[1]
	return nil, nil
}

func (s Store) fset(args ...interface{}) (interface{}, error) {
	if len(args) < 3 {
		return nil, errors.New(ErrInsufficientArgs)
	}
	key, err := asString(args[0])
	if err != nil {
		return nil, err
	}
	field, err := asString(args[1])
	if err != nil {
		return nil, err
	}
	if _, ok := s.structs[key]; !ok {
		s.structs[key] = make(map[string]interface{})
	}
	s.structs[key][field] = args[2]
	return nil, nil
}

func (s Store) Do(cmd string, args ...interface{}) (interface{}, error) {
	// Get commands.
	switch {
	case cmd == "GET":
		return s.get(args)
	case cmd == "HGET" || cmd == "SISMEMBER" || cmd == "ZSCORE":
		return s.fget(args)
	case cmd == "SET":
		return s.get(args)
	case cmd == "HSET" || cmd == "SADD" || cmd == "ZADD":
		return s.fget(args)
	default:
		return nil, fmt.Errorf(ErrBadCmd, cmd)
	}
}

func New() *Store {
	return &Store{
		top:     make(map[string]interface{}),
		structs: make(map[string]map[string]interface{}),
	}
}
