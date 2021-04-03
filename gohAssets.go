// gohAssets.go

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
	"embed"
	"log"
)

//go:embed assets/glade
//go:embed assets/images
var embeddedFiles embed.FS

// This functionality does not require explicit encoding of the files, at each
// compilation, the files are inserted into the resulting binary. Thus, updating
// assets is only required when new files are added to be embedded in order to
// create and declare the variables to which the files are linked.
// assetsDeclarationsUseEmbedded: Use native Go 'embed' package to include files
// content at runtime.
func assetsDeclarationsUseEmbedded(embedded ...bool) {
	mainGlade = readEmbedFile("assets/glade/main.glade")
	crossIcon48 = readEmbedFile("assets/images/Cross-icon-48.png")
	folder48 = readEmbedFile("assets/images/folder-48.png")
	linearProgressHorzBlue = readEmbedFile("assets/images/linear-progress-horz-blue.gif")
	logout48 = readEmbedFile("assets/images/logout-48.png")
	movieIcon = readEmbedFile("assets/images/movie-icon.png")
	options48 = readEmbedFile("assets/images/Options-48.png")
	stop48 = readEmbedFile("assets/images/Stop-48.png")
	tickIcon48 = readEmbedFile("assets/images/Tick-icon-48.png")
}

// readEmbedFile: read 'embed' file system and return []byte data.
func readEmbedFile(filename string) (out []byte) {
	var err error
	out, err = embeddedFiles.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read embedded file: %s, %v\n", filename, err)
	}
	return
}
