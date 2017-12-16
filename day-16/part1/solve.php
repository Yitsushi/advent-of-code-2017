<?php

class DanceFloor {
  var $abc;

  public function __construct() {
    $this->abc = 'abcdefghijklmnop';
  }

  public function spin(int $num) {
    $this->abc = substr(
      $this->abc,
      strlen($this->abc) - $num,
      strlen($this->abc)
    ) . substr(
      $this->abc,
      0,
      strlen($this->abc) - $num
    );
  }

  public function exchange(int $indexA, int $indexB) {
    $tmp = $this->abc[$indexA];
    $this->abc[$indexA] = $this->abc[$indexB];
    $this->abc[$indexB] = $tmp;
  }

  public function partner(string $a, string $b) {
    $this->exchange(strpos($this->abc, $a), strpos($this->abc, $b));
  }
}


$commands = explode(',', trim(file_get_contents($argv[1])));
$df = new DanceFloor();
foreach ($commands as $command) {
  switch ($command[0]) {
  case 's':
    list($num) = sscanf($command, 's%d');
    $df->spin($num);
    break;
  case 'x':
    list($ia, $ib) = sscanf($command, 'x%d/%d');
    $df->exchange($ia, $ib);
    break;
  case 'p':
    list($a, $b) = sscanf($command, 'p%c/%c');
    $df->partner($a, $b);
    break;
  }
}

echo $df->abc;
