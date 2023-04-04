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

fn = 'idf'
unless File.exists?(fn)
  exit(3)
end
idf = Hash.new(0)
File.open(fn, 'r') do |f|
  while l = f.gets()
    k, v = l.split(',')
    idf[k] = v.to_f()
  end
end

i = 0
while i < n
  fn = "#{i}.txt.tf"
  unless File.exists?(fn)
    exit(3)
  end
  tfidf = Hash.new(0)
  File.open(fn, 'r') do |f|
    while l = f.gets()
      k, v = l.split(',')
      # term frequency * inverse document frequency
      tfidf[k] = v.to_f() * idf[k]
    end
  end
  File.open("#{i}.txt.tfidf", 'w') do |f|
    tfidf.each_pair do |k, v|
      f.write(sprintf("%s,%0.10f\n", k, v))
    end
  end
  i += 1
end
