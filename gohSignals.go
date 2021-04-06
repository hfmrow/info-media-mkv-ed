// gohSignals.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 23:17:38 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
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
	obj.BoxEdit.Connect("notify", blankNotify)
	obj.BoxMain.Connect("notify", blankNotify)
	obj.ButtonExit.Connect("clicked", ButtonExitClicked)
	obj.ButtonProceed.Connect("clicked", ButtonProceedClicked)
	obj.EditButtonClose.Connect("clicked", EditButtonCloseClicked)
	obj.EditCheckAspectRatio.Connect("toggled", switchRemux)
	obj.EditCheckAudioDelay.Connect("toggled", switchRemux)
	obj.EditCheckAudioDelName.Connect("toggled", switchProp)
	obj.EditCheckAudioTrack.Connect("toggled", switchProp)
	obj.EditCheckAudioTrackDefault.Connect("notify", blankNotify)
	obj.EditCheckAudioTrackForced.Connect("notify", blankNotify)
	obj.EditCheckCumulativeDnD.Connect("toggled", EditCheckCumulativeDnDToggled)
	obj.EditCheckCut.Connect("toggled", switchRemux)
	obj.EditCheckGeneralCleanTags.Connect("toggled", switchProp)
	obj.EditCheckGeneralPreserveTracks.Connect("notify", blankNotify)
	obj.EditCheckGeneralRemux.Connect("toggled", switchRemux)
	obj.EditCheckOverwrite.Connect("toggled", EditCheckOverwriteToggled)
	obj.EditCheckSemiDarkMode.Connect("toggled", EditCheckSemiDarkModeToggled)
	obj.EditCheckSplit.Connect("toggled", switchRemux)
	obj.EditCheckTextDelName.Connect("toggled", switchProp)
	obj.EditCheckTextTrack.Connect("toggled", switchProp)
	obj.EditCheckTextTrackDefault.Connect("notify", blankNotify)
	obj.EditCheckTextTrackForced.Connect("notify", blankNotify)
	obj.EditCheckVideoARType.Connect("toggled", switchProp)
	obj.EditCheckVideoDelApp.Connect("toggled", switchProp)
	obj.EditCheckVideoDelDate.Connect("toggled", switchProp)
	obj.EditCheckVideoDelLibrary.Connect("toggled", switchProp)
	obj.EditCheckVideoDelName.Connect("toggled", switchProp)
	obj.EditCutRadioExtract.Connect("notify", blankNotify)
	obj.EditCutRadioFromEnd.Connect("notify", blankNotify)
	obj.EditCutRadioFromStart.Connect("notify", blankNotify)
	obj.EditEntryAspectRatio.Connect("notify", blankNotify)
	obj.EditEntryGeneralPreserveTracks.Connect("notify", blankNotify)
	obj.EditEntryOutputSuffix.Connect("notify", blankNotify)
	obj.EditFrameARType.Connect("notify", blankNotify)
	obj.EditFrameAspectRatio.Connect("notify", blankNotify)
	obj.EditFrameAudioTrack.Connect("notify", blankNotify)
	obj.EditFrameCut.Connect("notify", blankNotify)
	obj.EditFrameDelay.Connect("notify", blankNotify)
	obj.EditFrameSplit.Connect("notify", blankNotify)
	obj.EditFrameText.Connect("notify", blankNotify)
	obj.EditFrameTextTrack.Connect("notify", blankNotify)
	obj.EditFrameTitle.Connect("notify", blankNotify)
	obj.EditGridTag.Connect("notify", blankNotify)
	obj.EditGridVideo.Connect("notify", blankNotify)
	obj.EditLabelAudio.Connect("notify", blankNotify)
	obj.EditLabelGeneral.Connect("notify", blankNotify)
	obj.EditLabelGeneralTitle.Connect("notify", blankNotify)
	obj.EditLabelOutputSuffix.Connect("notify", blankNotify)
	obj.EditLabelText.Connect("notify", blankNotify)
	obj.EditLabelVideo.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleChange.Connect("toggled", switchProp)
	obj.EditRadioGeneralTitleKeep.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleRemove.Connect("toggled", switchProp)
	obj.EditRadioGeneralTitleUseFilename.Connect("notify", blankNotify)
	obj.EditRadioGeneralTitleUseTxtFile.Connect("toggled", EditRadioGeneralTitleUseTxtFileToggled)
	obj.EditRadioVideoARFixed.Connect("notify", blankNotify)
	obj.EditRadioVideoARFreeResize.Connect("notify", blankNotify)
	obj.EditRadioVideoARKeep.Connect("notify", blankNotify)
	obj.EditRadioVideoARRemove.Connect("notify", blankNotify)
	obj.EditSpinAudioDelay.Connect("notify", blankNotify)
	obj.EditSpinAudioTrack.Connect("notify", blankNotify)
	obj.EditSpinCutSec.Connect("notify", blankNotify)
	obj.EditSpinCutSecDuration.Connect("notify", blankNotify)
	obj.EditSpinSplit.Connect("notify", blankNotify)
	obj.EditSpinTextTrack.Connect("notify", blankNotify)
	obj.EditStatusbar.Connect("notify", blankNotify)
	obj.EditWindow.Connect("notify", blankNotify)
	obj.EditWindowEventboxResize.Connect("notify", blankNotify)
	obj.InfosButtonClose.Connect("clicked", InfosButtonCloseClicked)
	obj.InfosButtonShowFilesList.Connect("clicked", InfosButtonShowFilesListClicked)
	obj.InfosCheckExpandAll.Connect("toggled", InfosCheckExpandAllToggled)
	obj.InfosHeaderLabel.Connect("notify", blankNotify)
	obj.InfosStatusbar.Connect("notify", blankNotify)
	obj.InfosWindowEventboxResize.Connect("notify", blankNotify)
	obj.MainStatusbar.Connect("notify", blankNotify)
	obj.MainToolbar.Connect("notify", blankNotify)
	obj.MainToolButtonClear.Connect("clicked", MainToolButtonClearClicked)
	obj.MainToolButtonEdit.Connect("clicked", MainToolButtonEditClicked)
	obj.MainToolButtonInvertChecked.Connect("clicked", func() { changeCheckState(true) })
	obj.MainToolButtonUnckeckAll.Connect("clicked", func() { changeCheckState(false) })
	obj.MainWindow.Connect("delete-event", windowDestroy)
	obj.MainWindowEventboxMinimize.Connect("notify", blankNotify)
	obj.MainWindowEventboxResize.Connect("notify", blankNotify)
	obj.OptionsFrameGlobal.Connect("notify", blankNotify)
	obj.OptionsLabelGlobal.Connect("notify", blankNotify)
	obj.TreeViewFiles.Connect("row-activated", TreeViewFilesActivated)
	obj.TreeViewInfos.Connect("notify", blankNotify)
	obj.WindowInfos.Connect("notify", blankNotify)
}
