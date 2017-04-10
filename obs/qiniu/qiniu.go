package qiniu

import (
	"qiniupkg.com/api.v7/kodo"
	"github.com/fuyufjh/simple-gallery/obs"

	"io"
	"strings"
	"fmt"
)

const (
	kodoAccessKey  = "aLTwcrN4fPJ2Ed6_F2vqlxUKKCQjALjR9cMszPLl"
	kodoSecretKey  = "pR3IwBakQDUDx0rxE5ywsOsyRK5rmgK95lNgI6Or"
	kodoBucketName = "ericfu-album"
	kodoBucketURL  = "http://oo6ypz1jy.bkt.clouddn.com"
)

type QiniuBucket struct {
	bucket kodo.Bucket
}

func New() *QiniuBucket {
	client := kodo.New(0, &kodo.Config{
		AccessKey: kodoAccessKey,
		SecretKey: kodoSecretKey,
	})
	bucket := client.Bucket(kodoBucketName)
	return &QiniuBucket{
		bucket: bucket,
	}
}

func (b *QiniuBucket) List() ([]*obs.Category, error) {
	marker := ""
	var results []*obs.Category
	for {
		objects, _, nextMarker, err := b.bucket.List(nil, "", "", marker, 0)

		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("Qiniu list bucket error: %v", err)
		}

		for _, obj := range objects {
			var category, name string
			if slash := strings.LastIndex(obj.Key, "/"); slash != -1 {
				category = obj.Key[:slash]
				name = obj.Key[slash+1:]
			} else {
				category = ""
				name = obj.Key
			}

			if len(results) == 0 || results[len(results) - 1].Name != category {
				results = append(results, &obs.Category{
					Name: category,
					Photos: make([]*obs.Photo, 0),
				})
			}
			ctg := results[len(results) - 1]
			ctg.Photos = append(ctg.Photos, &obs.Photo{
				Name: name,
				URL: kodoBucketURL + "/" + obj.Key,
			})
		}

		if err == io.EOF {
			break
		}
		marker = nextMarker
	}
	return results, nil
}
