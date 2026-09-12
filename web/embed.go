package web

import "embed"

//go:embed index.html index.js index.css assets
var Files embed.FS
