<p align="center">

<a href="https://www.youtube.com/watch?v=qtuod29DDrU"><img src="https://img.youtube.com/vi/qtuod29DDrU/0.jpg"></a>

</p>

# Intro

Convert your [MuseScore .mscz](https://github.com/musescore/musescore) into a string accepted by [Warframe's](https://warframe.com) 
Shawzin music generator.

# Features

- Load data from zipped (.mscz) or not (.mscx) [MuseScore](https://github.com/musescore/musescore) files.
- convert notes into Shawzin format

## Supported notations

Currently this project supports **only** notes and rests (no chords, or any special charactes, except dot)

## Shawzin options

Shawzin has several scales. Currently only chromatic scale is supported.
Range on this scale is from C3 to B3 (well, I managed to play C4 and D4 using chords so they are
also implemented) with all sharps `#`.

# Installation

`go install github.com/gucio321/shawziner/cmd/shawziner@master`

Or choose any other preffered method to run go programs.

# Usage

`shawziner -f <file.mscz>` or alternatively `shawziner -plain -f <file.mscx>`

# License and Legal notes

* Warframe and Shawzin are trademarks of Digital Extremes Ltd. This project is **not** affiliated with Digital Extremes Ltd.
* This project is licensed under (attached) [GPL v3.0 license](./LICENSE).
