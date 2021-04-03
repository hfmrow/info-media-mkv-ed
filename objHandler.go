// objHandler.go

/*
	Source file auto-generated on Wed, 24 Feb 2021 11:15:56 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - mediainfo wraper library github.com/hfmrow
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * Spin
 */
func SpinSecChanged(spin *gtk.SpinButton) {
	opt.CutSec = spin.GetValueAsInt()
}

/*
 * CheckButtons
 */
func InfosCheckExpandAllToggled(chk *gtk.CheckButton) {
	opt.InfosExpandAll = chk.GetActive()
	if opt.InfosExpandAll {
		obj.TreeViewInfos.ExpandAll()
	} else {
		obj.TreeViewInfos.CollapseAll()
	}
}

func EditCutCheckShowProgressToggled(chk *gtk.CheckButton) {
	opt.EditShowProgress = chk.GetActive()
	obj.GridProgress.SetVisible(opt.EditShowProgress)
}

func EditCutCheckOverwriteToggled(chk *gtk.CheckButton) {
	opt.EditOverwrite = chk.GetActive()
}

func EditRadioGeneralTitleUseTxtFileToggled(rdio *gtk.RadioButton) {
	opt.TitleTextFile = rdio.GetActive()
}

func EditCheckCumulativeDnDToggled(chk *gtk.CheckButton) {
	opt.CumulativeDnD = chk.GetActive()
}

/*
 * Buttons
 */
func ButtonExitClicked(btn *gtk.Button) {

	InfosButtonCloseClicked(nil)
	EditButtonCloseClicked(nil)
	windowDestroy(nil)
}

func EditButtonCloseClicked(btn *gtk.Button) {

	if obj.EditWindow.IsVisible() {
		opt.EditWinWidth, opt.EditWinHeight = obj.EditWindow.GetSize()
		opt.EditWinPosX, opt.EditWinPosY = obj.EditWindow.GetPosition()
	}
	obj.EditWindow.Hide()
}

func MainToolButtonClearClicked(tb *gtk.ToolButton) {
	tvsFilesIn.Clear()
	displayedFiles = displayedFiles[:0]
}

func MainToolButtonEditClicked(tb *gtk.ToolButton) {

	obj.EditWindow.SetModal(false)
	obj.EditWindow.SetKeepAbove(true)
	obj.EditWindow.SetTitle("")
	obj.EditWindow.Show()

	obj.EditWindow.Resize(opt.EditWinWidth, opt.EditWinHeight)
	obj.EditWindow.Move(opt.EditWinPosX, opt.EditWinPosY)
}

func InfosButtonCloseClicked(btn *gtk.Button) {

	if obj.WindowInfos.IsVisible() {
		opt.InfosWinWidth, opt.InfosWinHeight = obj.WindowInfos.GetSize()
		opt.InfosWinPosX, opt.InfosWinPosY = obj.WindowInfos.GetPosition()
	}

	opt.InfosExpandAll = obj.InfosCheckExpandAll.GetActive()

	if standAloneWindow {
		windowDestroy(nil)
	} else {
		obj.WindowInfos.Hide()
	}
}

func InfosButtonShowEditClicked(btn *gtk.Button) {

	obj.MainWindow.ShowAll()
	obj.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)
	obj.MainWindow.Move(opt.MainWinPosX, opt.MainWinPosY)

	standAloneWindow = false
	InfosButtonCloseClicked(nil)
}

func ButtonProceedClicked(btn *gtk.Button) {

	var formatErrorList string

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
			filename, ok, err := FileChooser(obj.MainWindow, "open-entry", "Choose titles text file", filepath.Dir(filesIn[0]))
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
				tmpStrSl := strings.Split(string(textFileBytes), "\n")
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
				"\n"+`Title(s) error, please make sure you have the same number of title(s) as the file(s) and try again.`,
				nil, "Continue")
			return
		}
	}
	var title string
	for idx, file := range filesIn {
		/**********************
		 * Tags modifications
		 **********************/
		mediainfo, err = MediaInfoStructNew(file)
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

		cmd, err := buildCmdFromCtrl(file, title, mediainfo)
		if err != nil {
			Logger.Log(err, "ButtonProceedClicked/buildCmdFromCtrl")
			DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
			return
		}
		// Execute command
		if len(cmd) > 2 {
			out, err := ExecCommand(cmd)
			if err != nil {
				fmt.Println(out)
				Logger.Log(err, "ButtonProceedClicked/ExecCommand")
				DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
				return
			}
		}

		/*************
		 * Video Cut
		 *************/
		if obj.EditCheckCut.GetActive() ||
			obj.EditCheckAspectRatio.GetActive() {
			// go func() {
			obj.GridProgress.SetVisible(obj.EditCutCheckShowProgress.GetActive())
			err = FFmpegCmdMaker(
				file,
				buildFileOut(file),
				mediainfo,
				obj.EditCutCheckOverwrite.GetActive(),
				func(items map[string]string) {
					if obj.EditCutCheckShowProgress.GetActive() {
						glib.TimeoutAdd(10, func() bool {
							obj.LabelCurrentFilenameDisp.SetLabel(TruncatePath(buildFileOut(file), 2))
							obj.LabelFpsDisp.SetLabel(items["fps"])
							obj.LabelBitrateDisp.SetLabel(items["bitrate"])
							obj.LabelFrameDisp.SetLabel(items["frame"])
							obj.LabelSpeedDisp.SetLabel(items["speed"])
							return false
						})
					}
				})
			if err != nil {
				Logger.Log(err, "ButtonProceedClicked/FFmpegCmdMaker")
				DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
				return
			}
			// }()
		}
	}
	if len(formatErrorList) > 0 {
		DialogMessage(obj.MainWindow, "warning", "Warning",
			"\nEditing capability is only for Matroska files, some are not.\n\n"+formatErrorList,
			nil, "Continue")
	}
}

/*
 * TreeView
 */
func TreeViewFilesActivated(tv *gtk.TreeView, path *gtk.TreePath, col *gtk.TreeViewColumn) {
	var err error
	var iter *gtk.TreeIter

	iter, err = tvsFilesIn.ListStore.GetIter(path)
	if err == nil {
		filename := filepath.Join(
			tvsFilesIn.GetColValue(iter, colsFilesMap["Path"]).(string),
			tvsFilesIn.GetColValue(iter, colsFilesMap["Name"]).(string))

		err = treeViewInfosPopulate(filename)

	}
	if err != nil {
		Logger.Log(err, "TreeViewFilesActivated/treeViewInfosPopulate")
		DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
	}
}
