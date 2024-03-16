#!/bin/sh

clang++ -framework CoreVideo -framework IOKit -framework Cocoa -framework GLUT -framework OpenGL libraylib.a main.cc renderer.cc game.cc block.cc  -o blocks -I./raylib/src -I./

./blocks