defmodule Username do
  def sanitize([]), do: []

  def sanitize([head | tail]) do
    case head do
      ?ä -> 'ae' ++ sanitize(tail)
      ?ö -> 'oe' ++ sanitize(tail)
      ?ü -> 'ue' ++ sanitize(tail)
      ?ß -> 'ss' ++ sanitize(tail)
      char when char >= ?a and char <= ?z -> [char | sanitize(tail)]
      ?_ -> [?_ | sanitize(tail)]
      _ -> sanitize(tail)
    end
  end
end
