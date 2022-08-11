def factorical(num)
  (1..num).reduce(1, :*)
end

n, r = STDIN.gets.chomp.split(" ").map(&:to_i)
puts factorical(n) / (factorical(r) * factorical(n - r))
