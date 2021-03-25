#!/bin/bash

echo "running visualise-execution"
go run $(head -1 go.mod | cut -d' ' -f2)/cmd/visualise-execution

echo "running dot"
for f in dot/*.dot; do dot -Gsize=20,5\! -Gdpi=100 -Tpng $f > $f.png; done

echo "generating whitespace"
for f in dot/*.png; do convert $f -gravity center -background white -extent 2000x500 $f.final; done

echo "converting to gif"
convert -delay 50 -loop 0 dot/*.final animation.gif
