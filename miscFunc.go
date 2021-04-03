// miscFunc.go

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
	"fmt"
	"os"
	"path/filepath"
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
func FFmpegCmdMaker(in, out string, mediainfo *gltsmo.MediaInfoStruct, overwrite bool,
	progress ...func(items map[string]string)) error {
	var (
		err    error
		ovrWrt = "-n"
		// Define default progress function
		callbackProgress = func(items map[string]string) {
			for n, v := range items {
				fmt.Println(n, v)
			}
			time.Sleep(time.Second)
		}
	)
	if len(progress) > 0 {
		callbackProgress = progress[0]
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

	if obj.EditCheckCut.GetActive() {
		if len(mediainfo.Media) == 0 {
			return fmt.Errorf("Cut cannot be processed, no media stream")
		} else if len(mediainfo.Media[0].Streams) == 0 {
			return fmt.Errorf("Cut cannot be processed, no media stream")
		}

		duration := toFloat(mediainfo.Media[0].Streams[0].Duration)
		toTrim := float64(obj.EditCutSpinSec.GetValue())

		if toTrim > duration {
			return fmt.Errorf("Cut cannot be processed, cut > video duration")
		}

		start := "0"
		end := time.Date(2006, 01, 02, 0, 0, int(duration-toTrim), 0, time.Local).Format("15:04:05")

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

	if obj.EditCheckAspectRatio.GetActive() {
		if obj.EditEntryAspectRatio.GetTextLength() > 0 {
			ar, err := obj.EditEntryAspectRatio.GetText()
			if err != nil {
				return fmt.Errorf("Wrong Aspect Ratio [%s] for command\n%t", ar, err)
			}
			cmd = append(cmd, []string{"-aspect", ar}...)
		}
	}

	cmd = append(cmd, []string{
		"-map",
		"0",
		"-c",
		"copy",
		out}...)

	err = ExecCommandProgress(cmd, 12, func(lines []string) {
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
	return err
}

// buildCmdFromCtrl: Build the command line with selected options
// and arguments, use 'mkvpropedit'
func buildCmdFromCtrl(filename, title string, mediainfo *gltsmo.MediaInfoStruct) (cmd []string, err error) {

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
	if obj.EditCheckGeneralTitleCleanTags.GetActive() {
		cmd = append(cmd, []string{"--tags", "all:"}...)
	}
	/*********
	 * Title
	 *********/
	switch {
	case obj.EditRadioGeneralTitleRemove.GetActive():
		cmd = append(cmd, []string{"--delete", "title"}...)
	case obj.EditRadioGeneralTitleChange.GetActive():
		/* DEBUG */
		// title = BaseNoExt(filename)
		/* DEBUG */
		if len(title) == 0 {
			return cmd, fmt.Errorf("Missing title for command")
		}
		cmd = append(cmd, []string{"--set", "title=" + title}...)
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

// setDefForced: Add 'Default' 'Forced' tag to commandline
func setDefForced(def, forced bool, tmpCmd []string) (out []string) {
	d, f := 0, 0
	if def {
		d = 1
	}
	tmpCmd = append(tmpCmd, []string{"--set", fmt.Sprintf("flag-default=%d", d)}...)

	if forced {
		f = 1
	}
	tmpCmd = append(tmpCmd, []string{"--set", fmt.Sprintf("flag-forced=%d", f)}...)

	return tmpCmd
}

// buildFileOut: Add suffix for output filename
func buildFileOut(file string) string {
	return filepath.Join(
		filepath.Dir(file),
		BaseNoExt(file)+opt.OutputSuffix+filepath.Ext(file))
}

// checkIsReady: Used to determine whether an option must be modified
// for a 'track'
func checkIsReady(nb int, list []int) bool {
	if len(list) == 0 {
		return true
	}
	return isExistSlInt(list, nb,
		obj.EditCheckGeneralPreserveTracks.GetActive())
}

// isExistSlInt: if exist then  ... 'reverse' return inverted result
func isExistSlInt(slice []int, item int, reverse bool) bool {
	for _, i := range slice {
		if i == item {
			return !reverse
		}
	}
	return reverse
}

// IsExistSlStr: if exist then  ...
func IsExistSlStr(slice []string, item string) bool {
	for _, mainRow := range slice {
		if mainRow == item {
			return true
		}
	}
	return false
}

func checkExistingCommand() error {
	var missList []string
	if !CheckCmd("ffmpeg") {
		missList = append(missList, "ffmpeg")
	}
	if !CheckCmd("ffprobe") {
		missList = append(missList, "ffprobe")
	}
	if !CheckCmd("mediainfo") {
		missList = append(missList, "mediainfo")
	}
	if !CheckCmd("mkvpropedit") {
		missList = append(missList, "mkvpropedit")
	}
	if len(missList) > 0 {
		return fmt.Errorf("Required packages: 'mkvtoolnix', 'ffmpeg', 'mediainfo'.\nSome shell command(s) missing: %s", strings.Join(missList, ", "))
	}
	return nil
}
