// gohObjects.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 05:54:34 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.0.5 github.com/hfmrow/info-media-mkv-ed
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
	ButtonExit                       *gtk.Button
	ButtonProceed                    *gtk.Button
	EditButtonClose                  *gtk.Button
	EditCheckAspectRatio             *gtk.CheckButton
	EditCheckAudioDelName            *gtk.CheckButton
	EditCheckAudioTrack              *gtk.CheckButton
	EditCheckAudioTrackDefault       *gtk.CheckButton
	EditCheckAudioTrackForced        *gtk.CheckButton
	EditCheckCumulativeDnD           *gtk.CheckButton
	EditCheckCut                     *gtk.CheckButton
	EditCheckGeneralPreserveTracks   *gtk.CheckButton
	EditCheckGeneralTitleCleanTags   *gtk.CheckButton
	EditCheckTextDelName             *gtk.CheckButton
	EditCheckTextTrack               *gtk.CheckButton
	EditCheckTextTrackDefault        *gtk.CheckButton
	EditCheckTextTrackForced         *gtk.CheckButton
	EditCheckVideoDelName            *gtk.CheckButton
	EditCutCheckOverwrite            *gtk.CheckButton
	EditCutCheckShowProgress         *gtk.CheckButton
	EditCutRadioFromEnd              *gtk.RadioButton
	EditCutRadioFromStart            *gtk.RadioButton
	EditCutSpinSec                   *gtk.SpinButton
	EditEntryAspectRatio             *gtk.Entry
	EditEntryGeneralPreserveTracks   *gtk.Entry
	EditEntryOutputSuffix            *gtk.Entry
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
	EditSpinAudioTrack               *gtk.SpinButton
	EditSpinTextTrack                *gtk.SpinButton
	EditStatusbar                    *gtk.Statusbar
	EditWindow                       *gtk.Window
	EditWindowEventboxResize         *gtk.EventBox
	GridProgress                     *gtk.Grid
	InfosButtonClose                 *gtk.Button
	InfosButtonShowEdit              *gtk.Button
	InfosCheckExpandAll              *gtk.CheckButton
	InfosHeaderLabel                 *gtk.Label
	InfosStatusbar                   *gtk.Statusbar
	InfosWindowEventboxResize        *gtk.EventBox
	LabelBitrate                     *gtk.Label
	LabelBitrateDisp                 *gtk.Label
	LabelCurrentFilenameDisp         *gtk.Label
	LabelFps                         *gtk.Label
	LabelFpsDisp                     *gtk.Label
	LabelFrame                       *gtk.Label
	LabelFrameDisp                   *gtk.Label
	LabelSpeed                       *gtk.Label
	LabelSpeedDisp                   *gtk.Label
	MainStatusbar                    *gtk.Statusbar
	MainToolbar                      *gtk.Toolbar
	MainToolButtonClear              *gtk.ToolButton
	MainToolButtonEdit               *gtk.ToolButton
	mainUiBuilder                    *gtk.Builder
	MainWindow                       *gtk.Window
	MainWindowEventboxMinimize       *gtk.EventBox
	MainWindowEventboxResize         *gtk.EventBox
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
	obj.ButtonExit = loadObject("ButtonExit").(*gtk.Button)
	obj.ButtonProceed = loadObject("ButtonProceed").(*gtk.Button)
	obj.EditButtonClose = loadObject("EditButtonClose").(*gtk.Button)
	obj.EditCheckAspectRatio = loadObject("EditCheckAspectRatio").(*gtk.CheckButton)
	obj.EditCheckAudioDelName = loadObject("EditCheckAudioDelName").(*gtk.CheckButton)
	obj.EditCheckAudioTrack = loadObject("EditCheckAudioTrack").(*gtk.CheckButton)
	obj.EditCheckAudioTrackDefault = loadObject("EditCheckAudioTrackDefault").(*gtk.CheckButton)
	obj.EditCheckAudioTrackForced = loadObject("EditCheckAudioTrackForced").(*gtk.CheckButton)
	obj.EditCheckCumulativeDnD = loadObject("EditCheckCumulativeDnD").(*gtk.CheckButton)
	obj.EditCheckCut = loadObject("EditCheckCut").(*gtk.CheckButton)
	obj.EditCheckGeneralPreserveTracks = loadObject("EditCheckGeneralPreserveTracks").(*gtk.CheckButton)
	obj.EditCheckGeneralTitleCleanTags = loadObject("EditCheckGeneralTitleCleanTags").(*gtk.CheckButton)
	obj.EditCheckTextDelName = loadObject("EditCheckTextDelName").(*gtk.CheckButton)
	obj.EditCheckTextTrack = loadObject("EditCheckTextTrack").(*gtk.CheckButton)
	obj.EditCheckTextTrackDefault = loadObject("EditCheckTextTrackDefault").(*gtk.CheckButton)
	obj.EditCheckTextTrackForced = loadObject("EditCheckTextTrackForced").(*gtk.CheckButton)
	obj.EditCheckVideoDelName = loadObject("EditCheckVideoDelName").(*gtk.CheckButton)
	obj.EditCutCheckOverwrite = loadObject("EditCutCheckOverwrite").(*gtk.CheckButton)
	obj.EditCutCheckShowProgress = loadObject("EditCutCheckShowProgress").(*gtk.CheckButton)
	obj.EditCutRadioFromEnd = loadObject("EditCutRadioFromEnd").(*gtk.RadioButton)
	obj.EditCutRadioFromStart = loadObject("EditCutRadioFromStart").(*gtk.RadioButton)
	obj.EditCutSpinSec = loadObject("EditCutSpinSec").(*gtk.SpinButton)
	obj.EditEntryAspectRatio = loadObject("EditEntryAspectRatio").(*gtk.Entry)
	obj.EditEntryGeneralPreserveTracks = loadObject("EditEntryGeneralPreserveTracks").(*gtk.Entry)
	obj.EditEntryOutputSuffix = loadObject("EditEntryOutputSuffix").(*gtk.Entry)
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
	obj.EditSpinAudioTrack = loadObject("EditSpinAudioTrack").(*gtk.SpinButton)
	obj.EditSpinTextTrack = loadObject("EditSpinTextTrack").(*gtk.SpinButton)
	obj.EditStatusbar = loadObject("EditStatusbar").(*gtk.Statusbar)
	obj.EditWindow = loadObject("EditWindow").(*gtk.Window)
	obj.EditWindowEventboxResize = loadObject("EditWindowEventboxResize").(*gtk.EventBox)
	obj.GridProgress = loadObject("GridProgress").(*gtk.Grid)
	obj.InfosButtonClose = loadObject("InfosButtonClose").(*gtk.Button)
	obj.InfosButtonShowEdit = loadObject("InfosButtonShowEdit").(*gtk.Button)
	obj.InfosCheckExpandAll = loadObject("InfosCheckExpandAll").(*gtk.CheckButton)
	obj.InfosHeaderLabel = loadObject("InfosHeaderLabel").(*gtk.Label)
	obj.InfosStatusbar = loadObject("InfosStatusbar").(*gtk.Statusbar)
	obj.InfosWindowEventboxResize = loadObject("InfosWindowEventboxResize").(*gtk.EventBox)
	obj.LabelBitrate = loadObject("LabelBitrate").(*gtk.Label)
	obj.LabelBitrateDisp = loadObject("LabelBitrateDisp").(*gtk.Label)
	obj.LabelCurrentFilenameDisp = loadObject("LabelCurrentFilenameDisp").(*gtk.Label)
	obj.LabelFps = loadObject("LabelFps").(*gtk.Label)
	obj.LabelFpsDisp = loadObject("LabelFpsDisp").(*gtk.Label)
	obj.LabelFrame = loadObject("LabelFrame").(*gtk.Label)
	obj.LabelFrameDisp = loadObject("LabelFrameDisp").(*gtk.Label)
	obj.LabelSpeed = loadObject("LabelSpeed").(*gtk.Label)
	obj.LabelSpeedDisp = loadObject("LabelSpeedDisp").(*gtk.Label)
	obj.MainStatusbar = loadObject("MainStatusbar").(*gtk.Statusbar)
	obj.MainToolbar = loadObject("MainToolbar").(*gtk.Toolbar)
	obj.MainToolButtonClear = loadObject("MainToolButtonClear").(*gtk.ToolButton)
	obj.MainToolButtonEdit = loadObject("MainToolButtonEdit").(*gtk.ToolButton)
	obj.MainWindow = loadObject("MainWindow").(*gtk.Window)
	obj.MainWindowEventboxMinimize = loadObject("MainWindowEventboxMinimize").(*gtk.EventBox)
	obj.MainWindowEventboxResize = loadObject("MainWindowEventboxResize").(*gtk.EventBox)
	obj.TreeViewFiles = loadObject("TreeViewFiles").(*gtk.TreeView)
	obj.TreeViewInfos = loadObject("TreeViewInfos").(*gtk.TreeView)
	obj.WindowInfos = loadObject("WindowInfos").(*gtk.Window)
}
