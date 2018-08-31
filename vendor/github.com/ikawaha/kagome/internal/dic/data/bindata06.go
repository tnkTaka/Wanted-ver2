package data

import (
	"os"
	"time"
)


func dicUniUniDic002Bytes() ([]byte, error) {
	return bindataRead(
		_dicUniUniDic002,
		"dic/uni/uni.dic.002",
	)
}

func dicUniUniDic002() (*asset, error) {
	bytes, err := dicUniUniDic002Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/uni/uni.dic.002", size: 10485760, mode: os.FileMode(420), modTime: time.Unix(1515075871, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}