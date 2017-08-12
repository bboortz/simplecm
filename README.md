# simplecm
simplecm is a simple configuration management

## Build

* ./build.sh

## Usage

* sudo ./simplecm

## Features

* update os
* install packages via package manager
* uninstall packages via package manager
* sync time
* manage users
* manage groups
* create symlinks

## Support

* OS
 * Arch Linux
 * Debian / Ubuntu (PLANNED)

## Build Dependencies

These tools must be installed first

* docker
* github.com/bboortz/go-build 

## Runtime Dependencies

* github.com/bboortz/go-utils

## Architecture

*TBD*
   
main
|			|	|	|	|
task			start	exit	logger	reflect
|	|
|	helper
|	|
command

