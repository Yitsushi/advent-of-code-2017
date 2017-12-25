<?php

define('INFECTED', 1);
define('CLEAN', -1);

function rotate_vector($vector, $deg) {
  return [
    (int)(($vector[0] * cos($deg * M_PI/180)) - ($vector[1] * sin($deg * M_PI/180))),
    (int)(($vector[0] * sin($deg * M_PI/180)) + ($vector[1] * cos($deg * M_PI/180)))
  ];
}

function build_map($input_file) {
  $fh = fopen($input_file, "r");
  if (!$fh) {
    echo "File error", PHP_EOL;
    exit(0);
  }

  $map = [];
  while (($line = fgets($fh)) !== false) {
    $map[] = str_split(trim($line));
  }
  fclose($fh);

  $center = [
    'x' => floor(count($map[0]) / 2),
    'y' => floor(count($map) / 2)
  ];

  $valueMap = [
    '#' => INFECTED,
    '.' => CLEAN
  ];
  $real_map = [];
  foreach ($map as $key => $value) {
    $real_map[$key - $center['y']] = [];
    foreach ($value as $k => $v) {
      $real_map[$key - $center['y']][$k - $center['x']] = $valueMap[$v];
    }
  }

  return $real_map;
}

$map = build_map($argv[1]);
$facing = [-1, 0];
$position = [0, 0];
$infections = 0;

for ($burst = 0; $burst < 10000; $burst++) {
  # If the current node is infected, it turns to its right.
  # Otherwise, it turns to its left.
  $facing = rotate_vector($facing, -90 * $map[$position[0]][$position[1]]);

  # If the current node is clean, it becomes infected.
  # Otherwise, it becomes cleaned.
  $map[$position[0]][$position[1]] *= -1;

  # Count infections
  $infections += ($map[$position[0]][$position[1]] + 1) / 2;

  # The virus carrier moves forward one node
  # in the direction it is facing.
  $position = [
    $position[0] + $facing[0],
    $position[1] + $facing[1]
  ];

  # Extend map if need
  if (!isset($map[$position[0]][$position[1]])) {
    $map[$position[0]][$position[1]] = CLEAN;
  }
}

echo "Infections: ", $infections, PHP_EOL;
