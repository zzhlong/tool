package tools

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
	"image/png"
)

//Go 使用这个避免了协成panic导致程序退出
func Go(f func(v []interface{}), v ...interface{}) {
	go func(v []interface{}) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f(v)
	}(v)
}

type apkInfoModel struct {
	BaseImage   string //图标base64字符串
	PkgName     string //软件包名
	PackageName string //软件包包名
	VersionCode int32  //版本号
	VersionName string //版本名称
}

//ApkGetInfo 解析apk应用包
func ApkGetInfo(b *[]byte) (*apkInfoModel, error) {
	res := apkInfoModel{}
	r := bytes.NewReader(*b)
	pkg, err := apk.OpenZipReader(r, r.Size())
	if err != nil {
		return nil, err
	}
	icon, err := pkg.Icon(nil)
	if err != nil {
		return nil, err
	}
	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	png.Encode(emptyBuff, icon)       //img写入到buff
	res.BaseImage = base64.StdEncoding.EncodeToString(emptyBuff.Bytes())
	info := pkg.Manifest()
	res.VersionCode, _ = info.VersionCode.Int32()
	res.VersionName, _ = info.VersionName.String()
	res.PackageName = pkg.PackageName() // returns the package name
	resConfigEN := &androidbinary.ResTableConfig{
		Language: [2]uint8{uint8('e'), uint8('n')},
	}
	res.PkgName, _ = pkg.Label(resConfigEN) // get app label for en translation
	return &res, nil
}
