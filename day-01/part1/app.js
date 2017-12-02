var fs = require('fs');

// Prep
var args = process.argv.slice(2);

// Functions
var parse = function(err, data) {
  if (err) {
    throw err;
  }

  var sum = 0;

  data = data.trim();
  data += data[0];

  for (var i = 0, _l = data.length - 1; i < _l; i++) {
    if (data[i] == data[i+1]) {
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
