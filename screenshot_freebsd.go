package screenshot

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"image"
)

func Connect() (*xgb.Conn, error) {
	c, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Close(c *xgb.Conn) {
	c.Close()
}

func ScreenSize(c *xgb.Conn) [2]int {
	screen := xproto.Setup(c).DefaultScreen(c)
	x := screen.WidthInPixels
	y := screen.HeightInPixels
	return [2]int{int(x), int(y)}
}

func CaptureScreen(c *xgb.Conn) (*image.RGBA, error) {
	screen := xproto.Setup(c).DefaultScreen(c)
	x := screen.WidthInPixels
	y := screen.HeightInPixels
	xImg, err := xproto.GetImage(c, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(0), int16(0), x, y, 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}
	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}
	img := &image.RGBA{data, 4 * int(x), image.Rect(0, 0, int(x), int(y))}
	return img, nil
}
