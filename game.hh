#pragma once

#include "block.hh"

struct game {
    int x;

    block block;

    static game game_init();

    void start();
    void update();
    void draw();
};