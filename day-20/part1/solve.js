var fs = require('fs');

var LONG_TERM_MEANS = 100000;

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
};

ParticleSystem.prototype.Simulate = function() {
  for (var i = 0; i < LONG_TERM_MEANS; i++) {
    for (var j = 0, l = this.universe.length; j < l; j++) {
      this.universe[j].Move();
    }
  }

  var min_particle = { index: -1, value: Infinity };
  for (var i = 0, l = this.universe.length; i < l; i++) {
    var v = this.universe[i];
    if (v.ManhattanDistance() < min_particle.value) {
      min_particle.index = i;
      min_particle.value = v.ManhattanDistance();
    }
  }
  console.log(min_particle);
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
