lengths = File.open(ARGV[0]).read.strip.split('').map(&:ord) + [17, 31, 73, 47, 23]
numbers = (0..255).to_a
cursor = 0
skip = 0

64.times do
  l_buf = lengths.dup
  while l_buf.length > 0 do
    diff = 0
    length = l_buf.shift
    section = numbers[cursor...(cursor+length)]
    diff = length - section.length
    section += numbers[0...(diff)]
    section.reverse!
    numbers[cursor...(cursor+length-diff)] = section[0...(length - diff)]
    numbers[0...diff] = section[-diff..-1] unless diff.zero?
    cursor = (cursor +length + skip) % numbers.length
    skip += 1
  end
end

puts numbers.each_slice(16)
            .map { |b| b.reduce(:^) }
            .map { |c| '%02x' % c }
            .join
