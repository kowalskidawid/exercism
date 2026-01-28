defmodule Say do
  @doc """
  Translate a positive integer into English.
  """
  @spec in_english(integer) :: {atom, String.t()}
  def in_english(number) when number < 0 or number > 999_999_999_999 do
    {:error, "number is out of range"}
  end

  def in_english(0), do: {:ok, "zero"}

  def in_english(number) do
    {:ok, convert(number) |> String.trim()}
  end

  defp convert(0), do: ""

  defp convert(n) when n >= 1_000_000_000 do
    convert(div(n, 1_000_000_000)) <> " billion " <> convert(rem(n, 1_000_000_000))
  end

  defp convert(n) when n >= 1_000_000 do
    convert(div(n, 1_000_000)) <> " million " <> convert(rem(n, 1_000_000))
  end

  defp convert(n) when n >= 1_000 do
    convert(div(n, 1_000)) <> " thousand " <> convert(rem(n, 1_000))
  end

  defp convert(n) when n >= 100 do
    convert(div(n, 100)) <> " hundred " <> convert(rem(n, 100))
  end

  defp convert(n) when n >= 20 do
    prefix = elem({nil, nil, "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}, div(n, 10))
    suffix = convert(rem(n, 10))
    if suffix == "", do: prefix, else: prefix <> "-" <> suffix
  end

  defp convert(n) do
    elem({nil, "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
          "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}, n)
  end
end