package service

import (
	"errors"
	"strings"
)

// StringService 字符串服务的接口
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// NewStringService 返回 StringService 接口的一个实现
func NewStringService() *stringService {
	return &stringService{}
}
