package screenshot

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func TakeScreenshot(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ctx, cancelBrowser := chromedp.NewContext(ctx)
	defer cancelBrowser()

	var buf []byte

	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(&buf, 90),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return nil, err
	}

	return buf, nil
}
