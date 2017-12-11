var fs = require('fs');

// Prep
var args = process.argv.slice(2);

var parseFileContent = function(err, data) {
  if (err) {
    throw err;
  }

  data = data.trim();

  // remove ignored parts
  data = data.replace(/!./g, '');

  // remove garbage
  data = data.replace(/<[^>]*>/g, '');

  var score = 0; depth = 0;
  for (var i = 0, l = data.length; i < l; i++) {
    var v = data[i];
    if (v == '{') {
      depth++;
      score += depth;
    }
    if (v == '}') {
      depth--;
    }
  }

  console.log(depth, score);
};


if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parseFileContent);
