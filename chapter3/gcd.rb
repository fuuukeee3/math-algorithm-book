a, b = STDIN.gets.chomp.split(" ").map(&:to_i)

min = [a, b].min
answer = 1
(2..min).each do |i|
  if a % i == 0 && b % i == 0
    answer = i
  end
end

puts answer
