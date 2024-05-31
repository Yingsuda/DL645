package model

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

func str2bcd(addr string) ([]byte, error) {
	var address string
	if len(addr) > 12 {
		fmt.Println("addr length err:", len(addr))
		return nil, fmt.Errorf("addr length max 12 ,real id %d", len(addr))
	} else {
		address = addr
		for i := 0; i < 12-len(addr); i++ {
			address += "0"
		}
	}

	bcd := make([]byte, 6)
	for i := 0; i*4 < len(address); i++ {
		val, err := strconv.ParseUint(address[i*4:i*4+4], 16, 16)
		if err != nil {
			fmt.Println("parse err:", err)
			return nil, fmt.Errorf("addr is %s ;strconv parse err %s", addr, err.Error())
		}
		binary.BigEndian.PutUint16(bcd[i*2:], uint16(val))
	}

	return bcd, nil
}

func TestBCD(t *testing.T) {
	bcd, err := str2bcd("12345")
	if err != nil {
		fmt.Println("str2bcd err:", err)
		return
	}
	fmt.Printf("%x\n", bcd)
}
