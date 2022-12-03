IO.puts("Day 3")

{:ok, content} = File.read("input")

defmodule Day3 do
  def split_half(string) do
    half = div(String.length(string), 2)
    [String.slice(string, 0..half-1), String.slice(string, half..String.length(string))]
  end

  def get_value(char) when char in 65..90  do
    char - 64 + 26
  end

  def get_value(char) when char in 97..122  do
    char - 96
  end
end

sum1 = content
       |> String.split("\n", trim: true)
       |> Enum.map(&Day3.split_half/1)
       |> Enum.map(fn [a, b] -> [String.split(a, "", trim: true), String.split(b, "", trim: true)] end)
       |> Enum.map(fn [a, b] -> [MapSet.new(a), MapSet.new(b)] end)
       |> Enum.map(fn [a, b] -> MapSet.intersection(a, b) end)
       |> Enum.map(&MapSet.to_list/1)
       |> Enum.map(fn [a] -> :binary.first(a) end)
       |> Enum.map(&Day3.get_value/1)
       |> Enum.sum

IO.inspect(sum1)

sum2 = content
       |> String.split("\n", trim: true)
       |> Enum.chunk_every(3)
       |> Enum.map(fn a -> Enum.map(a, fn b -> String.split(b, "", trim: true) end) end)
       |> Enum.map(fn a -> Enum.map(a, fn b -> MapSet.new(b) end) end)
       |> Enum.map(fn a -> Enum.reduce(a, List.first(a), fn s, acc -> MapSet.intersection(s, acc) end) end)
       |> Enum.map(&MapSet.to_list/1)
       |> Enum.map(fn [a] -> :binary.first(a) end)
       |> Enum.map(&Day3.get_value/1)
       |> Enum.sum

IO.inspect(sum2)
