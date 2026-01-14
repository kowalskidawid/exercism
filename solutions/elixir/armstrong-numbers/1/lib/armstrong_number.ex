defmodule ArmstrongNumber do
  @spec valid?(integer) :: boolean
  def valid?(number) do
    digits = Integer.digits(number)
    len = length(digits)
    number == Enum.sum(for d <- digits, do: Integer.pow(d, len))
  end
end