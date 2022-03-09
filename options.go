package FFmpegBinding

import (
	"fmt"
)

type Options struct {
	Parame []string
}


//自定义绑定参数
func (e *Options)WithCustomParame(parame string,value interface{})  {

	if _, ok := value.(bool); ok {
		e.Parame = append(e.Parame, parame)
	}

	if sv, ok := value.(string); ok {
		e.Parame = append(e.Parame,parame,  sv)
	}


	if va, ok := value.([]string); ok {
		e.Parame = append(e.Parame,parame)
		for i := 0; i < len(va); i++ {
			item := va[i]
			e.Parame = append(e.Parame,item)
		}
	}


	if vm, ok := value.(map[string]string); ok {
		for k, v := range vm {
			e.Parame = append(e.Parame, parame, fmt.Sprintf("%v:%v", k, v))
		}
	}

	if vi, ok := value.(*int); ok {
		e.Parame = append(e.Parame, parame, fmt.Sprintf("%d", *vi))
	}

}


func (e *Options)Aspect(value string)  {
	e.WithCustomParame("-aspect", value)
}
func (e *Options)Resolution(value string)  {
	e.WithCustomParame("-s", value)
}
func (e *Options)VideoBitRate(value string)  {
	e.WithCustomParame("-b:v", value)
}
func (e *Options)VideoBitRateTolerance(value int)  {
	e.WithCustomParame("-bt", value)
}
func (e *Options)VideoMaxBitRate(value int)  {
	e.WithCustomParame("-maxrate", value)
}
func (e *Options)VideoMinBitrate(value int)  {
	e.WithCustomParame("-minrate", value)
}
func (e *Options)VideoCodec(value string)  {
	e.WithCustomParame("-c:v", value)
}
func (e *Options)Vframes(value int)  {
	e.WithCustomParame("-vframes", value)
}
func (e *Options)FrameRate(value int)  {
	e.WithCustomParame("-r", value)
}
func (e *Options)AudioRate(value int)  {
	e.WithCustomParame("-ar", value)
}
func (e *Options)KeyframeInterval(value int)  {
	e.WithCustomParame("-g", value)
}
func (e *Options)AudioCodec(value string)  {
	e.WithCustomParame("-c:a", value)
}
func (e *Options)AudioBitrate(value string)  {
	e.WithCustomParame("-ab", value)
}
func (e *Options)AudioChannels(value int)  {
	e.WithCustomParame("-ac", value)
}
func (e *Options)AudioVariableBitrate(value bool)  {
	e.WithCustomParame("-q:a", value)
}
func (e *Options)BufferSize(value int)  {
	e.WithCustomParame("-bufsize", value)
}
func (e *Options)Threadset(value bool)  {
	e.WithCustomParame("-threads", value)
}
func (e *Options)Threads(value int)  {
	e.WithCustomParame("-threads", value)
}
func (e *Options)Preset(value string)  {
	e.WithCustomParame("-preset", value)
}
func (e *Options)Tune(value string)  {
	e.WithCustomParame("-tune", value)
}
func (e *Options)AudioProfile(value string)  {
	e.WithCustomParame("-profile:a", value)
}
func (e *Options)VideoProfile(value string)  {
	e.WithCustomParame("-profile:v", value)
}
func (e *Options)Target(value string)  {
	e.WithCustomParame("-target", value)
}
func (e *Options)Duration(value string)  {
	e.WithCustomParame("-t", value)
}
func (e *Options)Qscale(value uint32)  {
	e.WithCustomParame("-qscale", value)
}
func (e *Options)Crf(value uint32)  {
	e.WithCustomParame("-crf", value)
}
func (e *Options)Strict(value int)  {
	e.WithCustomParame("-strict", value)
}
func (e *Options)MuxDelay(value string)  {
	e.WithCustomParame("-muxdelay", value)
}
func (e *Options)SeekTime(value string)  {
	e.WithCustomParame("-ss", value)
}
func (e *Options)SeekUsingTimestamp(value bool)  {
	e.WithCustomParame("-seek_timestamp", value)
}
func (e *Options)MovFlags(value string)  {
	e.WithCustomParame("-movflags", value)
}
func (e *Options)HideBanner(value bool)  {
	e.WithCustomParame("-hide_banner", value)
}
func (e *Options)OutputFormat(value string)  {
	e.WithCustomParame("-f", value)
}
func (e *Options)CopyTs(value bool)  {
	e.WithCustomParame("-copyts", value)
}
func (e *Options)NativeFramerateInput(value bool)  {
	e.WithCustomParame("-re", value)
}
func (e *Options)InputInitialOffset(value string)  {
	e.WithCustomParame("-itsoffset", value)
}
func (e *Options)RtmpLive(value string)  {
	e.WithCustomParame("-rtmp_live", value)
}
func (e *Options)HlsPlaylistType(value string)  {
	e.WithCustomParame("-hls_playlist_type", value)
}
func (e *Options)HlsListSize(value int)  {
	e.WithCustomParame("-hls_list_size", value)
}
func (e *Options)HlsSegmentDuration(value int)  {
	e.WithCustomParame("-hls_time", value)
}
func (e *Options)HlsMasterPlaylistName(value string)  {
	e.WithCustomParame("-master_pl_name", value)
}
func (e *Options)HlsSegmentFilename(value string)  {
	e.WithCustomParame("-hls_segment_filename", value)
}
func (e *Options)HTTPMethod(value string)  {
	e.WithCustomParame("-method", value)
}
func (e *Options)HTTPKeepAlive(value bool)  {
	e.WithCustomParame("-multiple_requests", value)
}
func (e *Options)Hwaccel(value string)  {
	e.WithCustomParame("-hwaccel", value)
}
func (e *Options)StreamIds(value map[string]string)  {
	e.WithCustomParame("-streamid", value)
}
func (e *Options)VideoFilter(value string)  {
	e.WithCustomParame("-vf", value)
}
func (e *Options)AudioFilter(value string)  {
	e.WithCustomParame("-af", value)
}
func (e *Options)SkipVideo(value bool)  {
	e.WithCustomParame("-vn", value)
}
func (e *Options)SkipAudio(value bool)  {
	e.WithCustomParame("-an", value)
}
func (e *Options)CompressionLevel(value int)  {
	e.WithCustomParame("-compression_level", value)
}
func (e *Options)MapMetadata(value string)  {
	e.WithCustomParame("-map_metadata", value)
}
func (e *Options)Metadata(value map[string]string)  {
	e.WithCustomParame("-metadata", value)
}
func (e *Options)EncryptionKey(value string)  {
	e.WithCustomParame("-hls_key_info_file", value)
}
func (e *Options)Bframe(value int)  {
	e.WithCustomParame("-bf", value)
}
func (e *Options)PixFmt(value string)  {
	e.WithCustomParame("-pix_fmt", value)
}
func (e *Options)WhiteListProtocols(value []string)  {
	e.WithCustomParame("-protocol_whitelist", value)
}
func (e *Options)Overwrite(value bool)  {
	e.WithCustomParame("-y", value)
}
func (e *Options)CodecV(value string)  {
	e.WithCustomParame("-codec:v", value)
}
func (e *Options)CodecA(value string)  {
	e.WithCustomParame("-codec:a", value)
}
func (e *Options)Filter(value string)  {
	e.WithCustomParame("-filter:v", value)
}
func (e *Options)Lavfi(value string)  {
	e.WithCustomParame("-lavfi", value)
}
func (e *Options)I(value string)  {
	e.WithCustomParame("-i", value)
}
func (e *Options)ForceKeyFrames(value string)  {
	e.WithCustomParame("-force_key_frames", value)
}
func (e *Options)ExtraArgs(value map[string]interface {})  {
	e.WithCustomParame("", value)
}
