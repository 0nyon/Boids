# Boids Flocking Simulation

A 2D flocking simulation built in Go utilizing Craig Reynolds' Boids algorithm. This project implements emergent collective behavior through distributed agent logic, leveraging Go's powerful concurrency primitives and the Raylib-go library for hardware-accelerated rendering.

## Features

*   **Emergent Behavior:** Watch individual boids interact using simple local rules to form complex, lifelike flocking patterns.
*   **Concurrent Architecture:** Uniquely structured so that each boid runs independently via native Go concurrency, rather than sequentially through a single thread.
*   **Real-time 2D Visualization:** Smooth, high-performance rendering powered by Raylib.

## The Three Rules of Flocking

The simulation relies on three fundamental steering behaviors applied to each boid based on its local neighborhood:

1.  **Separation:** Steer to avoid crowding or colliding with local flockmates.
2.  **Alignment:** Steer towards the average heading and velocity of local flockmates.
3.  **Cohesion:** Steer toward the "center of mass" (average position) of local flockmates.

## Prerequisites

Before running the project, ensure you have the following installed on your system:

*   **Go** (version 1.18 or higher recommended)
*   **CGO Dependencies:** Raylib-go requires a C compiler (like `gcc`) for compiling the underlying graphics bindings.
    *   *Ubuntu/Debian:* `sudo apt-get install alsa-utils libasound2-dev libx11-dev libxrandr-dev libxi-dev libxcursor-dev libxinerama-dev libxkbcommon-dev`
    *   *macOS:* Xcode Command Line Tools are required (`xcode-select --install`).
    *   *Windows:* Ensure `gcc` (via MinGW-w64 or similar) is configured in your system PATH.

## Getting Started

1. Clone the repository:
```bash
   git clone [https://github.com/yourusername/boids-simulation.git](https://github.com/yourusername/boids-simulation.git)
   cd boids-simulation

## Acknowledgments & Resources

* **Craig Reynolds:** For developing the original groundbreaking [Boids algorithm in 1986](https://www.red3d.com/cwr/boids/).
* **V. Hunter Adams (Cornell University):** Special thanks to Hunter Adams for his excellent [Boids Algorithm Write-up](https://vanhunteradams.com/Pico/Animal_Movement/Boids-algorithm.html). It really helps you understand how to implement each rule of the algorithm, and was very easy to understand!