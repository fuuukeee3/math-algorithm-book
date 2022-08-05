a, b = STDIN.gets.chomp.split(" ").map(&:to_i)

answer = 1
loop do
  if a < b
    b = b % a
  else
    a = a % b
  end

  if a == 0 || b == 0
    answer = [a, b].max
    break
  end
end

puts answer
