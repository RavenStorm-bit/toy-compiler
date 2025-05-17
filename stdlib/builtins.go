package stdlib

import (
	"fmt"
)

// Builtin represents a built-in function
type Builtin struct {
	Fn func(args ...interface{}) (interface{}, error)
}

// Builtins contains all built-in functions
var Builtins = map[string]*Builtin{
	"print": {
		Fn: func(args ...interface{}) (interface{}, error) {
			for _, arg := range args {
				fmt.Print(arg)
			}
			fmt.Println()
			return nil, nil
		},
	},
	"len": {
		Fn: func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case string:
				return int64(len(arg)), nil
			case []interface{}:
				return int64(len(arg)), nil
			default:
				return nil, fmt.Errorf("argument to `len` not supported, got %T", args[0])
			}
		},
	},
	"type": {
		Fn: func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}
			return fmt.Sprintf("%T", args[0]), nil
		},
	},
}

// GetBuiltin returns a built-in function by name
func GetBuiltin(name string) (*Builtin, bool) {
	builtin, ok := Builtins[name]
	return builtin, ok
}
