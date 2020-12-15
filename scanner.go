package rxscan

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

func parse(match string, arg interface{}) (err error) {
	switch v := arg.(type) {
	case *bool:
		*v, err = strconv.ParseBool(match)
	//case *complex64:
	//vv, err := strconv.ParseComplex(match, 64)
	//if err != nil {
	//	return err
	//}
	//*v = complex64(vv)
	//case *complex128:
	//vv, err := strconv.ParseComplex(match, 128)
	//if err != nil {
	//	return err
	//}
	//*v = vv
	case *int:
		*v, err = strconv.Atoi(match)
	case *int8:
		vv, err := strconv.ParseInt(match, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(vv)
	case *int16:
		vv, err := strconv.ParseInt(match, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(vv)
	case *int32:
		vv, err := strconv.ParseInt(match, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(vv)
	case *int64:
		vv, err := strconv.ParseInt(match, 10, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *uint:
		vv, err := strconv.ParseUint(match, 10, 64)
		if err != nil {
			return err
		}
		*v = uint(vv)
	case *uint8:
		vv, err := strconv.ParseUint(match, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(vv)
	case *uint16:
		vv, err := strconv.ParseUint(match, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(vv)
	case *uint32:
		vv, err := strconv.ParseUint(match, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(vv)
	case *uint64:
		vv, err := strconv.ParseUint(match, 10, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *uintptr:
		err = errors.New("uintptr is not supported yet")
	case *float32:
		vv, err := strconv.ParseFloat(match, 32)
		if err != nil {
			return err
		}
		*v = float32(vv)
	case *float64:
		vv, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return err
		}
		*v = vv
	case *string:
		*v = match
	case *[]byte:
		*v = []byte(match)
	default:
		err = errors.New("can't scan type: " + reflect.TypeOf(arg).String())
	}

	return err
}

// Scan string using regular expression to variables arguments.
// It returns the number variables successfully parsed. Variable arguments can be less than
// the capture group but it will return an error if variables are more than the capture group
func Scan(re *regexp.Regexp, s string, args ...interface{}) (n int, err error) {
	matches := re.FindStringSubmatch(s)
	if len(matches) <= 1 {
		return 0, nil
	}

	if len(args) > len(matches)-1 {
		return 0, errors.New("got " + strconv.Itoa(len(args)) + " arguments for " + strconv.Itoa(len(matches)-1) + " matches")
	}

	for i, arg := range args {
		if arg == nil {
			continue
		}
		if err := parse(matches[i+1], arg); err != nil {
			return 0, err
		}
		n++
	}
	return n, err
}

type Scanner struct {
	matches [][]string
	i       int
	args    []interface{}
	err     error
}

func (s *Scanner) Error() error {
	return s.err
}

// NewScanner returns a scanner that can scan all repeating regular expression within a text
func NewScanner(re *regexp.Regexp, s string) *Scanner {
	return &Scanner{
		matches: re.FindAllStringSubmatch(s, -1),
	}
}

// More returns true if there is more matches
func (s *Scanner) More() bool {
	return s.err == nil && s.i < len(s.matches)
}

// Scan matched regular expresion to the variables.
// It returns the number variables successfully parsed. Variable arguments can be less than
// the capture group but it will return an error if variables are more than the capture group
func (s *Scanner) Scan(args ...interface{}) (int, error) {
	m := s.matches[s.i]
	if len(s.args) > len(m)-1 {
		s.err = errors.New("got " + strconv.Itoa(len(s.args)) + " arguments for " + strconv.Itoa(len(m)-1) + " matches")
		return 0, s.err
	}

	parsed := 0
	if len(m) > 1 {
		for i, arg := range args {
			if arg == nil {
				continue
			}
			if s.err = parse(m[i+1], arg); s.err != nil {
				return parsed, s.err
			}
			parsed++
		}
	}

	s.i++

	return parsed, s.err
}
