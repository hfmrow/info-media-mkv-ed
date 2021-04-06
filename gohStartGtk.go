// gohStartGtk.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 18:48:50 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

/*******************************/
/* Gtk3 Window Initialisation */
/*****************************/
func mainStartGtk(winTitle string, width, height int, center bool) {
	obj = new(MainControlsObj)
	gtk.Init(nil)
	if err := newBuilder(mainGlade); err == nil {

		/* Init tempDir and plan to delete it when leaving. */
		if doTempDir {
			tempDir = tempMake(Name)
			defer os.RemoveAll(tempDir)
		}

		/* Parse Gtk objects */
		gladeObjParser()

		/* Fill control with images */
		assignImages()

		/* Set Window Properties */
		// if center {
		// 	obj.MainWindow.SetPosition(gtk.WIN_POS_CENTER)
		// }
		obj.MainWindow.Connect("delete-event", func() { ButtonExitClicked(nil) })

		obj.MainWindow.SetTitle(winTitle)

		// obj.Progressbar.Hide()

		/* Start main application ... */
		mainApplication()

		/* Objects Signals initialisations */
		signalsPropHandler()

		/* Execute after signals initialisation */
		afterSignals()

		/* Start Gui loop */
		gtk.Main()
	} else {
		log.Fatal("Builder initialisation error.", err.Error())
	}
}

// windowDestroy: on closing/destroying the gui window.
func windowDestroy(window interface{}) {
	if onShutdown() {
		gtk.MainQuit()
	}
}
