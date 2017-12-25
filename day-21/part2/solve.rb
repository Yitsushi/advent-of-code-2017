rules = {}
File.open(ARGV[0]).read.strip
    .scan(/(\S+) => (\S+)/)
    .map { |k, v| [k.split('/'), v.split('/')] }
    .each { |k, v| rules[k] = v };

def rotateClockwise(pattern)
  ## clockwise
  pattern.map(&:chars).transpose.map(&:join).map(&:reverse)
  ## anti-clockwise
  #pattern.map(&:chars).transpose.map(&:join).reverse
  ## 180
  #pattern.reverse.map(&:reverse)
end

def flip(pattern)
  pattern.reverse
end

def draw(pattern)
  on = 0
  puts "Length: #{pattern.length**2}"
  puts "On state: #{pattern.map { |c| c.count('#') }.inject(&:+)}"
  #pattern.each { |r| puts r }
  puts ''
end

picture = ['.#.', '..#', '###']
draw picture

18.times do
  size = picture.length
  block_size = 2 if size % 2 == 0
  block_size = 3 if block_size.nil? && size % 3 == 0

  number_of_blocks = size / block_size

  parts = []
  # row
  number_of_blocks.times do |i|
    # column
    number_of_blocks.times do |j|
      pattern = picture[i*block_size...(i*block_size+block_size)].map() { |c| c[j*block_size...(j*block_size+block_size)] }

      unless rules[pattern].nil?
        parts << rules[pattern]
        next
      end

      found = false
      [pattern, flip(pattern)].each do |pat|
        3.times do
          pat = rotateClockwise(pat)
          unless rules[pat].nil?
            parts << rules[pat]
            found = true
            break
          end
        end unless found
      end
    end
  end

  picture = []
  parts.each_with_index do |part, i|
    part.each_with_index do |b, j|
      row = (i / number_of_blocks).floor * b.size
      picture[row + j] ||= ''
      picture[row + j] += b
    end
  end

  draw picture
end


