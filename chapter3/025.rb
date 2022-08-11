n = STDIN.gets.chomp.to_i
a = STDIN.gets.chomp.split(" ").map(&:to_i)
b = STDIN.gets.chomp.split(" ").map(&:to_i)

result = 0.0
n.times do |i|
  result += (a[i].to_f * 2 / 6) + (b[i].to_f * 4 / 6)
end

puts result
