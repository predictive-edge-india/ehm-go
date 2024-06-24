package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type TextArray []string

func (x *TextArray) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []string{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			*x = append(*x, "")
			continue
		}
		if s[0:1] == `"` {
			var err error
			s = strings.ReplaceAll(s, `\\u`, `\u`)
			s = strings.ReplaceAll(s, "\n", "\\n")
			s, err = strconv.Unquote(s)
			if err != nil {
				panic(err.Error() + ", string: " + fmt.Sprint(parts))
			}
		}
		*x = append(*x, s)
	}
	return nil
}

func (x TextArray) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		s = strings.ReplaceAll(s, ",", `\u002c`)
		s = strings.ReplaceAll(s, `\`, `\\`)
		s = `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
		formatted = append(formatted, s)
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}

type Float4Array []float32

func (x *Float4Array) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []float32{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			panic("empty string is not a valid float32")
		}
		num, err := strconv.ParseFloat(s, 32)
		if err != nil {
			panic("not a valid float32")
		}
		*x = append(*x, float32(num))
	}
	return nil
}

func (x Float4Array) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		formatted = append(formatted, fmt.Sprint(s))
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}

type Float8Array []float64

func (x *Float8Array) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []float64{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			panic("empty string is not a valid float64")
		}
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("not a valid float64")
		}
		*x = append(*x, num)
	}
	return nil
}

func (x Float8Array) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		formatted = append(formatted, fmt.Sprint(s))
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}

type IntArray []int

func (x *IntArray) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []int{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			panic("empty string is not a valid integer")
		}
		num, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			panic("not a valid integer")
		}
		*x = append(*x, int(num))
	}
	return nil
}

func (x IntArray) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		formatted = append(formatted, fmt.Sprint(s))
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}

type Int4Array []int32

func (x *Int4Array) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []int32{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			panic("empty string is not a valid int32")
		}
		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			panic("not a valid int32")
		}
		*x = append(*x, int32(num))
	}
	return nil
}

func (x Int4Array) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		formatted = append(formatted, fmt.Sprint(s))
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}

type Int8Array []int64

func (x *Int8Array) Scan(value any) error {
	if value == nil {
		*x = nil
		return nil
	}
	str := value.(string)
	if str == "{}" {
		*x = []int64{}
		return nil
	}
	str, _ = strings.CutPrefix(str, "{")
	str, _ = strings.CutSuffix(str, "}")
	parts := strings.Split(str, ",")
	for _, s := range parts {
		if s == "" || s == `""` {
			panic("empty string is not a valid int64")
		}
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic("not a valid int64")
		}
		*x = append(*x, num)
	}
	return nil
}

func (x Int8Array) Value() (driver.Value, error) {
	if len(x) == 0 {
		return "{}", nil
	}
	formatted := []string{}
	for _, s := range x {
		formatted = append(formatted, fmt.Sprint(s))
	}
	str := "{" + strings.Join(formatted, ",") + "}"
	return str, nil
}
