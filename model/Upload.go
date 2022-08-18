package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-vue-blog/utils"
	"go-vue-blog/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

// UploadFile 文件上传
// Params file multipart.File 	待访问的文件
// Params fileSize int64		待上传的文件大小
func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket, //对应七牛云中的某个具体的存储空间
	}
	mac := qbox.NewMac(AccessKey, SecretKey) //返回 AK SK 的一个组合struct
	upToken := putPolicy.UploadToken(mac)    //使用 AK 和 SK 生成上传 Token ,云端便知道要上传到哪个空间

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, //本空间建立在华南的机房
		UseCdnDomains: false,               //不使用CDN
		UseHTTPS:      false,               //不使用https
	}
	formUploader := storage.NewFormUploader(&cfg) //上传执行者

	ret := storage.PutRet{} //上传文件成功后的返回值，主要是 key
	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR

	}

	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
