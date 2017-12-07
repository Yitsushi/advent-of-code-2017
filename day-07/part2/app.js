var fs = require('fs');

// Prep
var args = process.argv.slice(2);

var Node = function(name, weight) {
  this.name = name;
  this.weight = weight;
  this.isChild = false;
  this.children = [];
  this.realWeight = null;
};

var nodeList = [];
var rootNode = null;

var findOrCreateNode = function(name, weight) {
  for (var i = 0, l = nodeList.length; i < l; i++) {
    if (nodeList[i].name == name) {
      if (weight !== null && nodeList[i].weight === null) {
        nodeList[i].weight = weight;
      }
      return nodeList[i];
    }
  }

  var n = new Node(name, weight);
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

Node.prototype.findUnbalanced = function() {
  var weightList = this.children
    .map(function(c) { return c.realWeight; })
    .reduce(function(memo, current) {
      if (!memo.hasOwnProperty(current)) { memo[current] = 0; }
      memo[current]++;
      return memo;
    }, {});

  var wrongValue = -1;
  for (var weight in weightList) {
    if (weightList[weight] == 1) {
      wrongValue = parseInt(weight, 10);
      break;
    }
  }
  if (wrongValue === -1) {
    return false;
  }

  console.log(wrongValue);

  printThisLevel = false;
  wrongOne = null;
  for (var i = 0, l = this.children.length; i < l; i++) {
    if (this.children[i].realWeight == wrongValue) {
      //console.log(this.children[i]);
      var ret = this.children[i].findUnbalanced();
      printThisLevel = !ret;
      wrongOne = this.children[i];
      break;
    }
  }

  if (printThisLevel) {
    console.log("Name:", wrongOne.name);
    console.log("This weight:", wrongOne.weight);
    delete weightList[wrongValue];
    var goodValue = parseInt(Object.keys(weightList)[0], 10);
    console.log("Should be:", wrongOne.weight + goodValue - wrongValue);
  }
  return true;
};

Node.prototype.calculateWeight = function() {
  if (this.realWeight !== null) {
    return this.realWeight;
  }

  this.realWeight = this.weight;
  for (var i = 0, l = this.children.length; i < l; i++) {
    this.realWeight += this.children[i].calculateWeight();
  }
  return this.realWeight;
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

  rootNode.calculateWeight();
  rootNode.findUnbalanced();
};

// Main
if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parseFileContent);
