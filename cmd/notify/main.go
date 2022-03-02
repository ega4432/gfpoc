package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

var (
	URL      = "https://notify-api.line.me/api/notify"
	filename string
)

func init() {
	flag.StringVar(&filename, "f", "image.png", "")
	flag.Parse()
}

func main() {
	msg := "send an image"
	accessToken := "<TOKEN>"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("message")
	if err != nil {
		panic(err)
	}
	if _, err = fw.Write([]byte(msg)); err != nil {
		panic(err)
	}

	part := make(textproto.MIMEHeader)
	part.Set("Content-Disposition", `form-data; name="imageFile"; filename=`+filename)

	img, format, err := checkImageFormat(f, filename)
	if err != nil {
		panic(err)
	}

	if format == "jpeg" {
		part.Set("Content-Type", "image/jpeg")
	} else if format == "png" {
		part.Set("Content-Type", "image/png")
	} else {
		panic("LINE Notify supports only jpeg/png image format")
	}

	fw, err = w.CreatePart(part)
	if err != nil {
		panic(err)
	}

	io.Copy(fw, img)
	w.Close() // boundaryの書き込み
	req, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic("failed to send image, get http status code: " + resp.Status)
	}
	fmt.Println(resp.Body)
}

// checkImageFormat validates an image file is not illegal and
// returns image as io.Reader and file format.
func checkImageFormat(r io.Reader, filename string) (io.Reader, string, error) {
	ext := filepath.Ext(filename)

	var b bytes.Buffer
	if ext == ".jpeg" || ext == ".jpg" || ext == ".JPEG" || ext == ".JPG" {
		ext = "jpeg"
		img, err := jpeg.Decode(r)
		if err != nil {
			return nil, "", err
		}

		if err := jpeg.Encode(&b, img, &jpeg.Options{Quality: 100}); err != nil {
			return nil, "", err
		}
	} else if ext == ".png" || ext == ".PNG" {
		ext = "png"
		img, err := png.Decode(r)
		if err != nil {
			return nil, "", err
		}

		if err = png.Encode(&b, img); err != nil {
			return nil, "", err
		}
	} else {
		return nil, "", errors.New("image format must be jpeg or png")
	}

	return &b, ext, nil
}
