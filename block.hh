#pragma once

enum class block_state {
    BLOCK_STATE_INACTIVE = 0,
    BLOCK_STATE_HIGHLITED = 1,
    BLOCK_STATE_SELECTED = 2,
    BLOCK_STATE_ACTIVE = 3
};

struct block {
    int x, y;

    block_state state;

    static block block_init(int x, int y);

    void draw();
};