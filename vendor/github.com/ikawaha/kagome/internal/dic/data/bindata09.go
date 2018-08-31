package data

import (
	"os"
	"time"
)


func dicUniUniDic005Bytes() ([]byte, error) {
	return bindataRead(
		_dicUniUniDic005,
		"dic/uni/uni.dic.005",
	)
}

func dicUniUniDic005() (*asset, error) {
	bytes, err := dicUniUniDic005Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/uni/uni.dic.005", size: 10485760, mode: os.FileMode(420), modTime: time.Unix(1515075872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}