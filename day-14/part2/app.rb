DISK_SIZE = 128
DISK_LINE_LENGTH = 128
EMPTY_CELL = '0'
FULL_CELL = '1'
GROUP_REPLACEMENT_CHARACTER = 'x'

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

def kill_group(disk, index)
  queue = [index]
  until queue.empty?
    index = queue.shift
    next if index < 0 || index > disk.length
    next if disk[index] != FULL_CELL

    disk[index] = GROUP_REPLACEMENT_CHARACTER
    queue << index - 1 unless index % DISK_LINE_LENGTH == 0
    queue << index + 1 unless index % DISK_LINE_LENGTH == DISK_LINE_LENGTH - 1
    queue << index - DISK_LINE_LENGTH
    queue << index + DISK_LINE_LENGTH
  end
end

base = ARGV[0];
disk = ''

DISK_SIZE.times do |i|
  hash = knot_hash("#{base}-#{i}")
  hash.each_byte do |b|
    b = b.chr.to_i(16)
    disk += "%04d" % b.to_s(2)
  end
end

number_of_found_groups = 0
while index = disk.index(FULL_CELL) do
  kill_group(disk, index)
  number_of_found_groups += 1
end

puts "#{number_of_found_groups}"
