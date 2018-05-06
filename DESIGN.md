# 文件上传服务

图片使用场景:

  // todo

## Design Goal

- support upload file with: standard single，multipart, pre-signed.
- images support watermark
- upload maxsize limit.
- images support multi type convert.
- images support multi resize.
- images support access auto resize.
- support s3 Object store version 4。
- support local mount (cluster storage) fileSystem ?
- OneShot : Files are destructed after the first download
- Removable : Give the ability to the uploader to remove files at any time
- TTL : Custom expiration date
- Password : Protect upload with login/password (Auth Basic)

## 需求分析

  设计需求变大了

  重新划分成几个小的设计

  1. upload 设计
  2. http get access 设计
  3. image convert 设计
  4. image resize 设计
  5. image waterMark 设计
  6. multiple backend storage interface
  7. multiple metadata backends: Bolt, badger,...


重新考虑的事:

- file 包含 image, audio, movie , ...
- upload file  需要和image区分开来吗?
- file 是可以有 mime 识别的.
- file 是可以有 getMeta 设计的.
- image 是可以有 convert 和 resize 设计的.
- audio ?
- video ?

使用设计一:

  1. upload file , with a msg register
  2. use msg register detect file format and trigger defined actions(like image convert common ImageType)

使用设计二:

  1. upload file
  2. access file with define action.(like convert or resize, etc...)

## 设计


REST API

POST	/upload/standard	({$UUID}, {$MIME}, err)

POST	/upload/multipart	({$UUID}, {$MIME}, err)
// all files with json( []{"file1", "$uuidFiel1", "$mimeFile1", err} )

POST	/upload/presigned	({$UUID}, {$MIME}, err)

GET   /image/{$UUID}/{$ImageType}/{$SIZE} ([]byte or err) // ImageType is [png|gif|heif|webp|...]

GET   /file/{$UUID}

design thinking:

- 一个 image type convert 接口.
- 一个 image resize 接口.
- 一个 backend storage 接口.

## 其他相关信息收集

- HEVC(High Efficiency Video Coding(H.265)
- HEIF(High Efficiency Image File Format) ( .hief or .heic) support 4k or 8k
  解析度。同时声称尺寸比 jpeg 小50%.
- 已知情况，.heic 通过ios分享出去到特定的服务接口的时候会自动转换为jpeg。
- 将图片通过google photo, dropbox等上传的时候是.heic文件。

# TL;DR

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
-

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

examples:

```golang
// example
import (
 "image/jpeg"
 "image/png"
 "io"
)

// convertJPEGToPNG converts from JPEG to PNG.
func convertJPEGToPNG(w io.Writer, r io.Reader) error {
 img, err := jpeg.Decode(r)
 if err != nil {
  return err
 }
 return png.Encode(w, img)
}

// example
import (
 "image"
 "image/png"
 "io"

 _ "code.google.com/p/vp8-go/webp"
 _ "image/jpeg"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
 img, _, err := image.Decode(r)
 if err != nil {
  return err
 }
 return png.Encode(w, img)
}
```
