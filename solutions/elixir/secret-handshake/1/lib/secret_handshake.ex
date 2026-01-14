defmodule SecretHandshake do
  use Bitwise

  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.
  """
  @spec commands(code :: integer) :: list(String.t())
  def commands(code) do
    actions = []
  
    actions = if (code &&& 1) == 1, do: actions ++ ["wink"], else: actions
    actions = if (code &&& 2) == 2, do: actions ++ ["double blink"], else: actions
    actions = if (code &&& 4) == 4, do: actions ++ ["close your eyes"], else: actions
    actions = if (code &&& 8) == 8, do: actions ++ ["jump"], else: actions
    if (code &&& 16) == 16 do
      Enum.reverse(actions)
    else
      actions
    end
  end
end