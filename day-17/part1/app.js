var steps = parseInt(process.argv[2], 10),
    buffer = [0, 1], index = 1, target = 2018;

for (var i = 2; i < target; i++) {
  buffer.splice(index = (index + steps) % buffer.length + 1, 0, i);
}

console.log(buffer[index + 1]);
