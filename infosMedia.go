// infosMedia.go

/*
	Source file auto-generated on Mon, 15 Mar 2021 03:37:06 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - mkv-edit-gui v1.0 github.com/hfmrow/mkv-edit-gui
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"strconv"
	"strings"
	"time"

	gltsmo "github.com/hfmrow/genLib/tools/mmia/mediainfo"
)

func muStyle(style string, n int, in []string) []string {
	in[n] = "<" + style + ">" + in[n] + "</" + style + ">"
	return in
}

func muColor(color string, n int, in []string) []string {
	markUpStart := `<span foreground="`
	markUpMiddle := `">`
	markUpEnd := `</span>`
	in[n] = markUpStart + color + markUpMiddle + in[n] + markUpEnd
	return in
}

func buildInfoRows(stream gltsmo.Track, filename string) (rows [][]string) {
	var (
		getTypeWthID = func() string {
			return stream.Type + "(" + stream.InStreamID + ")"
		}
	)

	switch stream.Type {
	case "General":
		rows = append(rows, []string{stream.Type, "Title", toCtrledStr(stream.Movie)})
		rows = append(rows, []string{stream.Type, "UniqueID", toCtrledStr(stream.UniqueID)})
		rows = append(rows, []string{stream.Type, "Size", toSize(stream.FileSize)})
		rows = append(rows, []string{stream.Type, "Duration", toTime(stream.Duration)})
		rows = append(rows, []string{stream.Type, "Format", toCtrledStr(stream.Format)})
		rows = append(rows, []string{stream.Type, "Version", toCtrledStr(stream.FormatVersion)})
		rows = append(rows, []string{stream.Type, "Overall Bitrate", toSize(stream.OverallBitRate, "/s")})
		rows = append(rows, []string{stream.Type, "Application", toCtrledStr(stream.EncodedApplication)})
		rows = append(rows, []string{stream.Type, "Library", toCtrledStr(stream.EncodedLibrary)})
		rows = append(rows, []string{stream.Type, "Lib. version", toCtrledStr(stream.EncodedLibraryVersion)})
		rows = append(rows, []string{stream.Type, "Encoded date", toCtrledStr(stream.EncodedDate)})

	case "Video":
		HRAspectRation, err := getHRAspectRatio(filename)
		if err != nil {
			Logger.Log(err, "buildInfoRows/getHRAspectRatio")
			HRAspectRation = ""
		}
		rows = append(rows, []string{getTypeWthID(), "ID", toCtrledStr(stream.ID)})
		rows = append(rows, []string{getTypeWthID(), "Format", toCtrledStr(stream.Format)})
		rows = append(rows, []string{getTypeWthID(), "Profile", toCtrledStr(stream.FormatProfile)})
		rows = append(rows, []string{getTypeWthID(), "Level", toCtrledStr(stream.FormatLevel)})
		rows = append(rows, []string{getTypeWthID(), "Ref. frames", toCtrledStr(stream.FormatSettingsRefFrames)})
		rows = append(rows, []string{getTypeWthID(), "Codec", toCtrledStr(stream.CodecID)})
		rows = append(rows, []string{getTypeWthID(), "Size", toSize(stream.StreamSize)})
		rows = append(rows, []string{getTypeWthID(), "Duration", toTime(stream.Duration)})
		rows = append(rows, []string{getTypeWthID(), "Aspect ratio", toCtrledStr(HRAspectRation)})
		rows = append(rows, []string{getTypeWthID(), "Width x height", toCtrledStr(stream.Width + " x " + stream.Height)})
		rows = append(rows, []string{getTypeWthID(), "Bit depth", toCtrledStr(stream.BitDepth) + " bits"})
		rows = append(rows, []string{getTypeWthID(), "Bitrate", toSize(stream.BitRate, "/s")})
		rows = append(rows, []string{getTypeWthID(), "Fame rate mode", toCtrledStr(stream.FrameRateMode)})
		rows = append(rows, []string{getTypeWthID(), "Fame count", toCtrledStr(stream.FrameCount)})
		rows = append(rows, []string{getTypeWthID(), "Frame rate", toCtrledStr(stream.FrameRate) + " FPS"})
		rows = append(rows, []string{getTypeWthID(), "Color space", toCtrledStr(stream.ColorSpace)})
		rows = append(rows, []string{getTypeWthID(), "Chroma subsampling", toCtrledStr(stream.ChromaSubsampling)})
		rows = append(rows, []string{getTypeWthID(), "Language", toCtrledStr(stream.Language)})
		rows = append(rows, []string{getTypeWthID(), "Library", toCtrledStr(stream.EncodedLibrary)})
		rows = append(rows, []string{getTypeWthID(), "Settings", toCtrledStr(stream.EncodedLibrarySettings)})
		rows = append(rows, []string{getTypeWthID(), "Default", toCtrledStr(stream.Default)})
		rows = append(rows, []string{getTypeWthID(), "Forced", toCtrledStr(stream.Forced)})
		rows = append(rows, []string{getTypeWthID(), "Color range", toCtrledStr(stream.ColourRange)})
		rows = append(rows, []string{getTypeWthID(), "Color primaries", toCtrledStr(stream.ColourPrimaries)})
		rows = append(rows, []string{getTypeWthID(), "Transfer Charac.", toCtrledStr(stream.TransferCharacteristics)})
		rows = append(rows, []string{getTypeWthID(), "Matrix coeff.", toCtrledStr(stream.MatrixCoefficients)})

	case "Audio":
		rows = append(rows, []string{getTypeWthID(), "ID", toCtrledStr(stream.ID)})
		rows = append(rows, []string{getTypeWthID(), "Format", toCtrledStr(stream.Format)})
		rows = append(rows, []string{getTypeWthID(), "Codec", toCtrledStr(stream.CodecID)})
		rows = append(rows, []string{getTypeWthID(), "Codec name", toCtrledStr(stream.FormatCommercialIfAny)})
		rows = append(rows, []string{getTypeWthID(), "Size", toSize(stream.StreamSize)})
		rows = append(rows, []string{getTypeWthID(), "Duration", toTime(stream.Duration)})
		rows = append(rows, []string{getTypeWthID(), "Bitrate mode", toCtrledStr(stream.BitRateMode)})
		rows = append(rows, []string{getTypeWthID(), "Bitrate", toSize(stream.BitRate, "/s")})
		rows = append(rows, []string{getTypeWthID(), "Channels", toCtrledStr(stream.Channels)})
		rows = append(rows, []string{getTypeWthID(), "Channels layout", toCtrledStr(stream.ChannelLayout)})
		rows = append(rows, []string{getTypeWthID(), "Sampling rate", toCtrledStr(stream.SamplingRate)})
		rows = append(rows, []string{getTypeWthID(), "Frame rate",
			toCtrledStr(stream.FrameRate) + " FPS (SPF " + toCtrledStr(stream.SamplesPerFrame) + ")"})
		rows = append(rows, []string{getTypeWthID(), "Compression mode", toCtrledStr(stream.CompressionMode)})
		rows = append(rows, []string{getTypeWthID(), "Delay", toTime(stream.Delay)})
		rows = append(rows, []string{getTypeWthID(), "Language", toCtrledStr(stream.Language)})
		rows = append(rows, []string{getTypeWthID(), "Default", toCtrledStr(stream.Default)})
		rows = append(rows, []string{getTypeWthID(), "Forced", toCtrledStr(stream.Forced)})

	case "Text":
		rows = append(rows, []string{getTypeWthID(), "ID", toCtrledStr(stream.ID)})
		rows = append(rows, []string{getTypeWthID(), "Format", toCtrledStr(stream.Format)})
		rows = append(rows, []string{getTypeWthID(), "Codec", toCtrledStr(stream.CodecID)})
		rows = append(rows, []string{getTypeWthID(), "Size", toSize(stream.StreamSize)})
		rows = append(rows, []string{getTypeWthID(), "Duration", toTime(stream.Duration)})
		rows = append(rows, []string{getTypeWthID(), "Bitrate", toSize(stream.BitRate, "/s")})
		rows = append(rows, []string{getTypeWthID(), "Count of elements", toCtrledStr(stream.ElementCount)})
		rows = append(rows, []string{getTypeWthID(), "Language", toCtrledStr(stream.Language)})
		rows = append(rows, []string{getTypeWthID(), "Title", toCtrledStr(stream.Title)})
		rows = append(rows, []string{getTypeWthID(), "Default", toCtrledStr(stream.Default)})
		rows = append(rows, []string{getTypeWthID(), "Forced", toCtrledStr(stream.Forced)})

	default:
		rows = append(rows, []string{getTypeWthID(), "ID", toCtrledStr(stream.ID)})
		rows = append(rows, []string{getTypeWthID(), "Size", toSize(stream.StreamSize)})
		rows = append(rows, []string{getTypeWthID(), "Bitrate", toSize(stream.BitRate, "/s")})
		rows = append(rows, []string{getTypeWthID(), "Duration", toTime(stream.Duration)})
		rows = append(rows, []string{getTypeWthID(), "Codec", toCtrledStr(stream.CodecID)})
		rows = append(rows, []string{getTypeWthID(), "Compression mode", toCtrledStr(stream.CompressionMode)})
		rows = append(rows, []string{getTypeWthID(), "Default", toCtrledStr(stream.Default)})
		rows = append(rows, []string{getTypeWthID(), "Forced", toCtrledStr(stream.Forced)})
	}
	return rows
}

func toFloat(in string) float64 {
	if len(in) > 0 {
		val, err := strconv.ParseFloat(in, 64)
		if err == nil {
			return val
		}
	}
	return 0
}
func toTime(in string) string {
	if len(in) > 0 {
		val, _ := strconv.ParseFloat(in, 64)
		duration := time.Duration(val * float64(time.Second))
		if duration > 0 {
			return HumanReadableTime(duration)
		}
	}
	return strNA
}
func toSize(in string, suffix ...string) string {
	if len(in) > 0 {
		val, _ := strconv.ParseFloat(in, 64)
		if len(suffix) > 0 {
			return HumanReadableSize(val) + suffix[0]
		}
		return HumanReadableSize(val)
	}
	return strNA
}
func toCtrledStr(in string) string {
	if len(in) > 0 {
		return in
	}
	return strNA
}

// getHRAspectRatio: retrieve fractionned aspect ratio instead of float value
func getHRAspectRatio(filename string) (string, error) {
	cmd := []string{
		"ffprobe",
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=display_aspect_ratio",
		"-of",
		"csv=s=x:p=0",
		filename,
	}
	out, err := ExecCommand(cmd)
	if err != nil {
		Logger.Log(err, "ExecCommand")
		return out, err
	}
	tmpSl := strings.Split(out, "\n")
	if len(tmpSl) > 1 {
		out = tmpSl[0]
	} else {
		out = ""
	}

	return out, nil
}

// getFileCodec: retrieve file codec & aspec ratio
func getFileCodec(filename string) (string, error) {
	cmd := []string{
		"ffprobe",
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=codec_name,display_aspect_ratio",
		"-of",
		"csv=s=~:p=0",
		filename,
	}
	out, err := ExecCommand(cmd)
	if err != nil {
		Logger.Log(err, "ExecCommand")
		return out, err
	}
	tmpSl := strings.Split(out, "\n")
	if len(tmpSl) > 1 {
		out = tmpSl[0]
	} else {
		out = ""
	}

	return out, nil
}
