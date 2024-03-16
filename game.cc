#include "game.hh"
#include "raylib.h"

game game::game_init() {
    game g = {};

    g.block = block::block_init(500, 300);

    return g;
}

void game::start() {

}

void game::update() {
    if (IsKeyDown(KEY_SPACE)) {
        x++;
    }
}

void game::draw() {
    ClearBackground(WHITE);

    block.draw();
}