// main.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 02:37:20 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.0.5 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/*
	Also used: Goffmpeg -> github.com/xfrr/goffmpeg/ffmpeg
	Author: Copyright (c) 2018 FlooStack
	MIT License

	Also used: go-mediainfo -> github.com/zelenin/go-mediainfo
	Author: Aleksandr Zelenin, e-mail: aleksandr@zelenin.me
	Unspecified license type.
*/

import (
	"fmt"
	"log"
	"os"
)

// main: And at the beginning ... this part is not modified on update.
// Build options informations:
// devMode: is used in some functions to control the behavior of the program
// When software is ready to be published, this flag must be set at "false"
// that mean:
// - options file will be stored in $HOME/.config/[Creat]/[softwareName],
// - translate function if used, will no more auto-update "sts" map sentences,
// - all built-in assets will be used instead of the files themselves.
//   Be aware to update assets via "Goh" and translations with "Got" before all.
func main() {

	// VSCode use a tmp directory to execute compiled code
	// so, we have to disable 'absoluteRealPath' at development
	devMode = true
	VSCode = false

	absoluteRealPath, optFilename = getAbsRealPath()

	/* Logger init. */
	Logger = Log2FileStructNew(optFilename, devMode)
	defer Logger.CloseLogger()

	// Initialization of assets according to the chosen mode (devMode).
	// you can set this flag to your liking without reference to devMode.
	assetsDeclarationsUseEmbedded(!devMode)

	// Create temp directory .. or not
	doTempDir = false

	/* Init & read options file */
	opt = new(MainOpt) // Assignate options' structure.
	opt.Read()         // Read values from options' file if exists.

	if devMode {
		os.Args = append(os.Args, []string{
			"/media/syndicate/storage/Vidéos-local/vid-cut/naming/WandaVision S01E01 - Filmé devant public (Filmed Before a Live Studio Audience).mkv",
			"/media/syndicate/storage/Vidéos-local/vid-cut/naming/WandaVision S01E02 - Ne zappez pas (Don't Touch That Dial).mkv",
			"/media/syndicate/storage/Vidéos-local/vid-cut/naming/WandaVision S01E03 - On passe à la couleur (Now in Color).mkv",
		}...)
	}

	if len(os.Args) > 1 {
		for _, file := range os.Args[1:] {
			filesIn = append(filesIn, file)
		}
	}

	standAloneWindow = len(filesIn) == 1

	/* Init gtk display */
	mainStartGtk(fmt.Sprintf("%s %s  %s %s %s",
		Name,
		Vers,
		"©"+YearCreat,
		Creat,
		LicenseAbrv),
		opt.MainWinWidth,
		opt.MainWinHeight, true)
}

/*******************************************\
/* Executed before signals initialisation. */
/******************************************/
func mainApplication() {
	var err error

	/* Ckeck required packages */
	err = checkExistingCommand()
	if err != nil {
		DialogMessage(obj.MainWindow, "warning", "Warning", "\n"+err.Error(), nil, "Continue")
		os.Exit(1)
	}

	/* Init windows decoration */
	if mainWinDeco, err = WinDecorationStructureNew(
		obj.MainWindow,
		obj.MainWindowEventboxResize,
		obj.MainWindowEventboxMinimize,
		nil); err == nil { // Deactivate 'Displacement' when hovering 'TreeViewFiles'
		mainWinDeco.SignalHandleBlockUnblock(obj.TreeViewFiles.ToWidget(), nil, nil)
	}
	Logger.Log(err, "mainApplication/WinDecorationStructureNew/MainWindow")

	infosWinDeco, err = WinDecorationStructureNew(
		obj.WindowInfos,
		obj.InfosWindowEventboxResize,
		nil,
		nil)
	Logger.Log(err, "mainApplication/WinDecorationStructureNew/WindowInfos")

	editWinDeco, err = WinDecorationStructureNew(
		obj.EditWindow,
		obj.EditWindowEventboxResize,
		nil,
		nil)
	Logger.Log(err, "mainApplication/WinDecorationStructureNew/EditWindow")

	/* Init Spinbuttons */
	spinSec, err := SpinScaleSetNew(obj.EditCutSpinSec, 0, 18000, float64(opt.CutSec), 1, nil)
	Logger.Log(err, "mainApplication/SpinScaleSetNew/spinSec")
	spinSec.SetDigits(0)
	spinAudio, err := SpinScaleSetNew(obj.EditSpinAudioTrack, 0, 64, float64(opt.EditAudioTrack), 1, nil)
	Logger.Log(err, "mainApplication/SpinScaleSetNew/spinAudio")
	spinAudio.SetDigits(0)
	spinText, err := SpinScaleSetNew(obj.EditSpinTextTrack, 0, 64, float64(opt.EditTextTrack), 1, nil)
	Logger.Log(err, "mainApplication/SpinScaleSetNew/spinText")
	spinText.SetDigits(0)

	/* Init DND */
	DragNDropStruct = DragNDropNew(obj.TreeViewFiles, &filesIn,
		func() {
			err = treeViewFilesPopulate()
			Logger.Log(err, "DragNDrop/callback")
		})

	/* Init Gtk3 objects content */
	opt.UpdateObjects()

	/* Init treeviews */
	if err = treeViewFilesSetup(); err == nil {
		if err = treeViewInfosSetup(); err == nil {
			/* If there is only one file in the list, only show infos media */
			err = treeViewFilesPopulate()
			Logger.Log(err, "afterSignals/treeViewFilesPopulate")
			if standAloneWindow {

				obj.WindowInfos.Connect("delete-event", windowDestroy)
				err = treeViewInfosPopulate(filesIn[0])
				Logger.Log(err, "afterSignals/treeViewInfosPopulate")
			}
		}
	}
	Logger.Log(err, "mainApplication/Init treeviews")
}

/******************************************\
/* Executed after signals initialisation. */
/*****************************************/
func afterSignals() {}

/*************************************\
/* Executed just before closing all. */
/************************************/
func onShutdown() bool {
	var err error

	// Update 'opt' with GtkObjects and save it
	if err = opt.Write(); err == nil {
		// What you want to execute before closing the app.
		// Return:
		// 	true for exit applicaton
		//	false does not exit application
	}
	if err != nil {
		log.Fatalf("Unexpected error on exit: %v", err)
	}
	return true
}
