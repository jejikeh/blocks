#include "raylib.h"

#include "renderer.hh"
#include "game.hh"

// static renderer renderer_init(std::string title, int width, int height);
// static game     game_init();

int main(void)
{
    auto r = renderer::renderer_init("blocks", 800, 600);
    auto g = game::game_init();

    r.run(&g);

    return 0;
}