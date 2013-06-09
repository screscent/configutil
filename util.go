package configutil

import (
        "github.com/screscent/config"
	"errors"
	"reflect"
)

type section_key_value struct {
	section string
	key     string
	value   interface{}
}

var conf map[string]map[string]*section_key_value

func init() {
	conf = make(map[string]map[string]*section_key_value)
}

func AddSectionKeyValue(section, key string, value interface{}) error {
	conf_sec, ok := conf[section]
	if !ok {
		conf_sec = make(map[string]*section_key_value)
		conf[section] = conf_sec
	}

	sec_key_v, ok := conf_sec[key]
	if !ok {
		sec_key_v = &section_key_value{section, key, value}
		conf_sec[key] = sec_key_v
	} else {
		return errors.New(section + " " + key + " is allready added")
	}
	return nil
}

func AddDefaultSectionKeyValue(key string, value interface{}) error {
	return AddSectionKeyValue(config.DEFAULT_SECTION, key, value)
}

func ReadAll(fname string) error {
	c, err := config.Read(fname)
	if err != nil {
		return nil
	}
	for _, section := range conf {
		for _, node := range section {
			switch v := node.value.(type) {
			case *bool:
				b, err := c.Bool(node.section, node.key)
				if err != nil {
					return err
				}
				*v = b

			case *string:
				s, err := c.String(node.section, node.key)
				if err != nil {
					return err
				}
				*v = string(s)

			case *float64:
				f, err := c.Float(node.section, node.key)
				if err != nil {
					return err
				}
				*v = f

			case *int:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = i

			case *int32:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = int32(i)

			case *int64:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = int64(i)

			case *uint:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = uint(i)

			case *uint32:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = uint32(i)

			case *uint64:
				i, err := c.Int(node.section, node.key)
				if err != nil {
					return err
				}
				*v = uint64(i)

			case *[]string:
				strs, err := c.Strings(node.section, node.key)
				if err != nil {
					return err
				}
				*v = strs

			default:
				return errors.New("do not support type " + reflect.TypeOf(node.value).String())
			}
		}
	}
	return nil
}
