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
    cond do
      head == "" ->
        count(tail, countMap)
      Map.has_key?(countMap, head) ->
        count(tail, %{countMap | head => countMap[head] + 1})
      true ->
        count(tail, Map.put(countMap, head, 1))
    end
  end

  @spec count([], map) :: map
  defp count([], countMap) do
    countMap
  end
end
