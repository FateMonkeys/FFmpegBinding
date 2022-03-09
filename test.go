package main

import (
	"github.com/FateMonkeys/FFmpegBinding"
	"log"
)

func main() {

	InputFile := "/1.mp4"
	OutPathFile :="/2.mp4"

	FFmpegConf := &FFmpegBinding.Config{
		FfmpegBinPath :   "",//ffmpeg目录
		FfprobeBinPath:   "",//ffprobe目录一般与ffmpeg在同目录
		ProgressEnabled:  true,
	}


	options := FFmpegBinding.Options{}//根据自己的需求设置。预留了一些设置 也可以自定义传入



	information:=&FFmpegBinding.Information{}//用于接收ffmpeg进度获取提前结束ffmpeg

	FFmpegBinding.
		New(FFmpegConf).
		SetInput(InputFile).
		SetTimeout(60).//超时单位秒
		SetOutput(OutPathFile).
		WithOptions(options).
		Run(information)


	progress:=information.Progress
	done:=information.Error

	for msg := range progress {
		log.Printf("%+v", msg)
	}


	err := <-done
	if err!=nil {
		log.Println(err)
	}

}
