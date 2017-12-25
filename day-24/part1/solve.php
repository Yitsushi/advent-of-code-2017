<?php

function collect_components($input_file) {
  $fh = fopen($input_file, "r");
  if (!$fh) {
    echo "File error", PHP_EOL;
    exit(0);
  }

  $components = [];
  while (($line = fgets($fh)) !== false) {
    $components[] = explode('/', trim($line));
  }
  fclose($fh);

  return $components;
}

function build_bridge($components, $base) {
  $possible_routes = [];
  foreach ($components as $index => $value) {
    if (!in_array($base, $value)) {
      continue;
    }

    $current = ['length' => 1, 'strength' => array_sum($value)];
    $rest = $components;
    array_splice($rest, $index, 1);

    $myindex = array_search($base, $value);
    $target_value = $value[1 - $myindex];

    $ret = build_bridge($rest, $target_value);
    if (count($ret) > 0) {
      foreach ($ret as $v) {
        if (count($v) < 1) {
          continue;
        }
        $possible_routes[] = [
          'length' => $v['length'] + $current['length'],
          'strength' => $v['strength'] + $current['strength']
        ];
      }
    } else {
      $possible_routes[] = $current;
    }
  }

  return $possible_routes;
}

$components = collect_components($argv[1]);

$clusterfuck = build_bridge($components, 0);

$max = [
  'longest' => [ 'length' => 0, 'strength' => 0 ],
  'strongest' => [ 'length' => 0, 'strength' => 0 ]
];

$max = array_reduce(
  $clusterfuck,
  function($carry, $item) {
    // Strongest
    if ($carry['strongest']['strength'] < $item['strength']) {
      $carry['strongest'] = $item;
    }
    // Longest
    if ($carry['longest']['length'] <= $item['length']) {
      if ($carry['longest']['strength'] < $item['strength']) {
        $carry['longest'] = $item;
      }
    }

    return $carry;
  },
  $max
);

var_dump($max);
