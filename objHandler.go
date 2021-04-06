// objHandler.go

/*
	Source file auto-generated on Wed, 24 Feb 2021 11:15:56 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"log"
	"path/filepath"

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
 * Check/RadioButtons
 */
func InfosCheckExpandAllToggled(chk *gtk.CheckButton) {
	opt.InfosExpandAll = chk.GetActive()
	if opt.InfosExpandAll {
		obj.TreeViewInfos.ExpandAll()
	} else {
		obj.TreeViewInfos.CollapseAll()
	}
}

func EditCheckSemiDarkModeToggled(chk *gtk.CheckButton) {
	opt.SemiDarkMode = chk.GetActive()
}

func EditCheckOverwriteToggled(chk *gtk.CheckButton) {
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

	opt.MainWinWidth, opt.MainWinHeight = obj.MainWindow.GetSize()
	opt.MainWinPosX, opt.MainWinPosY = obj.MainWindow.GetPosition()

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

func MainToolButtonClearClicked(tb *gtk.ToolButton) {
	tvsFilesIn.Clear()
	displayedFiles = displayedFiles[:0]
}

func MainToolButtonEditClicked(tb *gtk.ToolButton) {

	obj.EditWindow.SetModal(false)
	obj.EditWindow.SetKeepAbove(true)
	obj.EditWindow.SetTitle("")
	obj.EditWindow.Show()

	updWinPos(5)
}

func InfosButtonShowFilesListClicked(btn *gtk.Button) {

	var err error
	if standAloneWindow {
		err = treeViewFilesPopulate()
		Logger.Log(err, "mainApplication/treeViewFilesPopulate")

		obj.MainWindow.Show()
		updWinPos(5)
	}
	standAloneWindow = false
	InfosButtonCloseClicked(nil)
}

func ButtonProceedClicked(btn *gtk.Button) {

	getFileTitle()

	anim, err := GetPixBufAnimation(linearProgressHorzBlue)
	if err != nil {
		log.Fatalf("GetPixBufAnimation: %s\n", err.Error())
	}
	gifImage, err := gtk.ImageNewFromAnimation(anim)
	if err != nil {
		log.Fatalf("ImageNewFromAnimation: %s\n", err.Error())
	}
	pbs = ProgressGifNew(gifImage, obj.BoxMain, 1,
		func() error {
			glib.IdleAdd(func() {
				obj.ButtonProceed.SetSensitive(false)
				obj.BoxEdit.SetSensitive(false)
			})
			goEdit()
			return nil
		},
		func() error {
			obj.ButtonProceed.SetSensitive(true)
			obj.BoxEdit.SetSensitive(true)
			return nil
		})

	go func() {
		pbs.StartGif()
	}()
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

func switchRemux(x interface{}) {
	if obj.EditCheckAspectRatio.GetActive() ||
		obj.EditCheckCut.GetActive() ||
		obj.EditCheckSplit.GetActive() ||
		obj.EditCheckGeneralRemux.GetActive() ||
		obj.EditCheckAudioDelay.GetActive() {

		obj.EditFrameTitle.SetSensitive(false)
		obj.EditGridTag.SetSensitive(false)
		obj.EditFrameText.SetSensitive(false)
		obj.EditFrameAudioTrack.SetSensitive(false)
		obj.EditCheckAudioDelName.SetSensitive(false)
		obj.EditFrameARType.SetSensitive(false)
		obj.EditCheckVideoDelName.SetSensitive(false)

	} else {
		obj.EditFrameTitle.SetSensitive(true)
		obj.EditGridTag.SetSensitive(true)
		obj.EditFrameText.SetSensitive(true)
		obj.EditFrameAudioTrack.SetSensitive(true)
		obj.EditCheckAudioDelName.SetSensitive(true)
		obj.EditFrameARType.SetSensitive(true)
		obj.EditCheckVideoDelName.SetSensitive(true)

	}
	var setSensitive = func(cut, remux, ar, delay, split bool) {
		obj.EditFrameCut.SetSensitive(cut)
		obj.EditCheckGeneralRemux.SetSensitive(remux)
		obj.EditFrameAspectRatio.SetSensitive(ar)
		obj.EditFrameDelay.SetSensitive(delay)
		obj.EditFrameSplit.SetSensitive(split)
	}
	setSensitive(true, true, true, true, true)
	if obj.EditCheckAudioDelay.GetActive() {
		setSensitive(false, false, false, true, false)
	}
	if obj.EditCheckCut.GetActive() ||
		obj.EditCheckAspectRatio.GetActive() {
		setSensitive(true, false, true, false, false)
	}
	if obj.EditCheckGeneralRemux.GetActive() {
		setSensitive(false, true, false, false, false)
	}
	if obj.EditCheckSplit.GetActive() {
		setSensitive(false, false, false, false, true)
	}
}

func switchProp(x interface{}) {

	if obj.EditCheckVideoDelLibrary.GetActive() ||
		obj.EditCheckVideoDelApp.GetActive() ||
		obj.EditCheckVideoDelDate.GetActive() ||
		obj.EditCheckVideoDelName.GetActive() ||
		obj.EditCheckVideoARType.GetActive() ||
		obj.EditCheckTextDelName.GetActive() ||
		obj.EditCheckAudioDelName.GetActive() ||
		obj.EditCheckAudioTrack.GetActive() ||
		obj.EditCheckTextTrack.GetActive() ||
		obj.EditRadioGeneralTitleRemove.GetActive() ||
		obj.EditRadioGeneralTitleChange.GetActive() ||
		obj.EditCheckGeneralCleanTags.GetActive() {

		obj.EditFrameCut.SetSensitive(false)
		obj.EditCheckGeneralRemux.SetSensitive(false)
		obj.EditFrameAspectRatio.SetSensitive(false)
		obj.EditFrameDelay.SetSensitive(false)
		obj.EditFrameSplit.SetSensitive(false)
	} else {
		obj.EditFrameCut.SetSensitive(true)
		obj.EditCheckGeneralRemux.SetSensitive(true)
		obj.EditFrameAspectRatio.SetSensitive(true)
		obj.EditFrameDelay.SetSensitive(true)
		obj.EditFrameSplit.SetSensitive(true)
	}
	obj.EditFrameTitle.SetSensitive(true)
	obj.EditGridTag.SetSensitive(true)
	obj.EditFrameText.SetSensitive(true)
	obj.EditFrameAudioTrack.SetSensitive(true)
	obj.EditCheckAudioDelName.SetSensitive(true)
	obj.EditFrameARType.SetSensitive(true)
}
