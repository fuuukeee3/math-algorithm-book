n, x, y = STDIN.gets.chomp.split(" ").map(&:to_i)

count = 0
(1..n).each do |i|
  count += 1 if i % x == 0 || i % y == 0
end

puts count
