defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  def compare(a, b) do
    cond do
      a === b -> :equal
      is_sublist?(a, b) -> :sublist
      is_sublist?(b, a) -> :superlist
      true -> :unequal
    end
  end

  defp is_sublist?([], _), do: true
  defp is_sublist?(_, []), do: false
  defp is_sublist?(small, large) do
    List.starts_with?(large, small) or is_sublist?(small, tl(large))
  end
end
