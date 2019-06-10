package main

import (
	"context"
	"fmt"
)

type UserContext1 struct {
	userName string
	userAge   int
	userConfigs map[string]interface{}
}

func (c *UserContext1) GetUserAge() int {
	return c.userAge
}

func (c *UserContext1) GetUserName() string {
	return c.userName
}

func (c *UserContext1) GetUserConfMap() map[string]interface{} {
	return c.userConfigs
}

type key1 struct{}

var contextKey1 = &key1{}

func NewContext1(ctx context.Context,fc *UserContext1)context.Context{
	return context.WithValue(ctx,contextKey1,fc)
}
func FromContext1(ctx context.Context) (*UserContext1, bool) {
	fc, ok := ctx.Value(contextKey1).(*UserContext1)
	return fc, ok
}

func NewUserContext1() *UserContext1 {
	un := new(UserContext1)
	un.userAge = 20
	un.userName = "狗蛋"
	return un
}
func main() {
	un := NewUserContext1()
	ctx := NewContext1(context.Background(),un)

	if fc, ok := FromContext1(ctx); ok {
		fmt.Printf("user age is: %d \n", fc.GetUserAge())
		fmt.Printf("user name is: %s \n", fc.GetUserName())
	}
}