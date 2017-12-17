var steps = parseInt(process.argv[2], 10),
    index = 1, target = 50000000, after_zero = -1;

for (var i = 1; i < target; i++) {
  if ((index = ((index + steps) % i) + 1) == 1) {
    after_zero = i;
  }
}

console.log(after_zero);
