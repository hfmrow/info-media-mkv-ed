// treeView.go

/*
	Source file auto-generated on Sat, 27 Feb 2021 21:50:08 using Gotk3 Objects Handler v1.6.8 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - mkv-edit-gui v1.0 github.com/hfmrow/mkv-edit-gui
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	// gltsbh "github.com/hfmrow/genLib/tools/bench"

	"github.com/gotk3/gotk3/gtk"
)

/*
 * Files
 */

func treeViewFilesSetup() error {
	var err error

	tvsFilesIn, err = TreeViewStructureNew(obj.TreeViewFiles, false, false)
	if err != nil {
		return err
	}

	tvsFilesIn.AddColumns(tvColsFiles, false, true, true, true, true, true)

	tvsFilesIn.Columns[colsFilesMap["Toggle"]].ReadOnly = false
	tvsFilesIn.Columns[colsFilesMap["Toggle"]].Editable = true

	return tvsFilesIn.StoreSetup(new(gtk.ListStore))
}

func treeViewFilesPopulate(files ...[]string) error {

	var rows [][]interface{}

	filesToDisplay := filesIn
	if len(files) > 0 {
		filesToDisplay = files[0]
	}

	if !obj.EditCheckCumulativeDnD.GetActive() {
		tvsFilesIn.Clear()
		displayedFiles = displayedFiles[:0]
	}

	// b := gltsbh.BenchNew(true)
	// b.Start("start")
	for _, file := range filesToDisplay {

		if IsExistSlStr(displayedFiles, file) {
			continue
		}

		mediainfo, err := MediaInfoStructNew(file)
		if err != nil ||
			mediainfo.Media[0].AudioCount == 0 &&
				mediainfo.Media[0].VideoCount == 0 {
			continue
		}

		displayedFiles = append(displayedFiles, file)

		stream := mediainfo.Media[0].Streams[1]
		format := fmt.Sprintf("%s", strings.ToLower(stream.Format))
		wh := fmt.Sprintf("%sx%s", stream.Width, stream.Height)
		duration := strings.Split(toTime(stream.Duration), ".")
		size := toSize(mediainfo.Media[0].Streams[0].FileSize)
		time := "n/a"
		if len(duration) > 1 {
			time = duration[0]
		}

		rows = append(rows, tvsFilesIn.ColValuesStringSliceToIfaceSlice(
			filepath.Base(file),
			format,
			wh,
			time,
			size,
			filepath.Dir(file)))
	}

	if len(rows) > 0 {
		var iTrue []interface{}

		iTrue = append(iTrue, true)
		for _, row := range rows {
			row = append(iTrue, row...)
			tvsFilesIn.AddRow(nil, row...)
		}
	}
	// b.Stop()
	return nil
}

/*
 * Infos media
 */
// buildStreamInfos: create HR version ready to display with treestore
func treeViewInfosSetup() error {
	var err error

	tvsInfos, err = TreeViewStructureNew(obj.TreeViewInfos, false, false)
	if err != nil {
		return err
	}

	tvsInfos.AddColumns(tvColsInfos, false, true, true, true, true, true)
	// Hide 'toggle' column (we don't need it but it's necessary to build tree)
	tvsInfos.Columns[colsInfosMap["Toggle"]].Visible = false
	// Permit to select and copy 'Details' column
	tvsInfos.Columns[colsInfosMap["Details"]].Editable = true

	return tvsInfos.StoreSetup(new(gtk.TreeStore))
}

// Used to keep 1st 'details' in callback function
var tmpValue string

const strNA = "n/a"

func treeViewInfosPopulate(file string) error {
	var (
		err error

		addToTree = func(row []string) {
			iface := tvsInfos.ColValuesStringSliceToIfaceSlice(row...)

			tvsInfos.AddTree(colsInfosMap["Toggle"],
				colsInfosMap["Desc"],
				false,
				func(store *gtk.TreeStore, iter *gtk.TreeIter, currentAddIdx int, iRow *[]interface{}) bool {
					var value string

					row := *iRow
					// remove consumed value, this means that we add
					// manually the information without the help of
					// library method.
					if len(row) > 2 {
						*iRow = row[:2]
					}
					// This part control the informations displayed
					// inside treeviewinfo, when we have a category,
					// we don't care about details infos instead of
					// 'Desc' where we have to disp the full details.
					if currentAddIdx == 0 {
						tmpValue = row[2].(string)
						value = ""
					} else if len(row) == 2 {
						value = tmpValue
					} else {
						value = row[2].(string)
					}

					tvsInfos.SetColValue(iter, colsInfosMap["Details"], value)

					return true
				},
				iface...)
		}
	)

	tvsInfos.Clear()

	mediainfo, err = MediaInfoStructNew(file)
	if err != nil {
		Logger.Log(err, "treeViewInfosPopulate/MediaInfoStructNew")
	}
	// Display all 'streams'
	for _, media := range mediainfo.Media {
		for _, stream := range media.Streams {

			for _, row := range buildInfoRows(stream, file) {
				if row[2] == strNA {
					continue
				}
				row = muStyle("b", 0, muColor("#330000", 0, row))
				row = muStyle("i", 1, muColor("#110000", 1, row))
				row = muColor("#332244", 2, row)
				addToTree(row)
			}
		}
		// Display 'chapters' if they exists
		for _, entry := range media.Chapters.Entries {
			addToTree([]string{"Chapters", entry.Time, entry.Text})
		}
	}

	InfosCheckExpandAllToggled(obj.InfosCheckExpandAll)
	tvsInfos.TreeView.SetHeadersVisible(false)
	obj.WindowInfos.SetModal(false)
	obj.WindowInfos.SetKeepAbove(true)

	obj.WindowInfos.SetTitle("")
	obj.InfosHeaderLabel.SetLabel(TruncatePath(file, 2))
	obj.InfosButtonShowEdit.SetVisible(!obj.MainWindow.GetVisible())
	obj.WindowInfos.Show()

	obj.WindowInfos.Resize(opt.InfosWinWidth, opt.InfosWinHeight)
	obj.WindowInfos.Move(opt.InfosWinPosX, opt.InfosWinPosY)

	obj.TreeViewInfos.GrabFocus()

	return nil
}
