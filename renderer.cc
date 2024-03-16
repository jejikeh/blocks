#include "renderer.hh"
#include "raylib.h"

#define MAX_FPS 60

renderer renderer::renderer_init(std::string title, int width, int height) {
    renderer r = {};

    r.title = title;
    r.width = width;
    r.height = height;
    r.target_fps = MAX_FPS;

    InitWindow(width, height, title.c_str());
    SetTargetFPS(r.target_fps);

    return r;
}

void renderer::run(game *g) {
    g->start();

    while (!WindowShouldClose()) {
        g->update();

        BeginDrawing();

        g->draw();

        EndDrawing();
    }

    CloseWindow();
}
