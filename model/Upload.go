package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go_blog/utils"
	"go_blog/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUpLoader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{} // 可选配置
	err = formUpLoader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)

	if err != nil {
		return "", errmsg.ERROR

	}

	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
