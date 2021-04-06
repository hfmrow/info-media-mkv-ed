// cssBuilder.go

/*
	Source file auto-generated on Sat, 03 Apr 2021 18:48:50 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2021 hfmrow - Info Media mkv Ed v1.1 github.com/hfmrow/info-media-mkv-ed
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import "github.com/gotk3/gotk3/gdk"

func applyCss() (err error) {
	if err = mainWinDeco.CssApply(buildCss(
		mainWinDeco.DefaultCssName,
		opt.MainFgCol.ToGdkRGBA(),
		opt.SecondaryFgCol.ToGdkRGBA(),
		opt.MainBgCol.ToGdkRGBA(),
		opt.CellviewBgCol.ToGdkRGBA(),
		opt.ButtonFgCol.ToGdkRGBA(),
		opt.ButtonBgCol.ToGdkRGBA(),
		opt.SpinBgCol.ToGdkRGBA(),
		opt.ToolbarBgCol.ToGdkRGBA())); err == nil {

		if err = infosWinDeco.CssApply(buildCss(
			infosWinDeco.DefaultCssName,
			opt.MainFgCol.ToGdkRGBA(),
			opt.SecondaryFgCol.ToGdkRGBA(),
			opt.MainBgCol.ToGdkRGBA(),
			opt.CellviewBgCol.ToGdkRGBA(),
			opt.ButtonFgCol.ToGdkRGBA(),
			opt.ButtonBgCol.ToGdkRGBA(),
			opt.SpinBgCol.ToGdkRGBA(),
			opt.ToolbarBgCol.ToGdkRGBA())); err == nil {

			err = editWinDeco.CssApply(buildCss(
				editWinDeco.DefaultCssName,
				opt.MainFgCol.ToGdkRGBA(),
				opt.SecondaryFgCol.ToGdkRGBA(),
				opt.MainBgCol.ToGdkRGBA(),
				opt.CellviewBgCol.ToGdkRGBA(),
				opt.ButtonFgCol.ToGdkRGBA(),
				opt.ButtonBgCol.ToGdkRGBA(),
				opt.SpinBgCol.ToGdkRGBA(),
				opt.ToolbarBgCol.ToGdkRGBA()))
		}
	}
	return
}

func buildCss(cssName string, MainFgColor, SecondaryFgColor, MainBgColor, CellviewBgColor, ButtonFgColor, ButtonBgColor, SpinBgColor, ToolbarBgColor *gdk.RGBA) string {
	return `
@define-color MainFgColor ` + MainFgColor.String() + `;
@define-color SecondaryFgColor ` + SecondaryFgColor.String() + `;
@define-color MainBgColor ` + MainBgColor.String() + `;
@define-color CellviewBgColor ` + CellviewBgColor.String() + `;
@define-color ButtonFgColor ` + ButtonFgColor.String() + `;
@define-color ButtonBgColor ` + ButtonBgColor.String() + `;
@define-color SpinBgColor ` + SpinBgColor.String() + `;
@define-color ToolbarBgColor ` + ToolbarBgColor.String() + `;
/*
* {
	color: @MainFgColor;
}
*/
#` + cssName + ` {
	border-radius: 10px 10px 10px 10px;
/*	background-color: @MainBgColor;*/
}

treeview.view {
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.01,rgba(0,0,0,0.05)),
		color-stop(0.015,rgba(0,0,0,0.06)),
		color-stop(0.50,rgba(0,0,0,0.01)),
		color-stop(0.985,rgba(0,0,0,0.06)),
		color-stop(0.99,rgba(0,0,0,0.05)));
}

treeview.view:selected:focus, treeview.view:selected {
	color: @SecondaryFgColor;
	background-color: rgba(255, 255, 210, 0.9);
}
treeview.view:hover {
	background-color: rgba(255, 255, 230, 0.8);
}

/*
entry, text {
	border-radius: 10px 10px 10px 10px;
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.0,rgba(0,0,0,0.02)),
		color-stop(0.10,rgba(0,0,0,0.01)),
		color-stop(0.90,rgba(0,0,0,0.01)),
		color-stop(1.00,rgba(0,0,0,0.02)));
}
*/
entry {
	border-radius: 10px 10px 10px 10px;
	color: @MainFgColor;
	background-color: @MainBgColor;
}
/*
entry:disabled, text:disabled {
	border-radius: 10px 10px 10px 10px;
	color: rgba(0,0,0,0.5);
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.0,rgba(0,0,0,0.08)),
		color-stop(0.10,rgba(0,0,0,0.07)),
		color-stop(0.90,rgba(0,0,0,0.07)),
		color-stop(1.00,rgba(0,0,0,0.08)));
}
*/
toolbar {
	background-color: @ToolbarBgColor;
}

frame border {
	border-radius: 10px 10px 10px 10px;
}

combobox window menu {
	background-color: @MainBgColor;
}


combobox window menu cellview {
	border-radius: 15px 15px 15px 15px;
	background-color: @CellviewBgColor;
}

spinbutton entry {
	background-color: @SpinBgColor;
	border-radius: 10px 10px 10px 10px;
}

checkbutton check {
background-color: @MainBgColor;
}

button {
	border-radius: 10px 30px 10px 30px;
	color: @ButtonFgColor;
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.00,rgba(0,0,0,0.1)),
		color-stop(0.30,rgba(0,0,0,0.7)),
		color-stop(0.50,rgba(0,0,0,0.9)),
		color-stop(0.70,rgba(0,0,0,0.7)),
		color-stop(1.00,rgba(0,0,0,0.1)));
	background-color: @ButtonBgColor;
}

/* Button hovered */
button:hover {
	border-radius: 05px 15px 05px 15px;
	color: @ButtonFgColor;
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.00,rgba(0,0,0,0.0)),
		color-stop(0.30,rgba(0,0,0,0.6)),
		color-stop(0.50,rgba(0,0,0,0.8)),
		color-stop(0.70,rgba(0,0,0,0.6)),
		color-stop(1.00,rgba(0,0,0,0.0)));
}

/* Button clicked */
button:active {
	border-radius: 20px 20px 20px 20px;
	color: @ButtonFgColor;
	background-image: -gtk-gradient (linear,
		left top,
		left bottom,
		color-stop(0.00,rgba(0,0,0,0.0)),
		color-stop(0.30,rgba(0,0,0,0.5)),
		color-stop(0.50,rgba(0,0,0,0.7)),
		color-stop(0.70,rgba(0,0,0,0.5)),
		color-stop(1.00,rgba(0,0,0,0.0)));
}
`
}
