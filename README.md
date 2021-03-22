# Cracking the coding interview

This is my take on exercises from the "Cracking the coding interview book.

One thing I wanted to try for a long while, is to generate useful visualisations
of how the solutions are executed, so I'm generating graphviz dot files,
and am turning them into gifs.


## Commands to generate a gif from a bunch of dot files

```sh
for f in dot/*.dot; do dot -Tpng $f > $f.png; done
convert -delay 50 -loop 0 dot/*.png animation.gif
```
