// gohImages.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 23:17:38 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/**********************************************************/
/* This section preserve user modifications on update.   */
/* Images declarations, used to initialize objects with */
/* The SetPict() func, accept both kind of variables:  */
/* filename or []byte content in case of using        */
/* embedded binary data. The variables names are the */
/* same. "assetsDeclarationsUseEmbedded(bool)" func */
/* could be used to toggle between filenames and   */
/* embedded binary type. See SetPict()            */
/* declaration to learn more on how to use it.   */
/************************************************/
func assignImages() {
	var iconsSize = 18
	SetPict(obj.ButtonExit, "")
	SetPict(obj.ButtonProceed, "")
	SetPict(obj.EditButtonClose, "")
	SetPict(obj.EditSpinAudioDelay, "")
	SetPict(obj.EditSpinAudioTrack, "")
	SetPict(obj.EditSpinCutSec, "")
	SetPict(obj.EditSpinCutSecDuration, "")
	SetPict(obj.EditSpinSplit, "")
	SetPict(obj.EditSpinTextTrack, "")
	SetPict(obj.EditWindow, "")
	SetPict(obj.InfosButtonClose, "")
	SetPict(obj.InfosButtonShowFilesList, "")
	SetPict(obj.MainToolButtonClear, "")
	SetPict(obj.MainToolButtonEdit, "")
	SetPict(obj.MainToolButtonInvertChecked, "")
	SetPict(obj.MainToolButtonUnckeckAll, "")
	SetPict(obj.MainWindow, movieIcon, iconsSize)
	SetPict(obj.WindowInfos, movieIcon, iconsSize)
}

/**********************************************************/
/* This section is rewritten on assets update.           */
/* Assets var declarations, this step permit to make a  */
/* bridge between the differents types used, string or */
/* []byte, and to simply switch from one to another.  */
/*****************************************************/
var mainGlade interface{}              // assets/glade/main.glade
var linearProgressHorzBlue interface{} // assets/images/linear-progress-horz-blue.gif
var movieIcon interface{}              // assets/images/movie-icon.png
