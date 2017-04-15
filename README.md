# go-envfs
Useful for dynamically creating a set of files in a docker container which runs in an orchestrated/clustered environment that may not have persistent volumes available.

## Usage
In the spirit of 12-factor apps, this binary will create a filesystem based on a base64 encoded yaml file set as an environment variable named `ENVFS_YAML_B64`.

Example:

```bash
# create a yaml file describing the filesystem
~$ cat fs.yaml 
---

files:
  /tmp/f1: "first file."
  /tmp/f2: |
    second file.
    this file has two lines

# install this binary
~$ go get github.com/gbolo/go-envfs

# run it with the yaml file in base64 env var
~$ ENVFS_YAML_B64=$(base64 fs.yaml) $GOPATH/bin/go-envfs
Initializing go-envfs version 0.1 ...
Attempting base64 decode of: ENVFS_YAML_B64 ...
Attempting to Unmarshal yaml ...
Number of files to write: 2
Writing file with mode 0644: /tmp/f1
Writing file with mode 0644: /tmp/f2
go-envfs has completed execution. Exiting.

# check the files
~$ cat /tmp/f1 && cat /tmp/f2
```
