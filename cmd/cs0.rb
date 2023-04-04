#!/usr/bin/env ruby

def usage()
  STDERR.puts("usage: #{File.basename(__FILE__)} n a b")
  exit(1)
end

if ARGV.length != 3
  usage()
end

n = ARGV[0].to_i()
if n <= 0
  usage()
end

a = ARGV[1].to_i()
if a < 0
  usage()
end

b = ARGV[2].to_i()
if b <= a or b > n
  usage()
end

hf = File.open("h-#{a}-#{b}.csv", 'w')
mf = File.open("m-#{a}-#{b}.csv", 'w')
lf = File.open("l-#{a}-#{b}.csv", 'w')

i = a
while i < b
  v1 = Hash.new(0)
  s1 = 0.0
  File.open("#{i}.txt.tfidf", 'r') do |f|
    while l = f.gets()
      k, v = l.split(',')
      v = v.to_f()
      v1[k] = v
      s1 += v * v
    end
  end
  j = i + 1
  while j < n
    v2 = Hash.new(0)
    s2 = 0.0
    File.open("#{j}.txt.tfidf", 'r') do |f|
      while l = f.gets()
        k, v = l.split(',')
        v = v.to_f()
        v2[k] = v
        s2 += v * v
      end
    end
    m = Math.sqrt(s1) * Math.sqrt(s2)
    if m > 0.0
      vs = v1
      vl = v2
      if vs.length > vl.length
        t = vl
        vl = vs
        vs = t
      end
      d = 0.0
      vs.each_pair do |t, f|
        d += f * vl[t]
      end
      c = d / m
      if c >= 0.80
        hf.write(sprintf("%d,%d,%0.4f\n", i, j, c))
      elsif c >= 0.60
        mf.write(sprintf("%d,%d,%0.4f\n", i, j, c))
      elsif c >= 0.40
        lf.write(sprintf("%d,%d,%0.4f\n", i, j, c))
      end
    end
    j += 1
  end
  i += 1
end

hf.close
mf.close
lf.close
