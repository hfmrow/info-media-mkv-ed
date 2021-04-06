// editMedia.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 23:17:38 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	gltsmo "github.com/hfmrow/genLib/tools/mmia/mediainfo"
)

/*
 * NOT USED ANYMORE
 */
// parseCmd: parse command line with 'filenameIn/Out' and 'title'
func parseCmd(command, fileIn, fileOut, title string) (cmd []string, err error) {

	cmd = strings.Split(command, " ")
	if len(cmd) < 2 {
		return cmd, fmt.Errorf("Not anought arguments in command:\n%s", command)
	}

	title = BaseNoExt(fileIn)

	for idx, arg := range cmd {
		switch {
		case strings.Contains(arg, "[$IN$]"):
			cmd[idx] = strings.ReplaceAll(cmd[idx], "[$IN$]", fileIn)

		case strings.Contains(arg, "[$OUT$]"):
			cmd[idx] = strings.ReplaceAll(cmd[idx], "[$OUT$]", fileOut)

		case strings.Contains(arg, "[$TITLE$]"):
			cmd[idx] = strings.ReplaceAll(cmd[idx], "[$TITLE$]", title)
		}
	}

	return
}

/**/
func FFmpegCmdMaker(in string, mediainfo *gltsmo.MediaInfoStruct, overwrite bool,
	progress ...func(items map[string]string)) error {
	var (
		ovrWrt = "-n"
		// Define default progress function
		callbackProgress = func(items map[string]string) {
			for n, v := range items {
				fmt.Println(n, v)
			}
			time.Sleep(time.Second)
		}
	)

	if len(progress) > 0 && progress[0] != nil {
		callbackProgress = progress[0]
	}

	var execProgress = func(c []string) (err error) {
		err = ExecCommandProgress(c, 12, func(lines []string) {
			items := make(map[string]string)
			for _, line := range lines {
				splitted := strings.Split(line, "=")
				if len(splitted) > 1 {
					items[splitted[0]] = splitted[1]
				}
			}
			callbackProgress(items)
		})
		if err != nil {
			if strings.Contains(err.Error(), "No such file or directory") {
				err = os.ErrNotExist
			}
			if strings.Contains(err.Error(), "already exists. Exiting.") {
				err = os.ErrExist
			}
		}
		return
	}

	if overwrite {
		ovrWrt = "-y"
	}
	// build command
	cmd := []string{
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"error",
		"-progress",
		"-",
		"-nostats",
		ovrWrt}

	duration := toFloat(mediainfo.Media[0].Streams[0].Duration)
	toTrim := float64(obj.EditSpinCutSec.GetValue())
	to := float64(obj.EditSpinCutSecDuration.GetValue())
	start := "0"
	end := time.Date(2006, 01, 02, 0, 0, int(duration-toTrim), 0, time.Local).Format("15:04:05")

	// Cut Head/Tail
	if obj.EditCheckCut.GetActive() && !obj.EditCutRadioExtract.GetActive() {
		if len(mediainfo.Media) == 0 {
			return fmt.Errorf("Cut cannot be processed, no media stream")
		} else if len(mediainfo.Media[0].Streams) == 0 {
			return fmt.Errorf("Cut cannot be processed, no media stream")
		}
		if toTrim > duration {
			return fmt.Errorf("Cut cannot be processed, cut > video duration")
		}
		if !obj.EditCutRadioFromEnd.GetActive() {
			start = time.Date(2006, 01, 02, 0, 0, int(toTrim), 0, time.Local).Format("15:04:05")
		}
		cmd = append(cmd, []string{
			"-ss",
			start,
			"-i",
			in,
			"-t",
			end}...)
	} else {
		cmd = append(cmd, []string{
			"-i",
			in}...)
	}
	// Aspect/ratio
	if obj.EditCheckAspectRatio.GetActive() {
		if obj.EditEntryAspectRatio.GetTextLength() > 0 {
			ar, err := obj.EditEntryAspectRatio.GetText()
			if err != nil {
				return fmt.Errorf("Wrong Aspect Ratio [%s] for command\n%v", ar, err)
			}
			cmd = append(cmd, []string{"-aspect", ar}...)
		}
	}

	// Extract
	if obj.EditCheckCut.GetActive() && obj.EditCutRadioExtract.GetActive() {
		start = time.Date(2006, 01, 02, 0, 0, int(toTrim), 0, time.Local).Format("15:04:05")
		end = time.Date(2006, 01, 02, 0, 0, int(to), 0, time.Local).Format("15:04:05")
		if toTrim+to > duration {
			return fmt.Errorf("Cut cannot be processed, cut > video duration")
		}
		cmd = append(cmd, []string{
			"-map",
			"0",
			"-c",
			"copy",
			"-ss",
			start,
			"-t",
			end,
			buildFileOut(in)}...)
	} else {
		cmd = append(cmd, []string{
			"-map",
			"0",
			"-c",
			"copy",
			buildFileOut(in)}...)
	}
	return execProgress(cmd)
}

// FFmpegRemux: remux only
func FFmpegRemux(in string, mediainfo *gltsmo.MediaInfoStruct, overwrite bool,
	progress ...func(items map[string]string)) error {
	var (
		err    error
		ovrWrt = "-n"
		// Define default progress function
		// callbackProgress = func(items map[string]string) {
		// 	for n, v := range items {
		// 		fmt.Println(n, v)
		// 	}
		// 	time.Sleep(time.Second)
		// }
	)

	// if len(progress) > 0 && progress[0] != nil {
	// 	callbackProgress = progress[0]
	// }

	if overwrite {
		ovrWrt = "-y"
	}
	// build command
	cmd := []string{
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"error",
		"-progress",
		"-",
		"-nostats",
		ovrWrt,
		"-i",
		in,
		"-map",
		"0",
		"-c",
		"copy"}

	// does not work
	// if obj.EditCheckVideoDelSettings.GetActive() && mediainfo.Media[0].Streams[0].Format == "AVC" {
	// 	if !obj.EditCheckGeneralRemux.GetActive() {
	// 		cmd = append(cmd, []string{"-bsf:v", "'filter_units=remove_types=6'"}...)
	// 	}
	// }
	cmd = append(cmd, buildFileOut(in))

	_, err = ExecCommand(cmd)

	if err != nil {
		if strings.Contains(err.Error(), "No such file or directory") {
			err = os.ErrNotExist
		}
		if strings.Contains(err.Error(), "already exists. Exiting.") {
			err = os.ErrExist
		}
	}
	return err
}

func buildMkvmergeCmd(inFilename string, mediainfo *gltsmo.MediaInfoStruct) (cmd []string, err error) {

	var (
		absId = func(relId int, mi *gltsmo.MediaInfoStruct) int {
			for _, stream := range mi.Media[0].Streams {
				if stream.InStreamID == relId {
					return int(toFloat(stream.ID))
				}
			}
			return -1
		}
	)

	switch {
	case obj.EditCheckSplit.GetActive():
		cmd = append(cmd, []string{
			"mkvmerge",
			"--split",
			fmt.Sprintf("%dM", obj.EditSpinSplit.GetValueAsInt()),
			"-o",
			buildFileOut(inFilename),
			inFilename}...)

	case obj.EditCheckAudioDelay.GetActive():
		cmd = append(cmd, []string{
			"mkvmerge",
			"-o",
			buildFileOut(inFilename)}...)
		absAudioTrackId := absId(obj.EditSpinAudioTrack.GetValueAsInt(), mediainfo)
		if absAudioTrackId < 1 {
			return cmd, fmt.Errorf("[%s], Invalid Audio track ID: %d", TruncatePath(inFilename, 2), obj.EditSpinAudioTrack.GetValueAsInt())
		}
		cmd = append(cmd, []string{
			"--sync",
			fmt.Sprintf(
				"%d:%d",
				absAudioTrackId,
				obj.EditSpinAudioDelay.GetValueAsInt()),
			inFilename}...)
	}

	return
}

// buildPropeditCmd: Build the command line with selected options
// and arguments, use 'mkvpropedit'
func buildPropeditCmd(filename, title string, mediainfo *gltsmo.MediaInfoStruct) (cmd []string, err error) {

	var splittedInt []int
	var i int
	list, err := obj.EditEntryGeneralPreserveTracks.GetText()
	if err != nil {
		Logger.Log(err, "buildCmdFromCtrl/GetText")
		return
	}
	if len(list) > 0 {
		var reWhtSpce = regexp.MustCompile(`[\s]`)
		var reNonDigit = regexp.MustCompile(`[\D]`)
		list = reWhtSpce.ReplaceAllString(list, "")
		splittedStr := reNonDigit.Split(list, -1)
		for _, val := range splittedStr {
			if len(val) == 0 {
				continue
			}
			i, err = strconv.Atoi(val)
			if err != nil {
				return
			}
			splittedInt = append(splittedInt, i)
		}
	}

	cmd = append(cmd, []string{"mkvpropedit", filename}...)
	/**********
	 * Global
	 **********/
	if obj.EditCheckGeneralCleanTags.GetActive() {
		cmd = append(cmd, []string{"--tags", "all:"}...)
	}
	/*********
	 * Global
	 *********/
	switch {
	case obj.EditRadioGeneralTitleRemove.GetActive():
		cmd = append(cmd, []string{"--delete", "title"}...)
	case obj.EditRadioGeneralTitleChange.GetActive():
		if len(title) == 0 {
			return cmd, fmt.Errorf("Missing title for command")
		}
		cmd = append(cmd, []string{"--set", "title=" + title}...)
	}
	if obj.EditCheckVideoDelApp.GetActive() {
		cmd = append(cmd, []string{"--set", "writing-application=none"}...) // May be replaced with const 'n/a'
	}
	if obj.EditCheckVideoDelLibrary.GetActive() {
		cmd = append(cmd, []string{"--set", "muxing-application=none"}...)
	}
	if obj.EditCheckVideoDelDate.GetActive() {
		cmd = append(cmd, []string{"--delete", "date"}...)
	}

	vc := mediainfo.Media[0].VideoCount
	ac := mediainfo.Media[0].AudioCount
	tc := mediainfo.Media[0].TextCount
	streamsCount := 1

	var tmpCmd []string
	var defChg, forcedChg, globChg bool
	/*********
	 * Video
	 *********/
	for i := 1; i <= vc; i++ {
		tmpCmd = append(cmd, []string{"--edit", "track:" + fmt.Sprintf("v%d", i)}...)
		// Specific track modifications
		/*NOP*/
		// Global restricted modifications
		if checkIsReady(streamsCount, splittedInt) {
			if obj.EditCheckVideoDelName.GetActive() {
				globChg = true
				tmpCmd = append(tmpCmd, []string{"--delete", "name"}...)
			}
			if obj.EditCheckVideoARType.GetActive() {
				globChg = true
				if !obj.EditRadioVideoARRemove.GetActive() {
					value := 0
					switch {
					case obj.EditRadioVideoARKeep.GetActive():
						value = 1
					case obj.EditRadioVideoARFixed.GetActive():
						value = 2
					}
					tmpCmd = append(tmpCmd, []string{"--set", "aspect-ratio-type=" + fmt.Sprintf("%d", value)}...)
				} else {
					tmpCmd = append(tmpCmd, []string{"--delete", "aspect-ratio-type"}...)
				}
			}
		}
		if defChg || forcedChg || globChg {
			cmd = tmpCmd
		}
		streamsCount++
	}
	tmpCmd = tmpCmd[:0]
	/*********
	 * Audio
	 *********/
	for i := 1; i <= ac; i++ {
		tmpCmd = append(cmd, []string{"--edit", "track:" + fmt.Sprintf("a%d", i)}...)
		defChg, forcedChg, globChg = false, false, false

		// Specific track modifications
		if obj.EditCheckAudioTrack.GetActive() {

			// Default forced flags
			if i == obj.EditSpinAudioTrack.GetValueAsInt() {
				tmpCmd = setDefForced(
					obj.EditCheckAudioTrackDefault.GetActive(),
					obj.EditCheckAudioTrackForced.GetActive(),
					tmpCmd)
			} else {
				tmpCmd = setDefForced(defChg, forcedChg, tmpCmd)
			}
			defChg = true // flag to indicate modifications
		}
		// Global restricted modifications
		if checkIsReady(streamsCount, splittedInt) {
			if obj.EditCheckAudioDelName.GetActive() {
				globChg = true
				tmpCmd = append(tmpCmd, []string{"--delete", "name"}...)
			}
		}
		if defChg || forcedChg || globChg {
			cmd = tmpCmd
		}
		streamsCount++
	}
	tmpCmd = tmpCmd[:0]
	/********
	 * Text
	 ********/
	for i := 1; i <= tc; i++ {
		tmpCmd = append(cmd, []string{"--edit", "track:" + fmt.Sprintf("s%d", i)}...)
		defChg, forcedChg, globChg = false, false, false

		// Specific track modifications
		if obj.EditCheckTextTrack.GetActive() {

			if i == obj.EditSpinTextTrack.GetValueAsInt() {
				tmpCmd = setDefForced(
					obj.EditCheckTextTrackDefault.GetActive(),
					obj.EditCheckTextTrackForced.GetActive(),
					tmpCmd)
			} else {
				tmpCmd = setDefForced(defChg, forcedChg, tmpCmd)
			}
			defChg = true // flag to indicate modifications
		}
		// Global restricted modifications
		if checkIsReady(streamsCount, splittedInt) {
			if obj.EditCheckTextDelName.GetActive() {
				globChg = true
				tmpCmd = append(tmpCmd, []string{"--delete", "name"}...)
			}
		}
		if defChg || forcedChg || globChg {
			cmd = tmpCmd
		}
		streamsCount++
	}

	return
}
