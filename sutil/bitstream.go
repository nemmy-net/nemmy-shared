package sutil

import (
	"errors"
	"fmt"
)

type BitStream struct {
	Buffer   []byte
	ReadPos  int
	WritePos int
}

func (stream *BitStream) Get(args ...any) error {
	for i := range args {
		switch args[i].(type) {
		case *int8:
			i64, err := stream.GetInt(1)
			if err != nil {
				return err
			}
			*args[i].(*int8) = int8(i64)
		case *int16:
			i64, err := stream.GetInt(16 / 8)
			if err != nil {
				return err
			}
			*args[i].(*int16) = int16(i64)
		case *int32:
			i64, err := stream.GetInt(32 / 8)
			if err != nil {
				return err
			}
			*args[i].(*int32) = int32(i64)
		case *int64:
			i64, err := stream.GetInt(64 / 8)
			if err != nil {
				return err
			}
			*args[i].(*int64) = int64(i64)
		case *uint8:
			u64, err := stream.GetUint(1)
			if err != nil {
				return err
			}
			*args[i].(*uint8) = uint8(u64)
		case *uint16:
			u64, err := stream.GetUint(16 / 8)
			if err != nil {
				return err
			}
			*args[i].(*uint16) = uint16(u64)
		case *uint32:
			u64, err := stream.GetUint(32 / 8)
			if err != nil {
				return err
			}
			*args[i].(*uint32) = uint32(u64)
		case *uint64:
			u64, err := stream.GetUint(64 / 8)
			if err != nil {
				return err
			}
			*args[i].(*uint64) = uint64(u64)
		case int:
			return fmt.Errorf("args[%v] type is int, which has unknown size", i)
		case uint:
			return fmt.Errorf("args[%v] type is uint, which has unknown size", i)
		default:
			return fmt.Errorf("args[%v] type is a non-pointer or is not supported", i)
		}
	}
	return nil
}

func (stream *BitStream) Put(args ...any) error {
	for i := range args {
		var err error

		switch args[i].(type) {
		case int8:
			err = stream.PutInt(int64(args[i].(int8)), 1)
		case *int8:
			err = stream.PutInt(int64(*args[i].(*int8)), 1)
		case int16:
			err = stream.PutInt(int64(args[i].(int16)), 16/8)
		case *int16:
			err = stream.PutInt(int64(*args[i].(*int16)), 16/8)
		case int32:
			err = stream.PutInt(int64(args[i].(int32)), 32/8)
		case *int32:
			err = stream.PutInt(int64(*args[i].(*int32)), 32/8)
		case int64:
			err = stream.PutInt(int64(args[i].(int64)), 64/8)
		case *int64:
			err = stream.PutInt(int64(*args[i].(*int64)), 64/8)
		case uint8:
			err = stream.PutUint(uint64(args[i].(uint8)), 1)
		case *uint8:
			err = stream.PutUint(uint64(*args[i].(*uint8)), 1)
		case uint16:
			err = stream.PutUint(uint64(args[i].(uint16)), 16/8)
		case *uint16:
			err = stream.PutUint(uint64(*args[i].(*uint16)), 16/8)
		case uint32:
			err = stream.PutUint(uint64(args[i].(uint32)), 32/8)
		case *uint32:
			err = stream.PutUint(uint64(*args[i].(*uint32)), 32/8)
		case uint64:
			err = stream.PutUint(uint64(args[i].(uint64)), 64/8)
		case *uint64:
			err = stream.PutUint(uint64(*args[i].(*uint64)), 64/8)
		case int:
			return fmt.Errorf("args[%v] type is int, which has unknown size", i)
		case uint:
			return fmt.Errorf("args[%v] type is uint, which has unknown size", i)
		default:
			return fmt.Errorf("args[%v] type is a pointer or is not supported", i)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (stream *BitStream) PutStr(str string) error {
	length := len(str)
	if stream.WritePos+length > len(stream.Buffer) {
		return errors.New("attempt to write past end of bit stream")
	}
	for i := 0; i < length; i++ {
		stream.Buffer[stream.WritePos+i] = str[i]
	}
	stream.WritePos += length
	return nil
}
func (stream *BitStream) PutInt(ival int64, byteCount int) error {
	length := byteCount
	if stream.WritePos+length > len(stream.Buffer) {
		return errors.New("attempt to write past end of bit stream")
	}
	for i := 0; i < length; i++ {
		stream.Buffer[stream.WritePos+i] = byte((ival >> (8 * (byteCount - 1 - i))) & 0xFF)
	}
	stream.WritePos += length
	return nil
}
func (stream *BitStream) PutUint(ival uint64, byteCount int) error {
	length := byteCount
	if stream.WritePos+length > len(stream.Buffer) {
		return errors.New("attempt to write past end of bit stream")
	}
	for i := 0; i < length; i++ {
		stream.Buffer[stream.WritePos+i] = byte((ival >> (8 * (byteCount - 1 - i))) & 0xFF)
	}
	stream.WritePos += length
	return nil
}

func (stream *BitStream) GetStr(length int) (string, error) {
	if stream.ReadPos+length > len(stream.Buffer) {
		return "", errors.New("attempt to read past end of bit stream")
	}
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = stream.Buffer[stream.ReadPos+i]
	}
	stream.ReadPos += length
	return string(str), nil
}
func (stream *BitStream) GetInt(byteCount int) (int64, error) {
	length := byteCount
	if stream.ReadPos+length > len(stream.Buffer) {
		return 0, errors.New("attempt to read past end of bit stream")
	}
	var value int64 = 0
	for i := 0; i < length; i++ {
		value <<= 8
		value |= int64(stream.Buffer[stream.ReadPos+i]) & 0xFF
	}
	stream.ReadPos += length
	return value, nil
}
func (stream *BitStream) GetUint(byteCount int) (uint64, error) {
	length := byteCount
	if stream.ReadPos+length > len(stream.Buffer) {
		return 0, errors.New("attempt to read past end of bit stream")
	}
	var value uint64 = 0
	for i := 0; i < length; i++ {
		value <<= 8
		value |= uint64(stream.Buffer[stream.ReadPos+i]) & 0xFF
	}
	stream.ReadPos += length
	return value, nil
}
