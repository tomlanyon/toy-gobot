# toy-gobot

toy-gobot is an implementation of the Toy Robot puzzle in Go.

The spec for the puzzle was taken from [duncan-bayne/toy-robot](https://github.com/duncan-bayne/toy-robot) and is reproduced below as a reference.

## spec
toy-robot reads instructions from STDIN, executing them one at a time until EOF is reached.  (On a Linux or OSX system, type C-d to generate an EOF character).

### valid commands

#### PLACE X,Y,FACING

Put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.  If the robot is already placed, issuing another *valid* PLACE command will place the robot in the newly specified location.

#### MOVE

Moves the toy robot one unit forward in the direction it is currently facing.

#### LEFT

Rotates the robot 90 degrees to the left (i.e. counter-clockwise) without changing the position of the robot.

#### RIGHT

Rotates the robot 90 degrees to the right (i.e. clockwise) without changing the position of the robot.

#### REPORT

Announces the X,Y and F of the robot.

