// gohOptions.go

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
	gipfmcrr "github.com/hfmrow/gtk3Import/pixbuff/misc/RGBA"
	gits "github.com/hfmrow/gtk3Import/tools"
	gitsww "github.com/hfmrow/gtk3Import/tools/window"
	gitw "github.com/hfmrow/gtk3Import/treeview"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Application infos. Only this section could be [modified during an update].
// Except for "Descr" variable, it is not recommended to manualy change
// options values relative to the project. Use GOH instead to doing this,
// or respect strictly the original applied format.
var (
	Name         = "Info Media mkv Ed"
	Vers         = "v1.1"
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
	DECO_AUTO_SHOW_HIDE       = gitsww.DECO_AUTO_SHOW_HIDE
	DECO_INIT_TRANSPARENT     = gitsww.DECO_INIT_TRANSPARENT
	mainWinDeco,
	infosWinDeco,
	editWinDeco *gitsww.WinDecorationStructure

	// Progressbar
	ProgressGifNew = gimc.ProgressGifNew
	pbs            *gimc.ProgressBarStruct
	// Statusbar
	StatusBarStructureNew = gimc.StatusBarStructureNew
	sbs                   *gimc.StatusBar

	// RGBA
	RgbaNew = gipfmcrr.RgbaNew

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
	DialogMessage = func(window *gtk.Window, dlgType, title, text string, iconFileName interface{}, buttons ...string) (value int) {
		glib.IdleAdd(func() {
			value = gidg.DialogMessage(window, dlgType, title, text, iconFileName, buttons...)
		})
		return
	}
	FileChooser = gidgcr.FileChooser /*func(window *gtk.Window, dlgType, title, filename string, options ...bool) (outFilename string, result bool, err error) {
		glib.IdleAdd(func() {
			outFilename, result, err = gidgcr.FileChooser(window, dlgType, title, filename, options...)
		})
		return
	}*/

	// D&D
	DragNDropNew = gimc.DragNDropNew
	DragNDropStruct,
	DragNDropInfoMedia *gimc.DragNDropStruct

	// Files
	TruncatePath      = glsg.TruncatePath
	GetTextEOL        = glsg.GetTextEOL
	HumanReadableSize = gltsushe.HumanReadableSize
	HR_UNIT_SHORTEN   = gltsushe.UNIT_SHORTEN
	HR_UNIT_DECIMAL   = gltsushe.UNIT_DECIMAL
	HR_UNIT_LOWER     = gltsushe.UNIT_LOWER
	HumanReadableTime = gltsushe.HumanReadableTime
	BaseNoExt         = glfs.BaseNoExt

	// Internal vars
	tvColsFiles  [][]string
	colsFilesMap = make(map[string]int)
	tvColsInfos  [][]string
	colsInfosMap = make(map[string]int)

	standAloneWindow bool

	filesCount int

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
	EditAudioDelay,
	EditTextTrack int

	InfosExpandAll,
	SemiDarkMode,
	EditOverwrite,
	CumulativeDnD,
	TitleTextFile,
	AlreadyStarted bool

	FilesDuration []time.Duration

	// RGBA
	MainFgCol,
	SecondaryFgCol,
	MainBgCol,
	CellviewBgCol,
	ButtonFgCol,
	ButtonBgCol,
	SpinBgCol,
	ToolbarBgCol *gipfmcrr.Rgba
}

// Init: Main options initialisation, Put here default values for your application.
func (opt *MainOpt) Init() {

	opt.MainWinWidth = 800
	opt.MainWinHeight = 400
	opt.InfosWinWidth = 470
	opt.InfosWinHeight = 700
	opt.EditWinWidth = 638
	opt.EditWinHeight = 537

	opt.OutputSuffix = "_out"
	opt.SemiDarkMode = true

	tvColsFiles = [][]string{
		{"", "active"},
		{"Name", "text"},
		{"Fmt", "text"},
		{"W x H", "text"},
		{"Durat.", "text"},
		{"Size", "text"},
		{"Path", "text"},
		// {"", "text"},
	}
	colsFilesMap = map[string]int{
		"Toggle":   0,
		"Name":     1,
		"Type":     2,
		"WxH":      3,
		"Duration": 4,
		"Size":     5,
		"Path":     6,
		"blank":    7,
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

	opt.MainFgCol = RgbaNew(255, 255, 255, 1)
	opt.SecondaryFgCol = RgbaNew(128, 32, 32, 0.9)
	opt.MainBgCol = RgbaNew(64, 64, 64, 0.8)
	opt.CellviewBgCol = RgbaNew(64, 32, 16, 0.3)
	opt.ButtonFgCol = opt.MainFgCol
	opt.ButtonBgCol = RgbaNew(0, 0, 0, 0)
	opt.SpinBgCol = RgbaNew(128, 128, 128, 0.8)
	opt.ToolbarBgCol = RgbaNew(64, 64, 64, 0)
}

func updWinPos(c int) {

	var updt = func() {

		obj.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)
		obj.WindowInfos.Resize(opt.InfosWinWidth, opt.InfosWinHeight)
		obj.EditWindow.Resize(opt.EditWinWidth, opt.EditWinHeight)

		if opt.MainWinPosX+opt.MainWinPosY == 0 {
			obj.MainWindow.SetDefaultSize(opt.MainWinWidth, opt.MainWinHeight)
			obj.MainWindow.SetPosition(gtk.WIN_POS_CENTER)
			// obj.MainWindow.Move(opt.MainWinPosX-(int(float64(opt.MainWinWidth)/2)), opt.MainWinPosY)
		} else {
			obj.MainWindow.Move(opt.MainWinPosX, opt.MainWinPosY)
		}
		if opt.InfosWinPosX+opt.InfosWinPosY == 0 {
			obj.WindowInfos.SetDefaultSize(opt.InfosWinWidth, opt.InfosWinHeight)
			obj.WindowInfos.SetPosition(gtk.WIN_POS_CENTER)
		} else {
			obj.WindowInfos.Move(opt.InfosWinPosX, opt.InfosWinPosY)
		}
		if opt.EditWinPosX+opt.EditWinPosY == 0 {
			obj.EditWindow.SetDefaultSize(opt.EditWinWidth, opt.EditWinHeight)
			obj.EditWindow.SetPosition(gtk.WIN_POS_CENTER)
		} else {
			obj.EditWindow.Move(opt.EditWinPosX, opt.EditWinPosY)
		}
	}
	count := c
	glib.TimeoutAdd(uint(64), func() bool {

		updt()
		count--
		return count > 0
	})
}

// UpdateObjects: Options -> Objects. Put here options to assign to gtk3 objects at start
func (opt *MainOpt) UpdateObjects() {

	// updWinPos(5)

	/* Your own declarations here */
	obj.InfosCheckExpandAll.SetActive(opt.InfosExpandAll)
	// InfosCheckExpandAllToggled(obj.InfosCheckExpandAll)

	obj.EditEntryOutputSuffix.SetText(opt.OutputSuffix)
	obj.EditCheckOverwrite.SetActive(opt.EditOverwrite)

	if opt.TitleTextFile {
		obj.EditRadioGeneralTitleUseTxtFile.SetActive(opt.TitleTextFile)
	} else {
		obj.EditRadioGeneralTitleUseFilename.SetActive(!opt.TitleTextFile)
	}
	obj.EditCheckCumulativeDnD.SetActive(opt.CumulativeDnD)
	obj.EditCheckSemiDarkMode.SetActive(opt.SemiDarkMode)
}

// UpdateOptions: Objects -> Options. Put here the gtk3 objects whose values you want to
// save in the options structure on exit.
func (opt *MainOpt) UpdateOptions() {

	// opt.MainWinWidth, opt.MainWinHeight = obj.MainWindow.GetSize()
	// opt.MainWinPosX, opt.MainWinPosY = obj.MainWindow.GetPosition()

	// opt.InfosWinWidth, opt.InfosWinHeight = obj.WindowInfos.GetSize()
	// opt.InfosWinPosX, opt.InfosWinPosY = obj.WindowInfos.GetPosition()

	// opt.EditWinWidth, opt.EditWinHeight = obj.EditWindow.GetSize()
	// opt.EditWinPosX, opt.EditWinPosY = obj.EditWindow.GetPosition()

	opt.TitleTextFile = obj.EditRadioGeneralTitleUseTxtFile.GetActive()
	opt.OutputSuffix, _ = obj.EditEntryOutputSuffix.GetText()
	opt.SemiDarkMode = obj.EditCheckSemiDarkMode.GetActive()
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
