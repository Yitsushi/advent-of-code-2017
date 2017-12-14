def knot_hash(string)
  lengths = string.split('').map(&:ord) + [17, 31, 73, 47, 23]
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

  numbers.each_slice(16)
    .map { |b| b.reduce(:^) }
    .map { |c| '%02x' % c }
    .join
end

base = ARGV[0];

number_of_ones = 0
128.times do |i|
  hash = knot_hash("#{base}-#{i}")
  hash.each_byte do |b|
    b = b.chr.to_i(16)
    number_of_ones += b.to_s(2).count('1')
  end
end

puts "Number of ones: #{number_of_ones}"
