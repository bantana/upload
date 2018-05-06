# 文件上传服务

## Design Goal

- support upload file with: standard single，multipart, pre-signed.
- images support watermark
- upload maxsize limit.
- images support multi type translate.
- images support multi resize.
- images support access autoresize.
- support s3 Object store version 4。
- support local mount (cluster storage) fileSystem ?

## 设计

图片使用场景:

  // todo


REST API:

POST	/upload/standard	(OK, err)
POST	/upload/multipart	(OK, err)
POST	/upload/presigned	(OK, err)

design:

- 一个image type corvert 接口.
- 一个image resize 接口.

## 其他相关信息收集

- HEVC(High Efficiency Video Coding(H.265)
- HEIF(High Efficiency Image File Format) ( .hief or .heic) support 4k or 8k
  解析度。同时声称尺寸比 jpeg 小50%.
- 已知情况，.heic 通过ios分享出去到特定的服务接口的时候会自动转换为jpeg。
- 将图片通过google photo, dropbox等上传的时候是.heic文件。

- webp格式 vp8

- Apple HEIC/HEVC images.

import "go4.org/media/heif"


    Package heif reads HEIF containers, as found in Apple HEIC/HEVC images. This package does not decode images; it only reads the metadata.
    This package is a work in progress and makes no API compatibility promises

https://github.com/pushd/heif

    c++ 实现

ImageMagic 转化:

@bradfitz commented 5 days ago

ImageMagick is still slow and uses a lot of RAM.

We need to bump our ImageMagick git rev to a later version (I saw some fixes just the other day) and also probably use this:

http://www.imagemagick.org/Usage/resize/#read ("Resize During Image Read")

ffmpeg 转化:

标准的h264和h265开源实现, (bug应该是最少的)

nokiatech C++ 实现:

    支持c++, java, ios, android.

## Reference:

- http://nokiatech.github.io/heigh/ High Efficiency Image File Format (HEIGH)
- http://jpgtoheif.com

- https://github.com/kometen/http_post/blob/master/heif.go use ffmpeg

## thinking

- 针对标准实现
```golang
// import golang standard lib
import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// import golang.org/x/images
import (
	_ "golang.org/x/image/webp"
)

// import

```
- 非标实现可临时使用ffmpeg + pipline

		https://github.com/kometen/http_post/blob/master/heif.go

	建议转成标准接口 image 实现封装:


golang image 标准接口:

```golang
import _ "radevio.com/bantana/heif"
```

radevio.com/bantana/heif.go
```golang
func init() {
  image.RegisterFormat("heif", "RIFF????HEIF", Decode, DecodeConfig)
}

var errInvalidFormat = errors.New("heif: invalid format")

func decode(r io.Reader, configOnly bool) (image.Image, image.Config, error) {
		// todo : may be need implement
}

// Decode reads a HEIF image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
  m, _, err := decode(r, false)
    if err != nil {
        return nil, err

    }
      return m, err
}

// DecodeConfig returns the color model and dimensions of a HEIF image without
// decoding the entire image.
DecodeConfig(r io.Reader) (image.Config, error) {
  _, c, err := decode(r, true)
    return c, err
}

```
