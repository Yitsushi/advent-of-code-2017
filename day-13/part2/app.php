<?php

$fh = fopen($argv[1], "r");
if (!$fh) {
  echo "File error", PHP_EOL;
  exit(0);
}

$layers = [];
while (($line = fgets($fh)) !== false) {
  list($layer, $range) = sscanf($line, "%d: %d");
  $layers[$layer] = $range;
}
fclose($fh);

$wait = 0;
while (true) {
  $fail = false;
  foreach ($layers as $layer => $range) {
    if (($layer + $wait) % (($range - 1) * 2) == 0) {
      $fail = true;
      break;
    }
  }
  if (!$fail) {
    break;
  }
  $wait++;
}

echo "Wait time: ", $wait, PHP_EOL;
