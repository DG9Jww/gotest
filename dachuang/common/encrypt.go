package common

import (
	"crypto/md5"
	"time"
	"fmt"
	"math/rand"
	"strconv"
)

func MD5encrypt(p string) string {
	p1 := []byte(p)
	p2 := md5.Sum(p1)
	p = fmt.Sprintf("%x", p2)
	return p
}

func RandomNum() string {
	rand.Seed(time.Now().Unix())
	res := strconv.Itoa(rand.Int())
	return res
}