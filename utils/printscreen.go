package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetChromeScreenshot(site string, quality int) {
	screenshotURL := fmt.Sprintf("https://%s/", site)
	var buf []byte

	var ext string = "png"
	if quality < 100 {
		ext = "jpeg"
	}

	log.Printf("Capturing %s screenshot...", screenshotURL)
	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1280, 1024))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate(screenshotURL),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.CaptureScreenshot(&buf),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("%s-%d-standard.%s", strings.Replace(site, "/", "-", -1), time.Now().UTC().Unix(), ext)
	if err := os.WriteFile(filepath.Join("tmp", filename), buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("Captura armazenada em %s", filename)
}
