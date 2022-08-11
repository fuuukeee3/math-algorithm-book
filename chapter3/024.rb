n = STDIN.gets.chomp.to_i

answers = []
n.times do |i|
  a = {}
  a[:count], a[:score] = STDIN.gets.chomp.split(" ").map(&:to_i)
  answers << a
end

result = answers.reduce(0.0) do |a, answer|
  a += answer[:score] / (answer[:count]).to_f
end

puts result
