# bhq

Simple CLI tool for managing Backlog issues

# DESCRIPTION

`bhq` provides a way to organize backlog issues, like go get does.

```
$ bhq get https://spaceid.backlog.jp/view/PROJECT-1

# Runs `mkdir -p ~/.bhq/PROJECT-1`
```

# SYNOPSIS

```
bhg get      get backlog task
bhq whoami   get backlog user
bhq list     list backlog issues
bhq comment  add a comment to the current issue
bhq help     Shows a list of commands or help for one command
```
# DIRECTORY STRUCTURES

```
~/.bhq
|-- PROJECT-1
`-- PROJECT-2
```

# INSTALLATION

```sh
go get github.com/acro5piano/bhq
export BACKLOG_API_KEY=YOUR_API_KEY
```

# Thanks

[motemen/ghq: Remote repository management made easy](https://github.com/motemen/ghq)
