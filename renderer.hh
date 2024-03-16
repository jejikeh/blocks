#pragma once

#include <string>
#include "game.hh"

struct renderer {
    std::string title;

    int width;
    int height;

    int target_fps;

    static renderer renderer_init(std::string title, int width, int height);
    void run(game *g);
};