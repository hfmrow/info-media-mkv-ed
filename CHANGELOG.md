## Informations

At the bottom you can find a compiled standalone ".deb" version with its checksum. The ".tar.gz" & ".zip" sources contain a "vendor" directory ensuring you can always compile it even if the official libraries have been changed.

## Changelog

All notable changes to this project will be documented in this file.

### [1.1] 2021-04-11

#### Added

- Add Drag and drop to info media window.
- Audio delay edition.
- Personal CSS semi-dark style with the possibility of deactivation (restart needed).
- Remux function. Sometimes useful when the '.mkv' file is not properly multiplexed and cannot be properly inspected.
- Remove entries: application, date information (depending on used format).
- Set/Remove aspect/ratio type tag.
- Explicitly allow or disallow editing actions to prevent unwanted behavior between them. This means that if there is an action that cannot be combined with others, only compatible actions will be available. 
- Deselect all and invert the selection in the file list window.
- File(s) count in status bar.

#### Fixed

- Issue with SpinButtons that it works alone running up / down without stopping.

#### Changed

- Some user interface changes.
- File size unit changed to be shorten and in lowercase.
- Progress information has been replaced with a progress bar.
- The 'format ',  'width x height',  'duration',  'size'  information in the file list are calculated and displayed (asynchronously) after all the files was shown instead of doing so when sending the file to the file list (useful on low frequency processors).

---

### [1.0.5] 2021-04-03

First public release
