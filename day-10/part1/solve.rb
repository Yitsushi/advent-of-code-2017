lengths = File.open(ARGV[0]).read.split(',').map(&:to_i)
numbers = (0..255).to_a
cursor = 0
skip = 0

while lengths.length > 0 do
  diff = 0
  length = lengths.shift
  section = numbers[cursor...(cursor+length)]
  diff = length - section.length
  section += numbers[0...(diff)]
  section.reverse!
  numbers[cursor...(cursor+length-diff)] = section[0...(length - diff)]
  numbers[0...diff] = section[-diff..-1] unless diff.zero?
  cursor = (cursor +length + skip) % numbers.length
  skip += 1
end

puts numbers[0] * numbers[1]
