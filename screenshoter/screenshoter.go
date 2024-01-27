package screenshoter

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Screenshoter struct {
}

type ScreenshoterBehaviour interface {
	TakeScreenshot(ctx context.Context, url string) ([]byte, error)
}

func (s Screenshoter) TakeScreenshot(ctx context.Context, url string) ([]byte, error) {

	var buf []byte
	// capture entire browser viewport, returning png with quality=90

	err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf))

	return buf, err
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	fmt.Println("aaaaaaaaaaaaa")

	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
