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

$severity = 0;
foreach ($layers as $layer => $range) {
  if ($layer % (($range - 1) * 2) == 0) {
    $severity += $layer * $range;
  }
}
echo "Severity: ${severity}", PHP_EOL;
