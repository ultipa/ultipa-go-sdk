package structs

import "errors"

type Values struct {
	data map[string]interface{}
}

func NewValues() *Values {
	return &Values{
		data: map[string]interface{}{},
	}
}

func (v *Values) Set(key string, value interface{}) {
	v.data[key] = value
}

func (v *Values) Get(key string) interface{} {
	return v.data[key]
}

func (v *Values) Has(key string) bool {
	return v.Get(key) != nil
}

func (v *Values) ForEach(cb func(v interface{}, key string) error, order []string) error {

	if order == nil {

		for key, v := range v.data {
			err := cb(v,key)
			if err != nil {
				return err
			}
		}

		return nil
	}


	for _, key := range order {

		vv := v.Get(key)
		if vv == nil {
			return errors.New("Key : " + key + "does not exist!")
		}

		err := cb(vv, key)

		if err != nil {
			return err
		}
	}

	return nil
}