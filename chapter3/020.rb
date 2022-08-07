n = STDIN.gets.chomp.to_i
numbers = STDIN.gets.chomp.split(" ").map(&:to_i)

count = 0
(0...n).each do |a1|
  ((a1 + 1)...n).each do |a2|
    ((a2 + 1)...n).each do |a3|
      ((a3 + 1)...n).each do |a4|
        ((a4 + 1)...n).each do |a5|
          count += 1 if numbers[a1] + numbers[a2] + numbers[a3] + numbers[a4] + numbers[a5] == 1000
        end
      end
    end
  end
end

puts count
