// Create by Yale 2019/2/26 11:48
package coss


import (
"github.com/aliyun/aliyun-oss-go-sdk/oss"
"net/http"
)

type OSS struct {
	client *oss.Client
}
type OSSNode struct {
	Key string
	Name string
}
type OSSBucket struct {
	bucket *oss.Bucket
}

func NewOSS(endpoint, accessKeyID, accessKeySecret string, options ...oss.ClientOption) (*OSS,error) {
	clt,err:=oss.New(endpoint,accessKeyID,accessKeySecret,options...)
	if err!=nil {
		return nil,err
	}

	return &OSS{client:clt},nil
}
func (oss *OSS)Config(config *oss.Config)*OSS{
	if oss.client!=nil {
		oss.client.Config = config
	}
	return oss
}
func (oss *OSS)Bucket(name string) (*OSSBucket,error) {
	bucket, err := oss.client.Bucket(name)
	if err!=nil {
		return nil,err
	}
	return  &OSSBucket{bucket:bucket},nil
}
func (o*OSSBucket)SignURL(objectKey string,expiredInSec int64) (string ,error)  {
	return 	o.bucket.SignURL(objectKey,oss.HTTPGet,expiredInSec)
}
func (o *OSSBucket)SetObjectMeta(objectKey string ,opt ...oss.Option) error {
	return o.bucket.SetObjectMeta(objectKey,opt...)
}

func (o *OSSBucket)GetObjectMetaKey(objectKey,metaKey string) (string,error)  {
	h,err:= o.bucket.GetObjectDetailedMeta(objectKey)
	if err!=nil {
		return "",err
	}
	return h.Get("X-Oss-Meta-"+metaKey),nil
}
func (o *OSSBucket)GetObjectMeta(objectKey string) (http.Header,error)  {
	return o.bucket.GetObjectMeta(objectKey)
}
func (o *OSSBucket)GetObjectDetailedMeta(objectKey string, options ...oss.Option) (http.Header,error)  {
	return o.bucket.GetObjectDetailedMeta(objectKey,options...)
}
func (o *OSSBucket)UploadFile(objectKey, filePath string, opt ...oss.Option) error  {
	return o.bucket.PutObjectFromFile(objectKey,filePath,opt...)
}

/////
func (o *OSSBucket)IsObjectFileExist(objectKey string) bool  {
	ret,err:= o.bucket.IsObjectExist(objectKey)
	if err!=nil {
		return false
	}
	return ret
}

func (o *OSSBucket)GetObjectFileMetaKey(objectKey,metaKey string) (string,error)  {

	h,err:= o.bucket.GetObjectDetailedMeta(objectKey)
	if err!=nil {
		return "",err
	}
	return h.Get("X-Oss-Meta-"+metaKey),nil
}
//////
func (o *OSSBucket)IsObjectExist(objectKey string) bool  {
	ret,err:= o.bucket.IsObjectExist(objectKey)
	if err!=nil {
		return false
	}
	return ret
}
func (o *OSSBucket) ListKeyNames(max,maker,pre oss.Option) ([]*OSSNode,error){
	nodes:=make([]*OSSNode,0)
	mk:=maker
	for {
		lor, err := o.bucket.ListObjects(max, mk, pre)
		if err != nil {
			return nodes,err
		}

		pre = oss.Prefix(lor.Prefix)
		mk = oss.Marker(lor.NextMarker)
		for _, object := range lor.Objects {
			if len(object.Key)>0 {
				node:=OSSNode{Key:object.Key}
				/*mn,err:=o.GetObjectMetaKey(object.Key,"name")
				if err==nil {
					node.Name = mn
				}*/
				nodes = append(nodes,&node)
			}
		}
		if !lor.IsTruncated {
			break
		}

	}
	return nodes,nil
}
func (o *OSSBucket) ListKeysSize(max,maker,pre oss.Option) (int64,error){

	count:=0
	mk:=maker
	for {
		lor, err := o.bucket.ListObjects(max, mk, pre)
		if err != nil {
			return 0,err
		}

		pre = oss.Prefix(lor.Prefix)
		mk = oss.Marker(lor.NextMarker)
		for _, object := range lor.Objects {
			if len(object.Key)>0 {
				count++
			}

		}

		if !lor.IsTruncated {
			break
		}

	}
	return int64(count),nil

}

