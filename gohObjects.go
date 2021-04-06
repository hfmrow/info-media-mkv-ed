// gohObjects.go

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
	"github.com/gotk3/gotk3/gtk"
)

/* Control over all used objects from glade. */
var obj *MainControlsObj

/********************************************************/
/* This section preserve user modifications on update. */
/* Main structure Declaration: You may add your own   */
/* declarations (gotk3 objects only) here.           */
/****************************************************/
type MainControlsObj struct {
	BoxEdit                          *gtk.Box
	BoxMain                          *gtk.Box
	ButtonExit                       *gtk.Button
	ButtonProceed                    *gtk.Button
	EditButtonClose                  *gtk.Button
	EditCheckAspectRatio             *gtk.CheckButton
	EditCheckAudioDelay              *gtk.CheckButton
	EditCheckAudioDelName            *gtk.CheckButton
	EditCheckAudioTrack              *gtk.CheckButton
	EditCheckAudioTrackDefault       *gtk.CheckButton
	EditCheckAudioTrackForced        *gtk.CheckButton
	EditCheckCumulativeDnD           *gtk.CheckButton
	EditCheckCut                     *gtk.CheckButton
	EditCheckGeneralCleanTags        *gtk.CheckButton
	EditCheckGeneralPreserveTracks   *gtk.CheckButton
	EditCheckGeneralRemux            *gtk.CheckButton
	EditCheckOverwrite               *gtk.CheckButton
	EditCheckSemiDarkMode            *gtk.CheckButton
	EditCheckSplit                   *gtk.CheckButton
	EditCheckTextDelName             *gtk.CheckButton
	EditCheckTextTrack               *gtk.CheckButton
	EditCheckTextTrackDefault        *gtk.CheckButton
	EditCheckTextTrackForced         *gtk.CheckButton
	EditCheckVideoARType             *gtk.CheckButton
	EditCheckVideoDelApp             *gtk.CheckButton
	EditCheckVideoDelDate            *gtk.CheckButton
	EditCheckVideoDelLibrary         *gtk.CheckButton
	EditCheckVideoDelName            *gtk.CheckButton
	EditCutRadioExtract              *gtk.RadioButton
	EditCutRadioFromEnd              *gtk.RadioButton
	EditCutRadioFromStart            *gtk.RadioButton
	EditEntryAspectRatio             *gtk.Entry
	EditEntryGeneralPreserveTracks   *gtk.Entry
	EditEntryOutputSuffix            *gtk.Entry
	EditFrameARType                  *gtk.Frame
	EditFrameAspectRatio             *gtk.Frame
	EditFrameAudioTrack              *gtk.Frame
	EditFrameCut                     *gtk.Frame
	EditFrameDelay                   *gtk.Frame
	EditFrameSplit                   *gtk.Frame
	EditFrameText                    *gtk.Frame
	EditFrameTextTrack               *gtk.Frame
	EditFrameTitle                   *gtk.Frame
	EditGridTag                      *gtk.Grid
	EditGridVideo                    *gtk.Grid
	EditLabelAudio                   *gtk.Label
	EditLabelGeneral                 *gtk.Label
	EditLabelGeneralTitle            *gtk.Label
	EditLabelOutputSuffix            *gtk.Label
	EditLabelText                    *gtk.Label
	EditLabelVideo                   *gtk.Label
	EditRadioGeneralTitleChange      *gtk.RadioButton
	EditRadioGeneralTitleKeep        *gtk.RadioButton
	EditRadioGeneralTitleRemove      *gtk.RadioButton
	EditRadioGeneralTitleUseFilename *gtk.RadioButton
	EditRadioGeneralTitleUseTxtFile  *gtk.RadioButton
	EditRadioVideoARFixed            *gtk.RadioButton
	EditRadioVideoARFreeResize       *gtk.RadioButton
	EditRadioVideoARKeep             *gtk.RadioButton
	EditRadioVideoARRemove           *gtk.RadioButton
	EditSpinAudioDelay               *gtk.SpinButton
	EditSpinAudioTrack               *gtk.SpinButton
	EditSpinCutSec                   *gtk.SpinButton
	EditSpinCutSecDuration           *gtk.SpinButton
	EditSpinSplit                    *gtk.SpinButton
	EditSpinTextTrack                *gtk.SpinButton
	EditStatusbar                    *gtk.Statusbar
	EditWindow                       *gtk.Window
	EditWindowEventboxResize         *gtk.EventBox
	InfosButtonClose                 *gtk.Button
	InfosButtonShowFilesList         *gtk.Button
	InfosCheckExpandAll              *gtk.CheckButton
	InfosHeaderLabel                 *gtk.Label
	InfosStatusbar                   *gtk.Statusbar
	InfosWindowEventboxResize        *gtk.EventBox
	MainStatusbar                    *gtk.Statusbar
	MainToolbar                      *gtk.Toolbar
	MainToolButtonClear              *gtk.ToolButton
	MainToolButtonEdit               *gtk.ToolButton
	MainToolButtonInvertChecked      *gtk.ToolButton
	MainToolButtonUnckeckAll         *gtk.ToolButton
	mainUiBuilder                    *gtk.Builder
	MainWindow                       *gtk.Window
	MainWindowEventboxMinimize       *gtk.EventBox
	MainWindowEventboxResize         *gtk.EventBox
	OptionsFrameGlobal               *gtk.Frame
	OptionsLabelGlobal               *gtk.Label
	TreeViewFiles                    *gtk.TreeView
	TreeViewInfos                    *gtk.TreeView
	WindowInfos                      *gtk.Window
}

/******************************************************************/
/* This section preserve user modification on update.            */
/* GtkObjects initialisation: You may add your own declarations */
/* as you  wish, best way is to add them by grouping  same     */
/* objects names (below first declaration).                   */
/*************************************************************/
func gladeObjParser() {
	obj.BoxEdit = loadObject("BoxEdit").(*gtk.Box)
	obj.BoxMain = loadObject("BoxMain").(*gtk.Box)
	obj.ButtonExit = loadObject("ButtonExit").(*gtk.Button)
	obj.ButtonProceed = loadObject("ButtonProceed").(*gtk.Button)
	obj.EditButtonClose = loadObject("EditButtonClose").(*gtk.Button)
	obj.EditCheckAspectRatio = loadObject("EditCheckAspectRatio").(*gtk.CheckButton)
	obj.EditCheckAudioDelay = loadObject("EditCheckAudioDelay").(*gtk.CheckButton)
	obj.EditCheckAudioDelName = loadObject("EditCheckAudioDelName").(*gtk.CheckButton)
	obj.EditCheckAudioTrack = loadObject("EditCheckAudioTrack").(*gtk.CheckButton)
	obj.EditCheckAudioTrackDefault = loadObject("EditCheckAudioTrackDefault").(*gtk.CheckButton)
	obj.EditCheckAudioTrackForced = loadObject("EditCheckAudioTrackForced").(*gtk.CheckButton)
	obj.EditCheckCumulativeDnD = loadObject("EditCheckCumulativeDnD").(*gtk.CheckButton)
	obj.EditCheckCut = loadObject("EditCheckCut").(*gtk.CheckButton)
	obj.EditCheckGeneralCleanTags = loadObject("EditCheckGeneralCleanTags").(*gtk.CheckButton)
	obj.EditCheckGeneralPreserveTracks = loadObject("EditCheckGeneralPreserveTracks").(*gtk.CheckButton)
	obj.EditCheckGeneralRemux = loadObject("EditCheckGeneralRemux").(*gtk.CheckButton)
	obj.EditCheckOverwrite = loadObject("EditCheckOverwrite").(*gtk.CheckButton)
	obj.EditCheckSemiDarkMode = loadObject("EditCheckSemiDarkMode").(*gtk.CheckButton)
	obj.EditCheckSplit = loadObject("EditCheckSplit").(*gtk.CheckButton)
	obj.EditCheckTextDelName = loadObject("EditCheckTextDelName").(*gtk.CheckButton)
	obj.EditCheckTextTrack = loadObject("EditCheckTextTrack").(*gtk.CheckButton)
	obj.EditCheckTextTrackDefault = loadObject("EditCheckTextTrackDefault").(*gtk.CheckButton)
	obj.EditCheckTextTrackForced = loadObject("EditCheckTextTrackForced").(*gtk.CheckButton)
	obj.EditCheckVideoARType = loadObject("EditCheckVideoARType").(*gtk.CheckButton)
	obj.EditCheckVideoDelApp = loadObject("EditCheckVideoDelApp").(*gtk.CheckButton)
	obj.EditCheckVideoDelDate = loadObject("EditCheckVideoDelDate").(*gtk.CheckButton)
	obj.EditCheckVideoDelLibrary = loadObject("EditCheckVideoDelLibrary").(*gtk.CheckButton)
	obj.EditCheckVideoDelName = loadObject("EditCheckVideoDelName").(*gtk.CheckButton)
	obj.EditCutRadioExtract = loadObject("EditCutRadioExtract").(*gtk.RadioButton)
	obj.EditCutRadioFromEnd = loadObject("EditCutRadioFromEnd").(*gtk.RadioButton)
	obj.EditCutRadioFromStart = loadObject("EditCutRadioFromStart").(*gtk.RadioButton)
	obj.EditEntryAspectRatio = loadObject("EditEntryAspectRatio").(*gtk.Entry)
	obj.EditEntryGeneralPreserveTracks = loadObject("EditEntryGeneralPreserveTracks").(*gtk.Entry)
	obj.EditEntryOutputSuffix = loadObject("EditEntryOutputSuffix").(*gtk.Entry)
	obj.EditFrameARType = loadObject("EditFrameARType").(*gtk.Frame)
	obj.EditFrameAspectRatio = loadObject("EditFrameAspectRatio").(*gtk.Frame)
	obj.EditFrameAudioTrack = loadObject("EditFrameAudioTrack").(*gtk.Frame)
	obj.EditFrameCut = loadObject("EditFrameCut").(*gtk.Frame)
	obj.EditFrameDelay = loadObject("EditFrameDelay").(*gtk.Frame)
	obj.EditFrameSplit = loadObject("EditFrameSplit").(*gtk.Frame)
	obj.EditFrameText = loadObject("EditFrameText").(*gtk.Frame)
	obj.EditFrameTextTrack = loadObject("EditFrameTextTrack").(*gtk.Frame)
	obj.EditFrameTitle = loadObject("EditFrameTitle").(*gtk.Frame)
	obj.EditGridTag = loadObject("EditGridTag").(*gtk.Grid)
	obj.EditGridVideo = loadObject("EditGridVideo").(*gtk.Grid)
	obj.EditLabelAudio = loadObject("EditLabelAudio").(*gtk.Label)
	obj.EditLabelGeneral = loadObject("EditLabelGeneral").(*gtk.Label)
	obj.EditLabelGeneralTitle = loadObject("EditLabelGeneralTitle").(*gtk.Label)
	obj.EditLabelOutputSuffix = loadObject("EditLabelOutputSuffix").(*gtk.Label)
	obj.EditLabelText = loadObject("EditLabelText").(*gtk.Label)
	obj.EditLabelVideo = loadObject("EditLabelVideo").(*gtk.Label)
	obj.EditRadioGeneralTitleChange = loadObject("EditRadioGeneralTitleChange").(*gtk.RadioButton)
	obj.EditRadioGeneralTitleKeep = loadObject("EditRadioGeneralTitleKeep").(*gtk.RadioButton)
	obj.EditRadioGeneralTitleRemove = loadObject("EditRadioGeneralTitleRemove").(*gtk.RadioButton)
	obj.EditRadioGeneralTitleUseFilename = loadObject("EditRadioGeneralTitleUseFilename").(*gtk.RadioButton)
	obj.EditRadioGeneralTitleUseTxtFile = loadObject("EditRadioGeneralTitleUseTxtFile").(*gtk.RadioButton)
	obj.EditRadioVideoARFixed = loadObject("EditRadioVideoARFixed").(*gtk.RadioButton)
	obj.EditRadioVideoARFreeResize = loadObject("EditRadioVideoARFreeResize").(*gtk.RadioButton)
	obj.EditRadioVideoARKeep = loadObject("EditRadioVideoARKeep").(*gtk.RadioButton)
	obj.EditRadioVideoARRemove = loadObject("EditRadioVideoARRemove").(*gtk.RadioButton)
	obj.EditSpinAudioDelay = loadObject("EditSpinAudioDelay").(*gtk.SpinButton)
	obj.EditSpinAudioTrack = loadObject("EditSpinAudioTrack").(*gtk.SpinButton)
	obj.EditSpinCutSec = loadObject("EditSpinCutSec").(*gtk.SpinButton)
	obj.EditSpinCutSecDuration = loadObject("EditSpinCutSecDuration").(*gtk.SpinButton)
	obj.EditSpinSplit = loadObject("EditSpinSplit").(*gtk.SpinButton)
	obj.EditSpinTextTrack = loadObject("EditSpinTextTrack").(*gtk.SpinButton)
	obj.EditStatusbar = loadObject("EditStatusbar").(*gtk.Statusbar)
	obj.EditWindow = loadObject("EditWindow").(*gtk.Window)
	obj.EditWindowEventboxResize = loadObject("EditWindowEventboxResize").(*gtk.EventBox)
	obj.InfosButtonClose = loadObject("InfosButtonClose").(*gtk.Button)
	obj.InfosButtonShowFilesList = loadObject("InfosButtonShowFilesList").(*gtk.Button)
	obj.InfosCheckExpandAll = loadObject("InfosCheckExpandAll").(*gtk.CheckButton)
	obj.InfosHeaderLabel = loadObject("InfosHeaderLabel").(*gtk.Label)
	obj.InfosStatusbar = loadObject("InfosStatusbar").(*gtk.Statusbar)
	obj.InfosWindowEventboxResize = loadObject("InfosWindowEventboxResize").(*gtk.EventBox)
	obj.MainStatusbar = loadObject("MainStatusbar").(*gtk.Statusbar)
	obj.MainToolbar = loadObject("MainToolbar").(*gtk.Toolbar)
	obj.MainToolButtonClear = loadObject("MainToolButtonClear").(*gtk.ToolButton)
	obj.MainToolButtonEdit = loadObject("MainToolButtonEdit").(*gtk.ToolButton)
	obj.MainToolButtonInvertChecked = loadObject("MainToolButtonInvertChecked").(*gtk.ToolButton)
	obj.MainToolButtonUnckeckAll = loadObject("MainToolButtonUnckeckAll").(*gtk.ToolButton)
	obj.MainWindow = loadObject("MainWindow").(*gtk.Window)
	obj.MainWindowEventboxMinimize = loadObject("MainWindowEventboxMinimize").(*gtk.EventBox)
	obj.MainWindowEventboxResize = loadObject("MainWindowEventboxResize").(*gtk.EventBox)
	obj.OptionsFrameGlobal = loadObject("OptionsFrameGlobal").(*gtk.Frame)
	obj.OptionsLabelGlobal = loadObject("OptionsLabelGlobal").(*gtk.Label)
	obj.TreeViewFiles = loadObject("TreeViewFiles").(*gtk.TreeView)
	obj.TreeViewInfos = loadObject("TreeViewInfos").(*gtk.TreeView)
	obj.WindowInfos = loadObject("WindowInfos").(*gtk.Window)
}
