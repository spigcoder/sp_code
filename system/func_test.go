package main

import (
	"fmt"
	"github.com/spigcoder/sp_code/pkg/snowflake"
	"github.com/spigcoder/sp_code/system/utils/bcrypt"
	"testing"
)

func TestGenFunc(t *testing.T) {
	snowflake.Init(10)
	fmt.Println(snowflake.GenID())
	fmt.Println(bcrypt.Encrypt("123456"))
}
