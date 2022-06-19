# Twitch2Davinci

This programm will convert a csv export of Twitch stream markers to Davinci Resolve markers for easier video editing.

[![Current status](https://github.com/suroh1994/twitch-marker-to-edl/workflows/CI/badge.svg)](https://github.com/suroh1994/twitch-marker-to-edl/actions?workflow=CI)
[![Releases](https://img.shields.io/github/release/suroh1994/twitch-marker-to-edl.svg)](https://github.com/suroh1994/twitch-marker-to-edl/releases)

## Install

**Binaries**

Find [the latest pre-compiled binaries here](https://github.com/suroh1994/twitch-marker-to-edl/releases/latest). 

**Source**

Alternatively you can grab the source code and compile it yourself like this:

```sh
$ go get -v github.com/suroh1994/twitch-marker-to-edl
```

## How to use

Just drag and drop the .csv file onto the binary and it will create a new .edl file with the same name in the same directory as the .csv file.

Alternatively you can invoke it from the CLI like this:

### Linux / macOS

```sh
Twitch2Davinci ./path/to/your/export.csv
```

### Windows

```bat
.\Twitch2Davinci.exe .\path\to\your\export.csv
```