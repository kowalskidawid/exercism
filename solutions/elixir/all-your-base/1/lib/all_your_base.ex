defmodule AllYourBase do
  @spec convert(list, integer, integer) :: {:ok, list} | {:error, String.t()}
  def convert(_, input_base, _) when input_base < 2, do: {:error, "input base must be >= 2"}
  def convert(_, _, output_base) when output_base < 2, do: {:error, "output base must be >= 2"}

  def convert(digits, input_base, output_base) do
    if Enum.any?(digits, fn x -> x < 0 or x >= input_base end) do
      {:error, "all digits must be >= 0 and < input base"}
    else
      decimal = Integer.undigits(digits, input_base)
      {:ok, Integer.digits(decimal, output_base)}
    end
  end
end