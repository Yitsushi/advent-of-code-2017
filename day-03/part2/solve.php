<?php

$target = $argv[1];

# Initialize the memory
$memory = [[1]];

# Initialize Head
$x = 1; $y = 0; $r = 1;

while (true) {
  $nextValue = 0;
  # Calculate values based on neighbors
  for($i = $x - 1; $i <= $x + 1; $i++) {
    for($j = $y - 1; $j <= $y + 1; $j++) {
      $nextValue += (isset($memory[$i][$j]) ? $memory[$i][$j] : 0);
    }
  }

  # If it's larger or equal, we are done
  if ($nextValue >= $target) {
    echo "Distance: ", (abs($x) + abs($y)), ". Value: ", $nextValue, PHP_EOL;
    break;
  }

  # Save the value
  if (!isset($memory[$x])) { $memory[$x] = []; }
  $memory[$x][$y] = $nextValue;

  # Move the Head
  if ($x == $r) {
    if ($y > -$r) { $y--; } else { $x--; }
  } elseif ($x == -$r) {
    if ($y < $r) { $y++; } else { $x++; }
  } else {
    $x += (($y > 0) * 2) - 1;
  }

  if (isset($memory[$x][$y])) { $r++; $y++; $x++; }
}
