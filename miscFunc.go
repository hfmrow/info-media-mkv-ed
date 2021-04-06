// miscFunc.go

/*
	Source file auto-generated on Mon, 15 Mar 2021 03:37:06 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

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

	text, err := obj.EditEntryOutputSuffix.GetText()
	if err != nil {
		Logger.Log(err, "buildFileOut/EditEntryOutputSuffix", "'_out' as default value has been used")
		return "_out"
	}
	opt.OutputSuffix = text
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
	if !CheckCmd("mkvmerge") {
		missList = append(missList, "mkvmerge")
	}
	if len(missList) > 0 {
		return fmt.Errorf("Required packages: 'mkvtoolnix', 'ffmpeg', 'mediainfo'.\nSome shell command(s) missing: %s", strings.Join(missList, ", "))
	}
	return nil
}
