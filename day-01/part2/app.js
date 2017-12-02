var fs = require('fs');

// Prep
var args = process.argv.slice(2);

// Functions
var parse = function(err, data) {
  if (err) {
    throw err;
  }
  data = data.trim();
  pool = data + data;

  var sum = 0;

  var pad = data.length / 2;

  for (var i = 0, _l = data.length; i < _l; i++) {
    if (pool[i] == pool[i + pad]) {
      sum += parseInt(data[i], 10);
    }
  }

  console.log(sum);
}

// Main
if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parse);
