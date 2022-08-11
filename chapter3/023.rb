n = STDIN.gets.chomp.to_i
blue_dice = STDIN.gets.chomp.split(" ").map(&:to_i)
red_dice = STDIN.gets.chomp.split(" ").map(&:to_i)

puts (blue_dice.sum / n.to_f) + (red_dice.sum / n.to_f)
