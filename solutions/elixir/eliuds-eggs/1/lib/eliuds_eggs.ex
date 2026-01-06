defmodule EliudsEggs do
  @doc """
  Given the number, count the number of eggs.
  """
  @spec egg_count(number :: integer()) :: non_neg_integer()
  def egg_count(number) do
    do_count(number, 0)
  end

  defp do_count(0, count), do: count

  defp do_count(number, count) do
    do_count(div(number, 2), count + rem(number, 2))
  end
end
