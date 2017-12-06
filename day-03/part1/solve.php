<?php

$target = $argv[1];
# our sprital starts from zero
$target--;

# Calculate the radius
$radius = floor((sqrt($target + 1) - 1) / 2) + 1;
# because we start from the center
$target--;

# calculate radius - 1 points
$rsum = (8 * $radius * ($radius - 1)) / 2;
# face length
$len = $radius * 2;

# shift all the values so we get a spiral not just squares
$shifted = (1 + $target - $rsum) % ($radius * 8);

$position = [0, 0, $radius];

# 0 -> up
# 1 -> right
# 2 -> down
# 3 -> left
$direction = floor($shifted / ($radius * 2));
switch ($direction) {
  case 0:
    $position[0] = $shifted - $radius;
    $position[1] = -$radius;
    break;
  case 1:
    $position[0] = $radius;
    $position[1] = ($shifted % $len) - $radius;
    break;
  case 2:
    $position[0] = $radius - ($shifted % $len);
    $position[1] = $radius;
    break;
  case 3:
    $position[0] = -$radius;
    $position[1] = $radius - ($shifted % $len);
    break;
}

echo "Distance: ", (abs($position[0]) + abs($position[1])), PHP_EOL;


