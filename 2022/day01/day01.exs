IO.puts('Day1')

{:ok, content} = File.read('input')
summed_calories = content
        |> String.split("\n\n")
        |> Enum.map(fn s -> String.split(s, "\n") end)
        |> Enum.filter(fn s -> s != [""] end)
        |> Enum.map(fn l -> Enum.map(l, &String.to_integer/1) end)
        |> Enum.map(fn a -> Enum.sum(a) end)

max = Enum.max(summed_calories)
IO.inspect(max)

largest_three_sum = summed_calories
                    |> Enum.sort(:desc)
                    |> Enum.take(3)
                    |> Enum.sum
IO.inspect(largest_three_sum)
