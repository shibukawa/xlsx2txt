xlsx2txt
=================

Git diff wrapper for xlsx file. This program is designed for xlsx files for data table.

This program is inspired by https://raw.githubusercontent.com/yappo/p5-git-xlsx-textconv.pl.

# Install

Put `xlsx2txt` any folder in PATH.

# How to use

It basically assuems spread sheet like:

* Header names are in first row or second row
* First column stores id of data in row
* Worksheet name doesn't start with `'_'`.
Column header row is decided by `-c` option (default: 1).
Some case, first row contains flag (only not empty columns are processed) and headers are in second row, you can use `-f 1 -c 2`.
Rows before column header row and flag row are ignored.

You can cnage id column by using `-i 2`. In this case left cells from id cell are ignored.

If you want to ignore worksheet whose name starts with `Test`, you can pass `-s Test` as an option.

# ~/.gitconfig or $GIT_DIR/config

    [diff "xlsx"]
        binary = true
        textconv = $PATH/xlsx2txt

# .gitattributes in your project repository

    *.xlsx diff=xlsx

# build

It needs golang compiler.

    $ make get    # it gets dependency modules from github
    $ make build  # xlsx2txt is created at current folder
