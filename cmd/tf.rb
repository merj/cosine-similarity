#!/usr/bin/env ruby

def usage()
  STDERR.puts("usage: #{File.basename(__FILE__)} n")
  exit(1)
end

if ARGV.length != 1
  usage()
end

n = ARGV[0].to_i()
if n <= 0
  usage()
end

i = 0
while i < n
  fn = "#{i}.txt"
  unless File.exists?(fn)
    exit(3)
  end
  tf = Hash.new(0)
  d = IO.read(fn).split
  d.each do |k|
    tf[k] += 1.0
  end
  File.open("#{i}.txt.tf", 'w') do |f|
    s = d.length.to_f()
    tf.each_pair do |k, v|
      # term frequency adjusted for document length
      v /= s
      f.write(sprintf("%s,%0.10f\n", k, v))
    end
  end
  i += 1
end
