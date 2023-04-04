#!/usr/bin/env ruby

def usage()
  STDERR.puts("usage: #{File.basename(__FILE__)} o n [m] [p]")
  exit(1)
end

if ARGV.length < 2 or ARGV.length > 4
  usage()
end

o = ARGV[0].to_i()
if o < 0
  usage()
end

n = ARGV[1].to_i()
if n <= 0
  usage()
end

m = 1
if ARGV.length >= 3
  m = ARGV[2].to_i()
  if m <= 0 or m > n
    usage()
  end
end

p = m
if ARGV.length == 4
  p = ARGV[3].to_i()
  if p < m or p > n
    usage()
  end
end

i = 0
while i < n
  unless File.exists?("#{i}.txt.tfidf")
  	exit(3)
  end
  i += 1
end

csns = Array.new(m)
pp = 0
pq = 0
pc = 0
ps = 0

i = 0
j = (n.to_f() / p.to_f()).ceil()
while i < n
  if pc < m
    k = i + j
    b = [k, n].min
    c = "cs#{o} #{n} #{i} #{b}"
    d = c.gsub(/\s+/, '_')
    csn = Process.fork do
      exec("#{c} > #{d}.out 2> #{d}.err")
    end
    csns[pq % m] = csn
    pq += 1
    STDERR.puts(c)
    pc += 1
    i = k
  end
  if pc == m
    csn = csns[pp % m]
    pp += 1
    Process.wait(csn)
    s = $?.exitstatus()
    if s != 0
      ps = s
    end
    pc -= 1
  end
end

while pc > 0
  csn = csns[pp % m]
  pp += 1
  Process.wait(csn)
  s = $?.exitstatus()
  if s != 0
    ps = s
  end
  pc -= 1
end

exit(ps)
