instruction = File.open(ARGV[0]).read.split("\n").map(&:to_i)

head = 0
max = instruction.length
jumps = 0

while (head < max)
  n = instruction[head]
  instruction[head] += 1
  head += n
  jumps += 1
end

puts jumps
