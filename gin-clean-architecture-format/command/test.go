package command

import "kouhei-github/sample-gin/service"

func Test() {
	service.DownloadImage("https://d1gab3f923bew6.cloudfront.net/H-10928.JPG")
}
