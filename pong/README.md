
# The New Pong! 

In this game, we mix pong and bricks! Idea is to build step-by-step logic to:

* Learn Go on the fly.
* Build a Go web-service for supporting the online version of the game

Here is an image of what the game shoudl look like!

![Game Preview](game.png)

## Core Logic

1. Set the position of the play area, paddle, ball and smiley-bricks!
2. Ball will keep bouncing off the play area wall i.e. it will reflect off the wall.
3. Smiley-bricks are the target - everytime the ball destroys a brick, you get 1 point.
   * When a smiley-brick is destroyed, it appears again in another location

4. Move the paddle to bounce the ball around.
   * If the ball hits the paddle in the middle, it will reflect exactly!
   * The more awat from the middle the ball hits the paddel, the larger the angle of relection!
5. The game is timed and high scores are maintained for maximum number of smiley-bricks hit.
6. Every time 4 points are scored,

   * The speed of the ball increases by 10%!
   * The size of the smiley-brick reduces by 10%!

7. The ball destroys the smiely-brick and continues it's same trajectory

## Single-player mode:

1. Paddle is positioned in the middle and moves horizonaltally only!
2. Smiley-bricks are positioned in the top and bottom edges of the play-area

## Two-player-mode:

1. Paddles are positioned vertically on the left and right edges of the play area.
2. Smiley-brick is positioned in the middle 40% of the play-area
3. Player who hits the smiley-brick gets a point.

