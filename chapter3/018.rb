n = STDIN.gets.chomp.to_i
prices = STDIN.gets.chomp.split(" ").map(&:to_i)

counts = {
  100 => 0,
  200 => 0,
  300 => 0,
  400 => 0,
}

prices.each do |price|
  counts[price] += 1
end

puts counts[100] * counts[400] + counts[200] * counts[300]
