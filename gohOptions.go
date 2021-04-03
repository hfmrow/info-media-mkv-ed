// gohOptions.go

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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	glfs "github.com/hfmrow/genLib/files"
	glsg "github.com/hfmrow/genLib/strings"
	glts "github.com/hfmrow/genLib/tools"
	gltsle "github.com/hfmrow/genLib/tools/log2file"
	gltsushe "github.com/hfmrow/genLib/tools/units/human_readable"

	gltsmafg "github.com/hfmrow/genLib/tools/mmia/ffmpeg"
	gltsmo "github.com/hfmrow/genLib/tools/mmia/mediainfo"

	gidg "github.com/hfmrow/gtk3Import/dialog"
	gidgcr "github.com/hfmrow/gtk3Import/dialog/chooser"
	gimc "github.com/hfmrow/gtk3Import/misc"
	gits "github.com/hfmrow/gtk3Import/tools"
	gitsww "github.com/hfmrow/gtk3Import/tools/window"
	gitw "github.com/hfmrow/gtk3Import/treeview"

	"github.com/gotk3/gotk3/glib"
)

// Application infos. Only this section could be [modified during an update].
// Except for "Descr" variable, it is not recommended to manualy change
// options values relative to the project. Use GOH instead to doing this,
// or respect strictly the original applied format.
var (
	Name         = "Info Media mkv Ed"
	Vers         = "v1.0.5"
	Descr        = "Simple mkv info viewer with some limited editing features, titling, tag cleaner, default/forced track, head/tail video trimmer, aspect/ratio changer."
	Creat        = "hfmrow"
	YearCreat    = "2021"
	LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
	LicenseAbrv  = "License (MIT)"
	Repository   = "github.com/hfmrow/info-media-mkv-ed"

	// Internal var declarations
	opt *MainOpt

	devMode,
	VSCode,
	doTempDir bool

	absoluteRealPath,
	optFilename,
	tempDir string

	/*
	 * Lib mapping
	 */
	// Decoration
	WinDecorationStructureNew = gitsww.WinDecorationStructureNew
	mainWinDeco,
	infosWinDeco,
	editWinDeco *gitsww.WinDecorationStructure

	/* Lib mapping */
	// Command line
	ExecCommand         = glts.ExecCommand
	ExecCommandProgress = glts.ExecCommandProgress
	CheckCmd            = glts.CheckCmd

	// FFmpeg
	FFmpegMetadataNew = gltsmafg.FFmpegMetadataNew
	metadata          *gltsmafg.FFmpegMetadata

	// MediaInfo
	MediaInfoStructNew = gltsmo.MediaInfoStructNew
	mediainfo          *gltsmo.MediaInfoStruct

	// Errors handling
	Log2FileStructNew = gltsle.Log2FileStructNew
	Logger            *gltsle.Log2FileStruct

	// SpinButton / ScaleButton
	SpinScaleSetNew = gits.SpinScaleSetNew

	// TreeView
	TreeViewStructureNew = gitw.TreeViewStructureNew
	tvsFilesIn           *gitw.TreeViewStructure
	tvsInfos             *gitw.TreeViewStructure

	// Dialog
	DialogMessage = gidg.DialogMessage
	FileChooser   = gidgcr.FileChooser

	// D&D
	DragNDropNew    = gimc.DragNDropNew
	DragNDropStruct *gimc.DragNDropStruct

	// Files
	TruncatePath      = glsg.TruncatePath
	HumanReadableSize = gltsushe.HumanReadableSize
	HumanReadableTime = gltsushe.HumanReadableTime
	BaseNoExt         = glfs.BaseNoExt

	// Internal vars
	tvColsFiles  [][]string
	colsFilesMap = make(map[string]int)
	tvColsInfos  [][]string
	colsInfosMap = make(map[string]int)

	standAloneWindow bool

	filesIn,
	filesOut,
	displayedFiles []string
)

// MainOpt: This structure contains all the variables of the application, they
// will be saved when exiting and reloaded at launch.
type MainOpt struct {

	// File signature
	FileSign []string

	// Windows position
	MainWinWidth,
	MainWinHeight,
	MainWinPosX,
	MainWinPosY,

	InfosWinWidth,
	InfosWinHeight,
	InfosWinPosX,
	InfosWinPosY,

	EditWinWidth,
	EditWinHeight,
	EditWinPosX,
	EditWinPosY int

	LanguageFilename, // In case where GOTranslate is used.
	OutputSuffix string

	CutSec,
	EditAudioTrack,
	EditTextTrack int

	InfosExpandAll,
	EditOverwrite,
	EditShowProgress,
	CumulativeDnD,
	TitleTextFile bool

	FilesDuration []time.Duration
}

// Init: Main options initialisation, Put here default values for your application.
func (opt *MainOpt) Init() {

	opt.MainWinWidth = 800
	opt.MainWinHeight = 600
	opt.InfosWinWidth = 520
	opt.InfosWinHeight = 630

	opt.OutputSuffix = "_out"
	opt.EditShowProgress = false

	tvColsFiles = [][]string{
		{"", "active"},
		{"Name", "text"},
		{"Format", "text"},
		{"W x H", "text"},
		{"Durat.", "text"},
		{"Size", "text"},
		{"Path", "text"},
	}
	colsFilesMap = map[string]int{
		"Toggle":   0,
		"Name":     1,
		"Type":     2,
		"WxH":      3,
		"Duration": 4,
		"Size":     5,
		"Path":     6,
	}

	tvColsInfos = [][]string{
		{"Toggle", "active"},
		{"Desc", "markup"},
		{"Details", "markup"},
	}
	colsInfosMap = map[string]int{
		"Toggle":  0,
		"Desc":    1,
		"Details": 2,
	}
}

// UpdateObjects: Options -> Objects. Put here options to assign to gtk3 objects at start
func (opt *MainOpt) UpdateObjects() {

	var updWinPos = func() {

		obj.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)
		obj.MainWindow.Move(opt.MainWinPosX, opt.MainWinPosY)

		// obj.WindowInfos.Resize(opt.InfosWinWidth, opt.InfosWinHeight)
		// obj.WindowInfos.Move(opt.InfosWinPosX, opt.InfosWinPosY)

		// obj.EditWindow.Resize(opt.EditWinWidth, opt.EditWinHeight)
		// obj.EditWindow.Move(opt.EditWinPosX, opt.EditWinPosY)

		obj.GridProgress.SetVisible(opt.EditShowProgress)
	}

	count := 5
	glib.TimeoutAdd(uint(64), func() bool {

		updWinPos()
		count--
		return count > 0
	})

	/* Your own declarations here */
	obj.InfosCheckExpandAll.SetActive(opt.InfosExpandAll)
	// InfosCheckExpandAllToggled(obj.InfosCheckExpandAll)

	obj.GridProgress.SetVisible(false)
	obj.EditEntryOutputSuffix.SetText(opt.OutputSuffix)

	obj.EditCutCheckOverwrite.SetActive(opt.EditOverwrite)
	obj.EditCutCheckShowProgress.SetActive(opt.EditShowProgress)

	if opt.TitleTextFile {
		obj.EditRadioGeneralTitleUseTxtFile.SetActive(opt.TitleTextFile)
	} else {
		obj.EditRadioGeneralTitleUseFilename.SetActive(!opt.TitleTextFile)
	}
	obj.EditCheckCumulativeDnD.SetActive(opt.CumulativeDnD)
}

// UpdateOptions: Objects -> Options. Put here the gtk3 objects whose values you want to
// save in the options structure on exit.
func (opt *MainOpt) UpdateOptions() {

	opt.MainWinWidth, opt.MainWinHeight = obj.MainWindow.GetSize()
	opt.MainWinPosX, opt.MainWinPosY = obj.MainWindow.GetPosition()

	// opt.InfosWinWidth, opt.InfosWinHeight = obj.WindowInfos.GetSize()
	// opt.InfosWinPosX, opt.InfosWinPosY = obj.WindowInfos.GetPosition()

	// opt.EditWinWidth, opt.EditWinHeight = obj.EditWindow.GetSize()
	// opt.EditWinPosX, opt.EditWinPosY = obj.EditWindow.GetPosition()

	opt.TitleTextFile = obj.EditRadioGeneralTitleUseTxtFile.GetActive()
	opt.OutputSuffix, _ = obj.EditEntryOutputSuffix.GetText()
}

// Read: Options from file.
func (opt *MainOpt) Read() (err error) {
	opt.Init() // Init options with defaults values

	textFileBytes, err := ioutil.ReadFile(optFilename)
	if err != nil {
		return err
	}
	return json.Unmarshal(textFileBytes, &opt)
}

// Write: Options to file
func (opt *MainOpt) Write() error {
	var out bytes.Buffer
	opt.UpdateOptions()

	opt.FileSign = []string{Name, Vers, "©" + YearCreat, Creat, Repository, LicenseAbrv}
	jsonData, err := json.Marshal(&opt)
	if err != nil {
		return err
	} else if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
		return ioutil.WriteFile(optFilename, out.Bytes(), 0644)
	}
	return err
}
