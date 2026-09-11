package web

import "embed"

//go:embed index.html index.js index.css
var Files embed.FS
