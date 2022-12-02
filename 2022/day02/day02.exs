IO.puts("Day 2")

{:ok, content} = File.read('input')

defmodule Rps do
  def rps(["A", "X"]), do: 1 + 3
  def rps(["A", "Y"]), do: 2 + 6
  def rps(["A", "Z"]), do: 3 + 0
  def rps(["B", "X"]), do: 1 + 0
  def rps(["B", "Y"]), do: 2 + 3
  def rps(["B", "Z"]), do: 3 + 6
  def rps(["C", "X"]), do: 1 + 6
  def rps(["C", "Y"]), do: 2 + 0
  def rps(["C", "Z"]), do: 3 + 3

  def rps2(["A", "X"]), do: rps(["A", "Z"])
  def rps2(["A", "Y"]), do: rps(["A", "X"])
  def rps2(["A", "Z"]), do: rps(["A", "Y"])
  def rps2(["B", "X"]), do: rps(["B", "X"])
  def rps2(["B", "Y"]), do: rps(["B", "Y"])
  def rps2(["B", "Z"]), do: rps(["B", "Z"])
  def rps2(["C", "X"]), do: rps(["C", "Y"])
  def rps2(["C", "Y"]), do: rps(["C", "Z"])
  def rps2(["C", "Z"]), do: rps(["C", "X"])
end

rounds = content
        |> String.split("\n")
        |> Enum.filter(fn s -> s != "" end)
        |> Enum.map(fn s -> String.split(s, " ") end)

sum = rounds
      |> Enum.map(&Rps.rps/1)
      |> Enum.sum

IO.inspect(sum)

sum2 = rounds
       |> Enum.map(&Rps.rps2/1)
       |> Enum.sum

IO.inspect(sum2)
