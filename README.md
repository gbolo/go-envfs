# go-envfs
create filesystem from env var

## Usage
in the spirit of 12-factor apps, this binary will create a filesystem based on a base64 encoded yaml file.

```bash
# create a yaml file describing the filesystem
$ cat fs.yaml 
---

files:
  /tmp/f1: "first file."
  /tmp/f2: |
    second file.
    this file has two lines

# install this binary
$ go get github.com/gbolo/go-envfs

# run it with the yaml file in base64 env var
$ ENV_FS_B64=$(base64 fs.yaml) $GOPATH/bin/go-envfs
Initializing ENV-2-FS...
[DEBUG] ENV_FS_B64: LS0tCgpmaWxlczoKICAvdG1wL2YxOiAiZmlyc3QgZmlsZS4iCiAgL3RtcC9mMjogfAogICAgc2Vj
b25kIGZpbGUuCiAgICB0aGlzIGZpbGUgaGFzIHR3byBsaW5lcwo=
[DEBUG] []byte: [45 45 45 10 10 102 105 108 101 115 58 10 32 32 47 116 109 112 47 102 49 58 32 34 102 105 114 115 116 32 102 105 108 101 46 34 10 32 32 47 116 109 112 47 102 50 58 32 124 10 32 32 32 32 115 101 99 111 110 100 32 102 105 108 101 46 10 32 32 32 32 116 104 105 115 32 102 105 108 101 32 104 97 115 32 116 119 111 32 108 105 110 101 115 10] 

Num of Files: 2
Writing file: /tmp/f1
Writing file: /tmp/f2

# check the files
$ cat /tmp/f1 && cat /tmp/f2
```
