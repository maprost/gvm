# Go Version Manager

The version manager is a small lightweight manager to install and change your `go` version on your system.

**Important:** you need an installed go version (root path: `/usr/local/go`), 
also the root path need your user account as owner. 

## Install

`go get github.com/maprost/gvm`

## Usage

Fast usage: `gvm install latest` 

### `gvm list`

List all go versions on the system.

### `gvm get (version|lastest)`

Download the version (latest versin).

### `gvm install (version|lastest)`

Download the version (if necessary) and configure the system to use this version. 

## Future implementation

### `gvm lock (version)`

The version is locked and can't delete by clear.

### `gvm unlock (version)`

The version is unlocked and can delete by clear.

### `gvm clear`

Clear all not locked versions, the last two version are also locked. 