<?php

define('CLEAN', 0);
define('WEAKENED', 1);
define('INFECTED', 2);
define('FLAGGED', 3);

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

function draw($map) {
  system('clear');
  foreach ($map as $row) {
    echo implode('', array_map(function ($n) {
      return ['.', 'W', '#', 'F'][$n];
    }, $row)), PHP_EOL;
  }
  sleep(1);
}

$map = build_map($argv[1]);
$facing = [-1, 0];
$position = [0, 0];
$infections = 0;

for ($burst = 0; $burst < 10000000; $burst++) {
  if ($burst % 100000 == 0) {
    echo ($burst / 100000), "%\r";
  }
  # If it is clean, it turns left.
  # If it is weakened, it does not turn,
  #   and will continue moving in the same direction.
  # If it is infected, it turns right.
  # If it is flagged, it reverses direction,
  #   and will go back the way it came.
  switch ($map[$position[0]][$position[1]]) {
    case CLEAN:
      $facing = rotate_vector($facing, 90);
      break;
    case INFECTED:
      $facing = rotate_vector($facing, -90);
      break;
    case FLAGGED:
      $facing = rotate_vector($facing, 180);
      break;
  }

  # Clean nodes become weakened.
  # Weakened nodes become infected.
  # Infected nodes become flagged.
  # Flagged nodes become clean.
  $map[$position[0]][$position[1]] = ($map[$position[0]][$position[1]] + 1) % 4;

  # Count infections
  if ($map[$position[0]][$position[1]] == INFECTED) {
    $infections++;
  }

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
