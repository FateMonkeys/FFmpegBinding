package FFmpegBinding

type Config struct {
	FfmpegBinPath   string //FFmpeg 路径
	FfprobeBinPath  string //FFprobe 路径
	ProgressEnabled bool
	Verbose         bool
}
