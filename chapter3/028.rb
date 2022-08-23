n = STDIN.gets.chomp.to_i
h = STDIN.gets.chomp.split(" ").map(&:to_i)

dp = []
(0...n).each do |i|
  dp[i] = 0 if i == 0
  dp[i] = (h[i - 1] - h[i]).abs if i == 1

  if i >= 2
    tmp1 = dp[i - 1] + (h[i - 1] - h[i]).abs
    tmp2 = dp[i - 2] + (h[i - 2] - h[i]).abs
    dp[i] = [tmp1, tmp2].min
  end
end

puts dp[-1]
