n, s = STDIN.gets.chomp.split(" ").map(&:to_i)
a = STDIN.gets.chomp.split(" ").map(&:to_i)

result = false
(0..n).each do |i|
  if a.combination(i).to_a.map(&:sum).include?(s)
    result = true
    break
  end
end

puts result ? "Yes" : "No"
