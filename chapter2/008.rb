n, s = STDIN.gets.chomp.split(" ").map(&:to_i)

count = 0
(1..n).each do |red|
  (1..n).each do |blue|
    count += 1 if red + blue <= s
  end
end

puts count
