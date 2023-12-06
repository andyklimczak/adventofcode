defmodule DayOne do

  def solve do
    IO.puts('Day1')

    {:ok, content} = File.read('input')

    result1 = content
              |> String.split("\n")
              |> IO.inspect
              |> Enum.map(fn s -> Regex.replace(~r/\D/, s, "") end)
              |> IO.inspect
              |> Enum.filter(fn s -> s != "" end)
              |> Enum.map(fn s -> String.at(s, 0) <> String.at(s, -1) end)
              |> IO.inspect
              |> Enum.map(fn a -> String.to_integer(a) end)
              |> IO.inspect
              |> Enum.sum

    IO.puts result1

    IO.puts("---")

    result2 = content
              |> String.split("\n")
              |> IO.inspect
              |> Enum.map(fn s ->
                replace_num(s)
              end)
              |> IO.inspect
              |> Enum.map(fn s -> Regex.replace(~r/\D/, s, "") end)
              |> IO.inspect
              |> Enum.filter(fn s -> s != "" end)
              |> Enum.map(fn s -> String.at(s, 0) <> String.at(s, -1) end)
              |> IO.inspect
              |> Enum.map(fn a -> String.to_integer(a) end)
              |> IO.inspect
              |> Enum.sum

    IO.puts result2
  end

  def replace_num("one" <> rest), do: "1" <> replace_num("e" <> replace_num(rest))
  def replace_num("two" <> rest), do: "2" <> replace_num("o" <> replace_num(rest))
  def replace_num("three" <> rest), do: "3" <> replace_num("e" <> replace_num(rest))
  def replace_num("four" <> rest), do: "4" <> replace_num("r" <> replace_num(rest))
  def replace_num("five" <> rest), do: "5" <> replace_num("e" <> replace_num(rest))
  def replace_num("six" <> rest), do: "6" <> replace_num("x" <> replace_num(rest))
  def replace_num("seven" <> rest), do: "7" <> replace_num("n" <> replace_num(rest))
  def replace_num("eight" <> rest), do: "8" <> replace_num("t" <> replace_num(rest))
  def replace_num("nine" <> rest), do: "9" <> replace_num("e" <> replace_num(rest))
  def replace_num(<<char, rest::binary>>), do: <<char>> <> replace_num(rest)
  def replace_num(""), do: ""
end

DayOne.solve
