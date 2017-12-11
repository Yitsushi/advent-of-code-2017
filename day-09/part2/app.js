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
  var score = 0;
  var blocks = data.match(/<[^>]*>/g);
  for (var i = 0, l = blocks.length; i < l; i++) {
    score += blocks[i].length - 2;
  }
  console.log(score);
};


if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parseFileContent);
