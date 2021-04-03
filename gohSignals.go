// gohSignals.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 05:54:34 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.0.5 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/********************************************************/
/* This section preserve user modifications on update. */
/* Signals & Property implementations:                */
/* initialize signals used by gtk objects ...        */
/****************************************************/
func signalsPropHandler() {
	obj.ButtonExit.Connect("clicked", ButtonExitClicked)
	obj.ButtonProceed.Connect("clicked", ButtonProceedClicked)
	obj.EditButtonClose.Connect("clicked", EditButtonCloseClicked)
	obj.EditCheckAspectRatio.Connect("notify", blankNotify)
	obj.EditCheckAudioDelName.Connect("notify", blankNotify)
	obj.EditCheckAudioTrack.Connect("notify", blankNotify)
	obj.EditCheckAudioTrackDefault.Connect("notify", blankNotify)
	obj.EditCheckAudioTrackForced.Connect("notify", blankNotify)
	obj.EditCheckCumulativeDnD.Connect("toggled", EditCheckCumulativeDnDToggled)
	obj.EditCheckCut.Connect("notify", blankNotify)
	obj.EditCheckGeneralPreserveTracks.Connect("notify", blankNotify)
	obj.EditCheckGeneralTitleCleanTags.Connect("notify", blankNotify)
	obj.EditCheckTextDelName.Connect("notify", blankNotify)
	obj.EditCheckTextTrack.Connect("notify", blankNotify)
	obj.EditCheckTextTrackDefault.Connect("notify", blankNotify)
	obj.EditCheckTextTrackForced.Connect("notify", blankNotify)
	obj.EditCheckVideoDelName.Connect("notify", blankNotify)
	obj.EditCutCheckOverwrite.Connect("toggled", EditCutCheckOverwriteToggled)
	obj.EditCutCheckShowProgress.Connect("toggled", EditCutCheckShowProgressToggled)
	obj.EditCutRadioFromEnd.Connect("notify", blankNotify)
	obj.EditCutRadioFromStart.Connect("notify", blankNotify)
	obj.EditCutSpinSec.Connect("changed", SpinSecChanged)
	obj.EditEntryAspectRatio.Connect("notify", blankNotify)
	obj.EditEntryGeneralPreserveTracks.Connect("notify", blankNotify)
	obj.EditEntryOutputSuffix.Connect("notify", blankNotify)
	obj.EditLabelAudio.Connect("notify", blankNotify)
	obj.EditLabelGeneral.Connect("notify", blankNotify)
	obj.EditLabelGeneralTitle.Connect("notify", blankNotify)
	obj.EditLabelOutputSuffix.Connect("notify", blankNotify)
	obj.EditLabelText.Connect("notify", blankNotify)
	obj.EditLabelVideo.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleChange.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleKeep.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleRemove.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleUseFilename.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleUseTxtFile.Connect("toggled", EditRadioGeneralTitleUseTxtFileToggled)
	obj.EditSpinAudioTrack.Connect("notify", blankNotify)
	obj.EditSpinTextTrack.Connect("notify", blankNotify)
	obj.EditStatusbar.Connect("notify", blankNotify)
	obj.EditWindow.Connect("notify", blankNotify)
	obj.EditWindowEventboxResize.Connect("notify", blankNotify)
	obj.GridProgress.Connect("notify", blankNotify)
	obj.InfosButtonClose.Connect("clicked", InfosButtonCloseClicked)
	obj.InfosButtonShowEdit.Connect("clicked", InfosButtonShowEditClicked)
	obj.InfosCheckExpandAll.Connect("toggled", InfosCheckExpandAllToggled)
	obj.InfosHeaderLabel.Connect("notify", blankNotify)
	obj.InfosStatusbar.Connect("notify", blankNotify)
	obj.InfosWindowEventboxResize.Connect("notify", blankNotify)
	obj.LabelBitrate.Connect("notify", blankNotify)
	obj.LabelBitrateDisp.Connect("notify", blankNotify)
	obj.LabelCurrentFilenameDisp.Connect("notify", blankNotify)
	obj.LabelFps.Connect("notify", blankNotify)
	obj.LabelFpsDisp.Connect("notify", blankNotify)
	obj.LabelFrame.Connect("notify", blankNotify)
	obj.LabelFrameDisp.Connect("notify", blankNotify)
	obj.LabelSpeed.Connect("notify", blankNotify)
	obj.LabelSpeedDisp.Connect("notify", blankNotify)
	obj.MainStatusbar.Connect("notify", blankNotify)
	obj.MainToolbar.Connect("notify", blankNotify)
	obj.MainToolButtonClear.Connect("clicked", MainToolButtonClearClicked)
	obj.MainToolButtonEdit.Connect("clicked", MainToolButtonEditClicked)
	obj.MainWindow.Connect("delete-event", windowDestroy)
	obj.MainWindowEventboxMinimize.Connect("notify", blankNotify)
	obj.MainWindowEventboxResize.Connect("notify", blankNotify)
	obj.TreeViewFiles.Connect("row-activated", TreeViewFilesActivated)
	obj.TreeViewInfos.Connect("notify", blankNotify)
	obj.WindowInfos.Connect("notify", blankNotify)
}
