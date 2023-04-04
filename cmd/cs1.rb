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

vis = Array.new(n)
sis = Array.new(n)

hf = File.open("h-#{a}-#{b}.csv", 'w')
mf = File.open("m-#{a}-#{b}.csv", 'w')
lf = File.open("l-#{a}-#{b}.csv", 'w')

i = 0
while i < n
  vi = Hash.new(0)
  si = 0.0
  File.open("#{i}.txt.tfidf", 'r') do |f|
    while l = f.gets()
      k, v = l.split(',')
      v = v.to_f()
      vi[k] = v
      si += v * v
    end
  end
  vis[i] = vi
  sis[i] = si
  i += 1
end

i = a
while i < b
  v1 = vis[i]
  s1 = sis[i]
  j = i + 1
  while j < n
    v2 = vis[j]
    s2 = sis[j]
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
