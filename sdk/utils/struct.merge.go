package utils

import (
	"errors"
	"log"
	"reflect"
)

// MergeSameStruct merge struct2 values to struct1, struct2 must have same struct with struct1.
// struct1, and struct2 should be two Ptr.
func MergeSameStruct(struct1 interface{}, struct2 interface{}) error {
	var err error

	s1p := reflect.ValueOf(struct1)
	s2p := reflect.ValueOf(struct2)


	log.Println("Kinds", s1p.Type().Kind(), s2p.Type().Kind())

	if s1p.Type().Kind() != reflect.Ptr || s2p.Type().Kind() != reflect.Ptr {
		return errors.New("struct1 and struct2 should be struct Ptr")
	}

	s1v := s1p.Elem()
	s2v := s2p.Elem()

	if s1v.Type() != s2v.Type() {
		return errors.New("struct1 and struct2 should be same Type")
	}

	for i := 0; i < s2v.NumField(); i++ {
		s2i := s2v.Field(i)
		s2t := s2v.Type().Field(i)
		s1i := s1v.FieldByName(s2t.Name)

		if s1i.CanSet() == false {
			continue
		}

		log.Println("Start merging ", s2t.Name)
		switch s2i.Kind() {
		case reflect.Ptr:
			err := MergeSameStruct(s1i, s2i)
			if err != nil {
				return err
			}
		case reflect.Struct:
			err := MergeSameStruct(&s1i, &s2i)
			if err != nil {
				return err
			}
		case reflect.Slice:
			s1i.Set(reflect.AppendSlice(s1i, s2i))
		default:
			s1v.Set(s2v)
		}
	}

	return err
}
