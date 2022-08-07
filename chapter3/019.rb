n = STDIN.gets.chomp.to_i
colors = STDIN.gets.chomp.split(" ").map(&:to_i)

red = colors.count(1)
yellow = colors.count(2)
blue = colors.count(3)

red_r = red * (red - 1) / 2
yellow_r = yellow * (yellow - 1) / 2
blue_r = blue * (blue - 1) / 2

puts red_r + yellow_r + blue_r
