#!/usr/bin/env ruby

require 'nokogiri'

if ARGV.length != 1
  STDIN.puts("usage: #{File.basename(__FILE__)} n")
  exit(1)
end

n = ARGV[0].to_i()

BLACKLIST = %w[title script style]

i = 0
while i < n
	f = File.open("#{i}.html", 'r')
	h = Nokogiri::HTML(f)
	f.close()
    ns = h.search('//text()')
    BLACKLIST.each do |t|
      ns -= h.search("//#{t}/text()")
    end
    s = ns.text
    s.scrub!
    s.gsub!(/[\s\W[:space:]]+/, ' ')
    ts = s.split(' ')
    ts.select! do |t|
      t.length >= 2 or t.length <= 25
    end
    s = ts.join(' ')
    s.downcase!()
    IO.write("#{i}.txt", s)
	i += 1
end
