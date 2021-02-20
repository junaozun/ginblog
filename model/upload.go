package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	AccessKey = utils.AccessKey
	ScretKey  = utils.SecretKey
	Bucket    = utils.Bucket
	ImgUrl    = utils.QiniuServer
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, ScretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false, //收费
		UseHTTPS:      false, // 收费
	}

	putExtra := storage.PutExtra{}

	formUpLoader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUpLoader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS

}
