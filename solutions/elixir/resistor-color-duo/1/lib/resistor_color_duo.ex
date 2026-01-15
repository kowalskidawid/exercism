defmodule ResistorColorDuo do
  @colors [:black, :brown, :red, :orange, :yellow, :green, :blue, :violet, :grey, :white]

  @spec value(colors :: [atom]) :: integer
  def value([c1, c2 | _]) do
    d1 = Enum.find_index(@colors, &(&1 == c1))
    d2 = Enum.find_index(@colors, &(&1 == c2))
    d1 * 10 + d2
  end
end