def prime_number?(number)
  return false if number == 1

  limit = Math.sqrt(number).floor
  (2..limit).each do |i|
    return false if number % i == 0
  end
  return true
end

n = STDIN.gets.chomp.to_i

results = []
(2..n).each do |i|
  next unless prime_number?(i)

  flg = false
  loop do
    div, mod = n.divmod(i)
    if mod == 0
      results << i
      n = div
      if prime_number?(n)
        results << n
        flg = true
        break
      end
    else
      break
    end
  end

  break if flg
end

puts results.join(" ")
