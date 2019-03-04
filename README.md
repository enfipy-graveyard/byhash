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

Result running `byhash ls dist/`:

```bash
f1.txt 47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad
f2.txt bdda36034608af5b6b61c81585c67ad2458af25c62fc3f7be036c29de99739f4
```

Result running `byhash get dist/ 47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad`:

```bash
hello world
```
