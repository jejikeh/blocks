#include "block.hh"
#include "raylib.h"

#define BLOCK_SIZE 32


block block::block_init(int x, int y) {
    block b = {};

    b.x = x;
    b.y = y;
    b.state = block_state::BLOCK_STATE_INACTIVE;

    return b;
}

void block::draw() {
    Color color = { 0, 0, 0, 0 };

    switch (state)
    {
    case block_state::BLOCK_STATE_INACTIVE:
        color = GRAY;
        break;
    
    case block_state::BLOCK_STATE_HIGHLITED:
        color = YELLOW;
        break;

    case block_state::BLOCK_STATE_SELECTED:
        color = RED;
        break;

    case block_state::BLOCK_STATE_ACTIVE:
        color = GREEN;
        break;

    default:
        break;
    }
    
    DrawRectangle(x, y, BLOCK_SIZE, BLOCK_SIZE, color);
}