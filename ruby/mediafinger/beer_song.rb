class BeerSong
  def lyrics
    verses(99, 0)
  end

  def verses(higher, lower)
    song = (lower..higher).reverse_each.map do |bottles|
      verse(bottles)
    end
    song.join("\n")
  end

  def verse(bottles)
    case bottles
    when 1
      <<~TEXT
        1 bottle of beer on the wall, 1 bottle of beer.
        Take it down and pass it around, no more bottles of beer on the wall.
      TEXT
    when 0
      <<~TEXT
        No more bottles of beer on the wall, no more bottles of beer.
        Go to the store and buy some more, 99 bottles of beer on the wall.
      TEXT
    else
      <<~TEXT
        #{bottles} #{pluralize('bottle', bottles)} of beer on the wall, #{bottles} #{pluralize('bottle', bottles)} of beer.
        Take one down and pass it around, #{bottles - 1} #{pluralize('bottle', bottles - 1)} of beer on the wall.
      TEXT
    end
  end

  def pluralize(string, amount)
    return string if amount == 1
    "#{string}s"
  end
end
