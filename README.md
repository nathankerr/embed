# Embed

embeds external files into go programs

## Installation

    go get github.com/nathankerr/embed

## Use

1. Create directory in your project source called `embed`.
2. Run `embed` to create `embed.go`
3. Reference file contents by `embed["filename"]`

## Current Limitations

- directories under embed not supported
- no compression, etc. of file contents