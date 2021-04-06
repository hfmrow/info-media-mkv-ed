// edit.go

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
	"io/ioutil"
	"path/filepath"
	"strings"
)

// getFileTitle:
func getFileTitle() {
	// Retrieve selected filenames in treeview
	filesIn = filesIn[:0]
	list, err := tvsFilesIn.StoreToIfaceSlice()
	if err != nil {
		Logger.Log(err, "ButtonProceedClicked/StoreToIfaceSlice")
		DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
		return
	}
	for _, row := range list {
		if row[colsFilesMap["Toggle"]].(bool) {
			fileIn := filepath.Join(
				row[colsFilesMap["Path"]].(string),
				row[colsFilesMap["Name"]].(string))
			filesIn = append(filesIn, fileIn)
		}
	}

	// Titles generator
	if obj.EditRadioGeneralTitleChange.GetActive() {
		// Generate from filename
		if obj.EditRadioGeneralTitleUseFilename.GetActive() {
			for _, title := range filesIn {
				if len(title) > 0 {
					filesOut = append(filesOut, strings.TrimSpace(BaseNoExt(title)))
				}
			}
		} else {
			// Generate from text file
			var filename string
			var ok bool
			// glib.IdleAdd(func() {
			filename, ok, err = FileChooser(obj.MainWindow, "open-entry", "Choose titles text file", filepath.Dir(filesIn[0]))
			// })
			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/FileChooser")
				DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
				return
			}
			if ok {
				textFileBytes, err := ioutil.ReadFile(filename)
				if err != nil {
					Logger.Log(err, "ButtonProceedClicked/ReadFile")
					DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
					return
				}
				tmpStrSl := strings.Split(string(textFileBytes), GetTextEOL(textFileBytes))
				filesOut = filesOut[:0]
				for _, title := range tmpStrSl {
					if len(title) > 0 {
						filesOut = append(filesOut, strings.TrimSpace(title))
					}
				}
			}
		}
		if len(filesIn) != len(filesOut) {
			Logger.Log(err, "ButtonProceedClicked/Titles generator")
			DialogMessage(obj.MainWindow, "warning", "Warning",
				"\n"+`Title(s) error, please make sure you have the same number of title(s) as the file(s) count and try again.`,
				nil, "Continue")
			return
		}
	}
}

// goEdit:
func goEdit() {
	var formatErrorList,
		mkvmergeErrorList,
		ffmpegErrorList,
		propeditErrorList string

	var title string

	for idx, file := range filesIn {

		/**********************
		 * Tags modifications
		 **********************/
		mediainfo, err := MediaInfoStructNew(file)
		if obj.EditRadioGeneralTitleChange.GetActive() {
			title = filesOut[idx]
		}

		if err != nil {
			Logger.Log(err, "ButtonProceedClicked/MediaInfoStructNew")
			DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
			return
		}

		stream := mediainfo.Media[0].Streams[0]
		if !strings.Contains(stream.Format, "Matroska") {
			formatErrorList += fmt.Sprintf("%s: %s\n", stream.Format, TruncatePath(file, 2))
			continue
		}

		// Mkvpropedit
		cmd, err := buildPropeditCmd(file, title, mediainfo)
		if err != nil {
			Logger.Log(err, "ButtonProceedClicked/buildPropeditCmd")
			propeditErrorList += fmt.Sprintf("%v\n", err.Error())
			continue
		}

		// Execute command
		if len(cmd) > 2 {

			_, err := ExecCommand(cmd)

			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/buildPropeditCmd/ExecCommand")
				propeditErrorList += fmt.Sprintf("%v\n", err.Error())
				continue
			}
			// if err != nil {
			// 	fmt.Println(out)
			// 	Logger.Log(err, "ButtonProceedClicked/buildPropeditCmd/ExecCommand")
			// 	DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
			// 	return
			// }
		}

		// Remux only
		if obj.EditCheckGeneralRemux.GetActive() {

			err = FFmpegRemux(
				file,
				mediainfo,
				obj.EditCheckOverwrite.GetActive(), nil)

			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/buildMkvmergeCmd")
				ffmpegErrorList += fmt.Sprintf("%v\n", err.Error())
				continue
			}
		}

		/***************************
		 * Video Cut / AspectRatio
		 ***************************/
		if obj.EditCheckCut.GetActive() ||
			obj.EditCheckAspectRatio.GetActive() {

			err = FFmpegCmdMaker(
				file,
				mediainfo,
				obj.EditCheckOverwrite.GetActive(), nil)
			// func(items map[string]string) {
			// 	if obj.EditCheckShowProgress.GetActive() {
			// 		glib.TimeoutAdd(1000, func() bool {
			// 			obj.ProgressLabelFileDisp.SetLabel(TruncatePath(buildFileOut(file), 2))
			// 			obj.LabelFpsDisp.SetLabel(items["fps"])
			// 			obj.LabelBitrateDisp.SetLabel(items["bitrate"])
			// 			obj.LabelFrameDisp.SetLabel(items["frame"])
			// 			obj.LabelSpeedDisp.SetLabel(items["speed"])
			// 			return false
			// 		})
			// 	}
			// }

			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/buildMkvmergeCmd")
				ffmpegErrorList += fmt.Sprintf("%v\n", err.Error())
				continue
			}
		}

		// Mkvmerge
		if obj.EditCheckAudioDelay.GetActive() || obj.EditCheckSplit.GetActive() {
			cmd, err := buildMkvmergeCmd(file, mediainfo)
			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/buildMkvmergeCmd")
				mkvmergeErrorList += fmt.Sprintf("%v\n", err.Error())
				continue
			}
			// Execute command
			if len(cmd) > 2 {

				_, err := ExecCommand(cmd)

				if err != nil {
					// fmt.Println(out)
					Logger.Log(err, "ButtonProceedClicked/buildMkvmergeCmd/ExecCommand")
					DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
				}
			}
		}
	}

	if len(propeditErrorList) > 0 {
		formatErrorList += "\n<b>mkvpropedit edition failed:</b>\n" + propeditErrorList
	}

	if len(ffmpegErrorList) > 0 {
		formatErrorList += "\n<b>ffmpeg edition failed:</b>\n" + ffmpegErrorList
	}

	if len(mkvmergeErrorList) > 0 {
		formatErrorList += "\n<b>Mkvmerge edition failed:</b>\n" + mkvmergeErrorList
	}

	if len(formatErrorList) > 0 {
		formatErrorList = "\n<b>Editing capability is only for Matroska files, some are not.</b>\n" + formatErrorList
	}

	if len(formatErrorList) > 0 {
		DialogMessage(obj.MainWindow, "warningWithMarkup", "Warning",
			"\n"+formatErrorList,
			nil, "Continue")
	}
}
