defmodule Prime do
  @spec nth(pos_integer) :: pos_integer
  def nth(n) when n < 1, do: raise ArgumentError
  def nth(n), do: find(n, 2)

  defp find(1, current) do
    if prime?(current), do: current, else: find(1, current + 1)
  end

  defp find(count, current) do
    if prime?(current) do
      find(count - 1, current + 1)
    else
      find(count, current + 1)
    end
  end

  defp prime?(2), do: true
  defp prime?(n) when n < 2 or rem(n, 2) == 0, do: false
  defp prime?(n) do
    limit = trunc(:math.sqrt(n))
    check_div(n, 3, limit)
  end

  defp check_div(_n, i, limit) when i > limit, do: true
  defp check_div(n, i, limit) do
    if rem(n, i) == 0, do: false, else: check_div(n, i + 2, limit)
  end
end
