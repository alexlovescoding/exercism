defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    count(String.split(String.downcase(sentence), ~r/[^[:alnum:]-]/u), %{})
  end

  @spec count([String.t()], map) :: map
  defp count([head | tail], countMap) do
    if head == "" do
      count(tail, countMap)
    else
      count(tail, Map.update(countMap, head, 1, &(&1 + 1)))
    end
  end

  @spec count([], map) :: map
  defp count([], countMap) do
    countMap
  end
end
