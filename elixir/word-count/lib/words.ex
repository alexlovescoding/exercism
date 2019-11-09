defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence = String.downcase(sentence)
    words = String.split(sentence, ~r/[^[:alnum:]-]/u)
    words = Enum.filter(words, &(&1 != ""))
    count(words, %{})
  end

  @spec count([String.t()], map) :: map
  defp count([head | tail], countMap) do
    count(tail, Map.update(countMap, head, 1, &(&1 + 1)))
  end

  @spec count([], map) :: map
  defp count([], countMap) do
    countMap
  end
end
