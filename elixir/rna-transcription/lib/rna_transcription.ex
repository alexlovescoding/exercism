defmodule RnaTranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  defp to_rna([head | tail], rna) do
    dna_to_rna_map = %{
      ?G => ?C,
      ?C => ?G,
      ?T => ?A,
      ?A => ?U
    }
    to_rna(tail, [dna_to_rna_map[head] | rna])
  end

  defp to_rna([], rna) do
    Enum.reverse(rna)
  end

  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    to_rna(dna, [])
  end
end
