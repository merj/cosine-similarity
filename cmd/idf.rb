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

df = Hash.new(0)

i = 0
while i < n
  fn = "#{i}.txt.tf"
  unless File.exists?(fn)
  	exit(3)
  end
  File.open(fn, 'r') do |f|
    while l = f.gets()
      k, _ = l.split(',')
      df[k] += 1.0
    end
  end
  i += 1
end

File.open('idf', 'w') do |f|
  s = n.to_f()
  df.each_pair do |k, v|
    # inverse document frequency
    v = Math.log(s / (1 + v))
    f.write(sprintf("%s,%0.10f\n", k, v))
  end
end
