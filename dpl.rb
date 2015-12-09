require 'rubygems'
require 'dpl/cli'
require 'json'

EXCEPTED_KEYS = %w(image environment ports)

input = begin 
          JSON.parse(STDIN.read)
        rescue Exception => e
          puts e
          exit 1
        end

vargs = input['vargs']
worksapce = input['workspace']['path']

FileUtils.cd(worksapce)

if vargs.nil? || vargs.empty?
  puts "ERROR: no vargs given, check your .drone.yml"
  exit 1
end

args = vargs.inject([]) do |arr, (key, value)|
  if EXCEPTED_KEYS.include?(key)
    next
  end

  if value.is_a?(TrueClass) || value == 'true'
    arr << "--#{key}"
  else
    arr << "--#{key}=#{value}"
  end
end

args.insert(0, "--provider=rubygems")
args.uniq!

cli = DPL::CLI.new(args)
cli.run
