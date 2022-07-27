n = STDIN.gets.chomp.to_i
total = STDIN.gets.chomp.split(" ").map(&:to_i).sum
puts total % 100
