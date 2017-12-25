var fs = require('fs');

var LONG_TERM_MEANS = 1000;

// Prep
var args = process.argv.slice(2);

var Coordinate = function(x, y, z) {
  this.x = x;
  this.y = y;
  this.z = z;
};

Coordinate.prototype.ManhattanDistance = function() {
  return Math.abs(this.x) + Math.abs(this.y) + Math.abs(this.z);
};

Coordinate.prototype.ToString = function() {
  return '<' + [this.x, this.y, this.z].join(',') + '>';
};

var Particle = function(pArr, vArr, aArr) {
  this.position = new Coordinate(...pArr);
  this.velocity = new Coordinate(...vArr);
  this.acceleration = new Coordinate(...aArr);
};

Particle.prototype.Move = function() {
  this.velocity.x += this.acceleration.x;
  this.velocity.y += this.acceleration.y;
  this.velocity.z += this.acceleration.z;

  this.position.x += this.velocity.x;
  this.position.y += this.velocity.y;
  this.position.z += this.velocity.z;
};

Particle.prototype.ManhattanDistance = function() {
  return this.position.ManhattanDistance();
};

var ParticleSystem = function() {
  this.universe = [];
  this.positionCache = {};
};

ParticleSystem.prototype.Simulate = function() {
  console.log(this.universe.length);

  for (var i = 0; i < LONG_TERM_MEANS; i++) {
    this.positionCache = {};
    for (var j = 0, l = this.universe.length; j < l; j++) {
      if (this.universe[j] === false) { continue; }
      this.universe[j].Move();
      if (!this.positionCache.hasOwnProperty(this.universe[j].position.ToString())) {
        this.positionCache[this.universe[j].position.ToString()] = j;
      }
      if (this.positionCache[this.universe[j].position.ToString()] !== j) {
        this.universe[this.positionCache[this.universe[j].position.ToString()]] = false;
        this.universe[j] = false;
      }
    }
    this.universe = this.universe.filter(function(n) { return n !== false });
  }

  console.log("Length:", this.universe.length);
};

ParticleSystem.prototype.AddParticle = function(p) {
  this.universe.push(p);
};

var parseFileContent = function(err, data) {
  if (err) {
    throw err;
  }

  data = data.trim().split(/\n/);

  var particleSystem = new ParticleSystem();

  for (var i = 0, l = data.length; i < l; i++) {
    var v = data[i];
    var d = v.match(/p=<([0-9,\-]+)>, v=<([0-9,\-]+)>, a=<([0-9,\-]+)>/)
    //console.log(d[1], d[2], d[3])
    particleSystem.AddParticle(
      new Particle(
        d[1].split(',').map(function(n) { return parseInt(n, 10); }),
        d[2].split(',').map(function(n) { return parseInt(n, 10); }),
        d[3].split(',').map(function(n) { return parseInt(n, 10); })
      )
    );
  }

  particleSystem.Simulate();
};

if (args.length < 1) {
  console.log('First argument must be a file path');
  process.exit();
}

fs.readFile(args[0], 'utf8', parseFileContent);
