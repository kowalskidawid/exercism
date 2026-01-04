defmodule BinarySearch do
  @doc """
    Searches for a key in the tuple using the binary search algorithm.
    It returns :not_found if the key is not in the tuple.
    Otherwise returns {:ok, index}.

    ## Examples

      iex> BinarySearch.search({}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 5)
      {:ok, 2}

  """

  @spec search(tuple, integer) :: {:ok, integer} | :not_found
  def search(numbers, key) do
    search(numbers, key, 0, tuple_size(numbers) - 1)
  end
  
  def search(numbers, key) when tuple_size(numbers) == 1 do
    case elem(numbers, 0) == key do
      false -> :not_found
      true -> {:ok, 0}
    end
  end

  defp search(_numbers, _key, low, high) when low > high do
    :not_found
  end

  defp search(numbers, key, low, high) do
    mid = div(low + high, 2)
    current_value = elem(numbers, mid)

    cond do
      current_value == key ->
        {:ok, mid}

      current_value > key ->
        search(numbers, key, low, mid - 1)

      current_value < key ->
        search(numbers, key, mid + 1, high)
    end
  end
end
