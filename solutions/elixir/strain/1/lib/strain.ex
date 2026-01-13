defmodule Strain do

  @spec keep(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def keep([], _fun), do: []

  def keep([head | tail], fun) do
    if fun.(head) do
      [head | keep(tail, fun)]
    else
      keep(tail, fun)
    end
  end

  @spec discard(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def discard(list, fun) do
    keep(list, fn item -> not fun.(item) end)
  end
end