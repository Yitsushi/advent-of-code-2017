var fs = require('fs');

// Prep
var args = process.argv.slice(2);

var Node = function(name, num) {
  this.name = name;
  this.num = num;
  this.isChild = false;
  this.children = [];
};

var nodeList = [];
var rootNode = null;

var findOrCreateNode = function(name, num) {
  for (var i = 0, l = nodeList.length; i < l; i++) {
    if (nodeList[i].name == name) {
      if (num !== null && nodeList[i].num === null) {
        nodeList[i].num = num;
      }
      return nodeList[i];
    }
  }

  var n = new Node(name, num);
  nodeList.push(n);
  return n;
}

Node.prototype.addChild = function(child) {
  for (var i = 0, l = this.children.length; i < l; i++) {
    if (child == this.children[i]) {
      return false;
    }
  }

  this.children.push(child);
};

Node.prototype.findChild = function(name) {
  if (this.name == name) {
    return this;
  }

  for (var i = 0, _l = this.children.length; i < _l; i++) {
    var c = this.children[i].findChild(name);
    if (c !== null) return c;
  }

  return null;
};

var parseLine = function(line) {
  if (/->/.test(line)) {
    // has child
    var m = line.match(/^([a-z]+) \((\d+)\) -> (.*)$/);
    var n = findOrCreateNode(m[1], parseInt(m[2], 10));
    var children = m[3].split(', ');
    for (var i = 0, l = children.length; i < l; i++) {
      var sub = findOrCreateNode(children[i], null);
      sub.isChild = true;
      n.addChild(sub);
    }
  } else {
    // no child
    var m = line.match(/^([a-z]+) \((\d+)\)$/);
    var n = findOrCreateNode(m[1], parseInt(m[2], 10));
    n.isChild = true;
  }
};

var parseFileContent = function(err, data) {
  if (err) {
    throw err;
  }

  data = data.trim().split("\n");
  for (var i = 0, l = data.length; i < l; i++) {
    parseLine(data[i]);
  }

  for (var i = 0, l = nodeList.length; i < l; i++) {
    if (nodeList[i].isChild === false) {
      rootNode = nodeList[i];
      break;
    }
  }

  console.log(rootNode);
};

// Main
if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parseFileContent);
