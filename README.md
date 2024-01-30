# TOOLBOX

## Description

This is a collection of tools that I use in my daily work. I hope you find them useful.

## Network

### ping

Execute a ping command from the toolbox.

```shell
./build/toolbox net ping -u google.com
```

## Docker

### image
```shell
./build/toolbox docker image list
```

## File

### finder
```shell
./build/toolbox find text -c "cat /etc/passwd" -n 1,2 -d ":" -s "\t" -f 1
```
