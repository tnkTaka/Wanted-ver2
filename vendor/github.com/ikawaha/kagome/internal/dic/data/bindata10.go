package data

import (
	"os"
	"time"
)


func dicUniUniDic006Bytes() ([]byte, error) {
	return bindataRead(
		_dicUniUniDic006,
		"dic/uni/uni.dic.006",
	)
}

func dicUniUniDic006() (*asset, error) {
	bytes, err := dicUniUniDic006Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/uni/uni.dic.006", size: 10485760, mode: os.FileMode(420), modTime: time.Unix(1515075872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}