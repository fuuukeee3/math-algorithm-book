n = STDIN.gets.chomp.to_i
numbers = STDIN.gets.chomp.split(" ").map(&:to_i)

count = 0
numbers.combination(2) { |arr| count += 1 if arr.sum == 100000 }

puts count
