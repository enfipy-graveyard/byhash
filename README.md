<h1 align="center">
  <img src="./logo.svg" width="300"/>
  <p align="center" style="font-size: 0.5em">Utility to get file data by hash</p>
</h1>

Utility `byhash` - retrieves folder data by provided hash

## Installation:

To install via `go get`:

```
go get github.com/enfipy/byhash
```

## Usage:

To lock by key:

```
locker := locker.Initialize()
locker.Lock("key")
<code>
locker.Unlock("key")
```

To lock only read:

```
locker := locker.Initialize()
locker.RLock("key")
<code>
locker.RUnlock("key")
```
