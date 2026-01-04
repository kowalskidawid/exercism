defmodule BirdCount do
  def today([]) do
    nil
  end
  
  def today(list) do
    Enum.at(list, 0)
  end

  def increment_day_count([]), do: [1]
  
  def increment_day_count([today | rest]) do
    [ today + 1 | rest]
  end

  def has_day_without_birds?(list) do
    0 in list
  end

  def total(list) do
     Enum.sum(list)
  end

  def busy_days(list) do
    Enum.count(list, fn count -> count >= 5 end)
  end
end
