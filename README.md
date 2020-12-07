Go Version Manager

The version manager is a small lightweight manager to use multiply versions of `go` on your system.

Fast usage: `gvm install latest` 

## `gvm list`

List all go versions on the system.

## `gvm get (versions|lastest)`

Download the version.

## `gvm install (version|lastest)`

Download the version (if necessary) and configure the system to use this version. 

## `gvm lock (version)`

The version is locked and can't delete by clear.

## `gvm unlock (version)`

The version is unlocked and can delete by clear.

## `gvm clear`

Clear all not locked versions, the last two version are also locked. 