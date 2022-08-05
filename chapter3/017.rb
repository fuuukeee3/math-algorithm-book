def gcd(a, b)
  result = 1
  loop do
    if a > b
      a = a % b
    else
      b = b % a
    end

    if a == 0 || b == 0
      result = [a, b].max
      break
    end
  end
  result
end

n = STDIN.gets.chomp.to_i
a_arr = STDIN.gets.chomp.split(" ").map(&:to_i)

a = nil
answer = 1
while a_arr.size > 0
  if a.nil?
    a = a_arr.shift
  else
    a = answer
  end
  b = a_arr.shift

  answer = gcd(a, b)
end

puts answer
