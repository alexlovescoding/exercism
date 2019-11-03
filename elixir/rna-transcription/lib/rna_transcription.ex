defmodule RnaTranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  def to_rna_helper([head | tail], rna) do
    dna_to_rna_map = %{
      ?G => ?C,
      ?C => ?G,
      ?T => ?A,
      ?A => ?U
    }
    to_rna_helper(tail, rna ++ [dna_to_rna_map[head]])
  end

  def to_rna_helper([], rna) do
    rna
  end

  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    to_rna_helper(dna, [])
  end
end
