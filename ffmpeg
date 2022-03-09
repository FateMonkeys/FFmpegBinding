package FFmpegBinding

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/transcoder/utils"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Information struct {
	Progress chan Progress
	Error    chan error
	Cmd      *exec.Cmd
}

type Progress struct {
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
}

type FFmpeg struct {
	config   *Config
	input    string
	timeout  int
	output   string
	options  [][]string
	metadata Metadata
	setTime  int64
	timer *time.Timer
	information *Information

}

func New(config *Config) *FFmpeg {
	FFmpeg:= &FFmpeg{}
	FFmpeg.config = config
	return FFmpeg
}

func (e *FFmpeg) SetInput(input string) *FFmpeg {
	e.input = input
	return e
}

func (e *FFmpeg) SetTimeout(time int) *FFmpeg {
	e.timeout = time
	return e
}

func (e *FFmpeg) SetOutput(output string) *FFmpeg {
	e.output = output
	return e
}

func (e *FFmpeg) WithOptions(opts Options) *FFmpeg {
	e.options = [][]string{opts.Parame}
	return e
}

func (t *FFmpeg) GetMetadata(information *Information) error {

	if t.config.FfprobeBinPath == "" {
		t.config.FfprobeBinPath = "ffprobe"
	}
	var outb, errb bytes.Buffer

	input := t.input
	args := make([]string, 0)


	args = append(args, "-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error")

	information.Cmd = exec.Command(t.config.FfprobeBinPath, args...)
	information.Cmd.Stdout = &outb
	information.Cmd.Stderr = &errb

	err := information.Cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FfprobeBinPath, args, err, outb.String(), errb.String())
	}

	var metadata Metadata

	if err = json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
		return err
	}

	t.metadata = metadata

	return nil

}

func (t *FFmpeg) validate() error {

	if t.input == "" {
		return errors.New("missing input option")
	}
	if t.output == "" {
		return errors.New("missing input option")
	}

	return nil
}


func (t *FFmpeg)statrTimeOut()  {
	t.timer = time.NewTimer(1000 * time.Millisecond)

	go func() {
		for {
			select {
			case <-t.timer.C:
				timeout := time.Now().Unix() - t.setTime
				if int(timeout) > t.timeout {
					if t.information.Cmd !=nil {
						t.information.Cmd.Process.Kill()
					}
				}
				if t.timer !=nil {
					t.timer.Reset(1000 * time.Millisecond) // 每次使用完后需要人为重置下
				}
			}
		}

	}()
}

func (t *FFmpeg)stopTimeOut()  {
	if t.timer !=nil {
		t.timer.Stop()
		t.timer = nil
	}
}

func (t *FFmpeg) Run(information *Information) {

	if information == nil {
		information = &Information{}
	}

	t.setTime = time.Now().Unix()

	if t.config.ProgressEnabled && !t.config.Verbose {
		t.statrTimeOut()
	}

	t.information = information
	information.Error = make(chan error)
	information.Progress = make(chan Progress)
	options := t.options

	var stderrIn io.ReadCloser

	err := t.validate()
	if err != nil {
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	err = t.GetMetadata(information)
	t.setTime = time.Now().Unix()
	if err != nil {
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	// Append input file and standard options
	args := make([]string, 0)

	args = append(args, "-i", t.input)
	for _, v := range options {
		args = append(args, v...)
	}

	args = append(args, t.output)

	if t.config.FfmpegBinPath == "" {
		t.config.FfmpegBinPath = "ffmpeg"
	}
	information.Cmd = exec.Command(t.config.FfmpegBinPath, args...)

	if t.config.ProgressEnabled && !t.config.Verbose {
		stderrIn, err = information.Cmd.StderrPipe()
		if err != nil {
			go func(err error, information *Information) {
				information.Error <- err
			}(err, information)
			return
		}
	}

	if t.config.Verbose {
		information.Cmd.Stderr = os.Stdout
	}

	// Start process
	err = information.Cmd.Start()
	if err != nil {
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	if t.config.ProgressEnabled && !t.config.Verbose {
		go func() {
			t.progress(stderrIn, information.Progress)
		}()

		go func() {
			defer close(information.Progress)
			defer t.stopTimeOut()
			err = information.Cmd.Wait()
			if err != nil {
				go func(err error) {
					information.Error <- err
				}(err)
			} else {
				close(information.Error)
			}

		}()
	} else {
		err = information.Cmd.Wait()
		if err != nil {
			go func(err error) {
				information.Error <- err
			}(err)
		}
	}

}


// progress sends through given channel the transcoding status
func (t *FFmpeg) progress(stream io.ReadCloser, out chan Progress) {

	defer stream.Close()

	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(stream)
	scanner.Split(split)

	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		Progress := new(Progress)
		line := scanner.Text()

		if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
			var re = regexp.MustCompile(`=\s+`)
			st := re.ReplaceAllString(line, `=`)

			f := strings.Fields(st)

			var framesProcessed string
			var currentTime string
			var currentBitrate string
			var currentSpeed string

			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")

				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]

					if fieldname == "frame" {
						framesProcessed = fieldvalue
					}

					if fieldname == "time" {
						currentTime = fieldvalue
					}

					if fieldname == "bitrate" {
						currentBitrate = fieldvalue
					}
					if fieldname == "speed" {
						currentSpeed = fieldvalue
					}
				}
			}

			timesec := utils.DurToSec(currentTime)
			dursec, _ := strconv.ParseFloat(t.metadata.Format.Duration, 64)

			progress := (timesec * 100) / dursec
			Progress.Progress = progress

			Progress.CurrentBitrate = currentBitrate
			Progress.FramesProcessed = framesProcessed
			Progress.CurrentTime = currentTime
			Progress.Speed = currentSpeed
			ProtectRun(func() {
				t.setTime = time.Now().Unix()
				out <- *Progress
			})
		}
	}
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
		default:
		}
	}()
	entry()
}
