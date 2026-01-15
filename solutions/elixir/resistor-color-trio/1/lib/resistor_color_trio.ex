defmodule ResistorColorTrio do
  @colors [:black, :brown, :red, :orange, :yellow, :green, :blue, :violet, :grey, :white]

  @spec label(colors :: [atom]) :: {number, :ohms | :kiloohms | :megaohms | :gigaohms}
  def label([c1, c2, c3 | _]) do
    d1 = Enum.find_index(@colors, &(&1 == c1))
    d2 = Enum.find_index(@colors, &(&1 == c2))
    zeros = Enum.find_index(@colors, &(&1 == c3))

    value = (d1 * 10 + d2) * Integer.pow(10, zeros)

    cond do
      value >= 1_000_000_000 -> {value / 1_000_000_000, :gigaohms}
      value >= 1_000_000 -> {value / 1_000_000, :megaohms}
      value >= 1_000 -> {value / 1_000, :kiloohms}
      true -> {value, :ohms}
    end
  end
end