n = STDIN.gets.chomp.to_i

results = []
(2..n).each do |i|
  flg = true
  (2..(i - 1)).each do |ii|
    if i % ii == 0
      flg = false
      break
    end
  end
  results << i if flg
end

puts results.join(" ")
